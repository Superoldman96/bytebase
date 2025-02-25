package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"sort"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

const (
	maxDefaultCurrentTimeColumCount   = 2
	maxOnUpdateCurrentTimeColumnCount = 1
)

var (
	_ advisor.Advisor = (*ColumnCurrentTimeCountLimitAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLCurrentTimeColumnCountLimit, &ColumnCurrentTimeCountLimitAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLCurrentTimeColumnCountLimit, &ColumnCurrentTimeCountLimitAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLCurrentTimeColumnCountLimit, &ColumnCurrentTimeCountLimitAdvisor{})
}

// ColumnCurrentTimeCountLimitAdvisor is the advisor checking for current time column count limit.
type ColumnCurrentTimeCountLimitAdvisor struct {
}

// Check checks for current time column count limit.
func (*ColumnCurrentTimeCountLimitAdvisor) Check(ctx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnCurrentTimeCountLimitChecker{
		level:    level,
		title:    string(ctx.Rule.Type),
		tableSet: make(map[string]tableData),
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.generateAdvice(), nil
}

type columnCurrentTimeCountLimitChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	tableSet   map[string]tableData
}

func (checker *columnCurrentTimeCountLimitChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableElementList() == nil || ctx.TableName() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	for _, tableElement := range ctx.TableElementList().AllTableElement() {
		if tableElement.ColumnDefinition() == nil || tableElement.ColumnDefinition().FieldDefinition() == nil || tableElement.ColumnDefinition().FieldDefinition().DataType() == nil {
			continue
		}
		_, _, columnName := mysqlparser.NormalizeMySQLColumnName(tableElement.ColumnDefinition().ColumnName())
		checker.checkTime(tableName, columnName, tableElement.ColumnDefinition().FieldDefinition())
	}
}

func (checker *columnCurrentTimeCountLimitChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.AlterTableActions() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	if tableName == "" {
		return
	}
	// alter table add column, change column, modify column.
	for _, item := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if item == nil {
			continue
		}

		var columnName string
		switch {
		// add column
		case item.ADD_SYMBOL() != nil:
			switch {
			case item.Identifier() != nil && item.FieldDefinition() != nil:
				columnName := mysqlparser.NormalizeMySQLIdentifier(item.Identifier())
				checker.checkTime(tableName, columnName, item.FieldDefinition())
			case item.OPEN_PAR_SYMBOL() != nil && item.TableElementList() != nil:
				for _, tableElement := range item.TableElementList().AllTableElement() {
					if tableElement.ColumnDefinition() == nil || tableElement.ColumnDefinition().ColumnName() == nil || tableElement.ColumnDefinition().FieldDefinition() == nil {
						continue
					}
					_, _, columnName := mysqlparser.NormalizeMySQLColumnName(tableElement.ColumnDefinition().ColumnName())
					checker.checkTime(tableName, columnName, tableElement.ColumnDefinition().FieldDefinition())
				}
			}
		// change column.
		case item.CHANGE_SYMBOL() != nil && item.ColumnInternalRef() != nil && item.Identifier() != nil && item.FieldDefinition() != nil:
			if item.FieldDefinition().DataType() == nil {
				continue
			}
			// only focus on new column name.
			columnName = mysqlparser.NormalizeMySQLIdentifier(item.Identifier())
			checker.checkTime(tableName, columnName, item.FieldDefinition())
		// modify column.
		case item.MODIFY_SYMBOL() != nil && item.ColumnInternalRef() != nil && item.FieldDefinition() != nil:
			if item.FieldDefinition().DataType() == nil {
				continue
			}
			columnName = mysqlparser.NormalizeMySQLColumnInternalRef(item.ColumnInternalRef())
			checker.checkTime(tableName, columnName, item.FieldDefinition())
		default:
			continue
		}
	}
}

func (checker *columnCurrentTimeCountLimitChecker) checkTime(tableName string, _ string, ctx mysql.IFieldDefinitionContext) {
	if ctx.DataType() == nil {
		return
	}

	switch ctx.DataType().GetType_().GetTokenType() {
	case mysql.MySQLParserDATETIME_SYMBOL, mysql.MySQLParserTIMESTAMP_SYMBOL:
		if checker.isDefaultCurrentTime(ctx) {
			table, exists := checker.tableSet[tableName]
			if !exists {
				table = tableData{
					tableName: tableName,
				}
			}
			table.defaultCurrentTimeCount++
			table.line = checker.baseLine + ctx.GetStart().GetLine()
			checker.tableSet[tableName] = table
		}
		if checker.isOnUpdateCurrentTime(ctx) {
			table, exists := checker.tableSet[tableName]
			if !exists {
				table = tableData{
					tableName: tableName,
				}
			}
			table.onUpdateCurrentTimeCount++
			table.line = checker.baseLine + ctx.GetStart().GetLine()
			checker.tableSet[tableName] = table
		}
	}
}

func (*columnCurrentTimeCountLimitChecker) isDefaultCurrentTime(ctx mysql.IFieldDefinitionContext) bool {
	for _, attr := range ctx.AllColumnAttribute() {
		if attr == nil || attr.GetValue() == nil {
			continue
		}
		if attr.GetValue().GetTokenType() == mysql.MySQLParserDEFAULT_SYMBOL && attr.NOW_SYMBOL() != nil {
			return true
		}
	}
	return false
}

func (*columnCurrentTimeCountLimitChecker) isOnUpdateCurrentTime(ctx mysql.IFieldDefinitionContext) bool {
	for _, attr := range ctx.AllColumnAttribute() {
		if attr == nil || attr.GetValue() == nil {
			continue
		}
		if attr.GetValue().GetTokenType() == mysql.MySQLParserON_SYMBOL && attr.NOW_SYMBOL() != nil {
			return true
		}
	}
	return false
}

func (checker *columnCurrentTimeCountLimitChecker) generateAdvice() []*storepb.Advice {
	var tableList []tableData
	for _, table := range checker.tableSet {
		tableList = append(tableList, table)
	}
	sort.Slice(tableList, func(i, j int) bool {
		return tableList[i].line < tableList[j].line
	})
	for _, table := range tableList {
		if table.defaultCurrentTimeCount > maxDefaultCurrentTimeColumCount {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.DefaultCurrentTimeColumnCountExceedsLimit.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("Table `%s` has %d DEFAULT CURRENT_TIMESTAMP() columns. The count greater than %d.", table.tableName, table.defaultCurrentTimeCount, maxDefaultCurrentTimeColumCount),
				StartPosition: &storepb.Position{
					Line: int32(table.line),
				},
			})
		}
		if table.onUpdateCurrentTimeCount > maxOnUpdateCurrentTimeColumnCount {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.OnUpdateCurrentTimeColumnCountExceedsLimit.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("Table `%s` has %d ON UPDATE CURRENT_TIMESTAMP() columns. The count greater than %d.", table.tableName, table.onUpdateCurrentTimeCount, maxOnUpdateCurrentTimeColumnCount),
				StartPosition: &storepb.Position{
					Line: int32(table.line),
				},
			})
		}
	}

	return checker.adviceList
}
