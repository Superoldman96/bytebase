package mysql

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	parserbase "github.com/bytebase/bytebase/backend/plugin/parser/base"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	"github.com/bytebase/bytebase/backend/store/model"
)

var (
	_ advisor.Advisor = (*OnlineMigrationAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLOnlineMigration, &OnlineMigrationAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLOnlineMigration, &OnlineMigrationAdvisor{})
}

// OnlineMigrationAdvisor is the advisor checking for using gh-ost to migrate large tables.
type OnlineMigrationAdvisor struct {
}

// Check checks for using gh-ost to migrate large tables.
func (*OnlineMigrationAdvisor) Check(ctx context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	minRows := int64(payload.Number)

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	dbSchema := model.NewDatabaseSchema(checkCtx.DBSchema, nil, nil, storepb.Engine_MYSQL, checkCtx.IsObjectCaseSensitive)
	title := string(checkCtx.Rule.Type)

	// Check gh-ost database existence first if the change type is gh-ost.
	if checkCtx.ChangeType == storepb.PlanCheckRunConfig_DDL_GHOST {
		ghostDatabaseName := common.BackupDatabaseNameOfEngine(storepb.Engine_MYSQL)
		if !advisor.DatabaseExists(ctx, checkCtx, ghostDatabaseName) {
			return []*storepb.Advice{
				{
					Status:        level,
					Title:         title,
					Content:       fmt.Sprintf("Needs database %q to save temporary data for online migration but it does not exist", ghostDatabaseName),
					Code:          advisor.DatabaseNotExists.Int32(),
					StartPosition: common.FirstLinePosition,
				},
			}, nil
		}
	}

	var adviceList []*storepb.Advice
	// Check statements.
	for _, stmt := range stmtList {
		checker := &useGhostChecker{
			currentDatabase:  checkCtx.CurrentDatabase,
			changedResources: make(map[string]parserbase.SchemaResource),
			baseline:         int32(stmt.BaseLine),
			checkCtx:         checkCtx,
		}

		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)

		if !checker.ghostCompatible {
			continue
		}

		for _, resource := range checker.changedResources {
			var tableRows int64
			if table := dbSchema.GetDatabaseMetadata().GetSchema(resource.Schema).GetTable(resource.Table); table != nil {
				tableRows = table.GetRowCount()
			}
			if tableRows >= minRows {
				adviceList = append(adviceList, &storepb.Advice{
					Status:        level,
					Code:          advisor.AdviseOnlineMigrationForStatement.Int32(),
					Title:         title,
					Content:       fmt.Sprintf("Estimated table row count of %q is %d exceeding the set value %d. Consider using online migration for this statement", fmt.Sprintf("%s.%s", resource.Schema, resource.Table), tableRows, minRows),
					StartPosition: checker.start,
					EndPosition:   checker.end,
				})
			}
		}
	}

	// More than one statements need online migration.
	// Advise running each statement in separate issues.
	if len(adviceList) > 1 {
		return adviceList, nil
	}
	// One statement needs online migration, others don't.
	// Advise running the statement in another issue.
	if len(adviceList) == 1 && len(stmtList) > 1 {
		return adviceList, nil
	}

	// We have only one statement, and the statement
	// needs online migration.
	// Advise to enable online migration for the issue, or return OK if it's already enabled.
	if len(adviceList) == 1 && len(stmtList) == 1 {
		if checkCtx.ChangeType == storepb.PlanCheckRunConfig_DDL_GHOST {
			return nil, nil
		}

		adviceList[0].Code = advisor.AdviseOnlineMigration.Int32()
		return adviceList, nil
	}

	// No statement needs online migration.
	// Advise to disable online migration if it's enabled.
	if len(adviceList) == 0 {
		if checkCtx.ChangeType == storepb.PlanCheckRunConfig_DDL_GHOST {
			return []*storepb.Advice{{
				Status:  level,
				Code:    advisor.AdviseNoOnlineMigration.Int32(),
				Title:   title,
				Content: "Advise to disable online migration because found no statements that need online migration",
			}}, nil
		}
		return nil, nil
	}

	// Should never reach this.
	return nil, nil
}

type useGhostChecker struct {
	*mysql.BaseMySQLParserListener
	checkCtx advisor.Context

	currentDatabase  string
	changedResources map[string]parserbase.SchemaResource
	ghostCompatible  bool

	baseline int32
	start    *storepb.Position
	end      *storepb.Position
}

func (c *useGhostChecker) EnterAlterStatement(ctx *mysql.AlterStatementContext) {
	c.start = common.ConvertANTLRPositionToPosition(
		&common.ANTLRPosition{
			Line:   int32(ctx.GetStart().GetLine()),
			Column: int32(ctx.GetStart().GetColumn()),
		},
		c.checkCtx.Statements,
	)
}

func (c *useGhostChecker) ExitAlterStatement(ctx *mysql.AlterStatementContext) {
	c.end = common.ConvertANTLRPositionToPosition(
		&common.ANTLRPosition{
			Line:   c.baseline + int32(ctx.GetStop().GetLine()),
			Column: int32(ctx.GetStop().GetColumn() + len([]rune(ctx.GetStop().GetText()))),
		},
		c.checkCtx.Statements,
	)
}

func (c *useGhostChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	resource := parserbase.SchemaResource{
		Database: c.currentDatabase,
	}
	db, table := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	if db != "" {
		resource.Database = db
	}
	resource.Table = table
	c.changedResources[resource.String()] = resource
}

func (c *useGhostChecker) EnterAlterTableActions(ctx *mysql.AlterTableActionsContext) {
	c.ghostCompatible = ctx.AlterCommandList() != nil || ctx.PartitionClause() != nil || ctx.RemovePartitioning() != nil
}
