package mysql

// Framework code is generated by the generator.

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*IndexPrimaryKeyTypeAllowlistAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLPrimaryKeyTypeAllowlist, &IndexPrimaryKeyTypeAllowlistAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLPrimaryKeyTypeAllowlist, &IndexPrimaryKeyTypeAllowlistAdvisor{})
}

// IndexPrimaryKeyTypeAllowlistAdvisor is the advisor checking for primary key type allowlist.
type IndexPrimaryKeyTypeAllowlistAdvisor struct {
}

// Check checks for primary key type allowlist.
func (*IndexPrimaryKeyTypeAllowlistAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalStringArrayTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	allowlist := make(map[string]bool)
	for _, tp := range payload.List {
		allowlist[strings.ToLower(tp)] = true
	}
	checker := &indexPrimaryKeyTypeAllowlistChecker{
		level:            level,
		title:            string(checkCtx.Rule.Type),
		allowlist:        allowlist,
		catalog:          checkCtx.Catalog,
		tablesNewColumns: make(tableColumnTypes),
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type indexPrimaryKeyTypeAllowlistChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine         int
	adviceList       []*storepb.Advice
	level            storepb.Advice_Status
	title            string
	allowlist        map[string]bool
	catalog          *catalog.Finder
	tablesNewColumns tableColumnTypes
}

func (checker *indexPrimaryKeyTypeAllowlistChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableName() == nil {
		return
	}
	if ctx.TableElementList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	for _, tableElement := range ctx.TableElementList().AllTableElement() {
		if tableElement == nil {
			continue
		}
		switch {
		case tableElement.ColumnDefinition() != nil:
			if tableElement.ColumnDefinition().FieldDefinition() == nil {
				continue
			}
			_, _, columnName := mysqlparser.NormalizeMySQLColumnName(tableElement.ColumnDefinition().ColumnName())
			checker.checkFieldDefinition(tableName, columnName, tableElement.ColumnDefinition().FieldDefinition())
		case tableElement.TableConstraintDef() != nil:
			checker.checkConstraintDef(tableName, tableElement.TableConstraintDef())
		}
	}
}

func (checker *indexPrimaryKeyTypeAllowlistChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
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
	if ctx.TableRef() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	for _, alterListItem := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if alterListItem == nil {
			continue
		}

		switch {
		// add column
		case alterListItem.ADD_SYMBOL() != nil && alterListItem.Identifier() != nil:
			switch {
			case alterListItem.Identifier() != nil && alterListItem.FieldDefinition() != nil:
				columnName := mysqlparser.NormalizeMySQLIdentifier(alterListItem.Identifier())
				checker.checkFieldDefinition(tableName, columnName, alterListItem.FieldDefinition())
			case alterListItem.OPEN_PAR_SYMBOL() != nil && alterListItem.TableElementList() != nil:
				for _, tableElement := range alterListItem.TableElementList().AllTableElement() {
					if tableElement.ColumnDefinition() == nil || tableElement.ColumnDefinition().ColumnName() == nil || tableElement.ColumnDefinition().FieldDefinition() == nil {
						continue
					}
					_, _, columnName := mysqlparser.NormalizeMySQLColumnName(tableElement.ColumnDefinition().ColumnName())
					checker.checkFieldDefinition(tableName, columnName, tableElement.ColumnDefinition().FieldDefinition())
				}
			}
		// modify column
		case alterListItem.MODIFY_SYMBOL() != nil && alterListItem.ColumnInternalRef() != nil:
			columnName := mysqlparser.NormalizeMySQLColumnInternalRef(alterListItem.ColumnInternalRef())
			checker.checkFieldDefinition(tableName, columnName, alterListItem.FieldDefinition())
		// change column
		case alterListItem.CHANGE_SYMBOL() != nil && alterListItem.ColumnInternalRef() != nil && alterListItem.Identifier() != nil:
			oldColumnName := mysqlparser.NormalizeMySQLColumnInternalRef(alterListItem.ColumnInternalRef())
			checker.tablesNewColumns.delete(tableName, oldColumnName)
			newColumnName := mysqlparser.NormalizeMySQLIdentifier(alterListItem.Identifier())
			checker.checkFieldDefinition(tableName, newColumnName, alterListItem.FieldDefinition())
		// add constriant.
		case alterListItem.ADD_SYMBOL() != nil && alterListItem.TableConstraintDef() != nil:
			checker.checkConstraintDef(tableName, alterListItem.TableConstraintDef())
		}
	}
}

func (checker *indexPrimaryKeyTypeAllowlistChecker) checkFieldDefinition(tableName, columnName string, ctx mysql.IFieldDefinitionContext) {
	if ctx.DataType() == nil {
		return
	}
	// columnType := ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx.DataType())
	// columnType = strings.ToLower(columnType)
	columnType := mysqlparser.NormalizeMySQLDataType(ctx.DataType(), true /* compact */)
	for _, attribute := range ctx.AllColumnAttribute() {
		if attribute.PRIMARY_SYMBOL() != nil {
			if _, exists := checker.allowlist[columnType]; !exists {
				checker.adviceList = append(checker.adviceList, &storepb.Advice{
					Status:        checker.level,
					Code:          advisor.IndexPKType.Int32(),
					Title:         checker.title,
					Content:       fmt.Sprintf("The column `%s` in table `%s` is one of the primary key, but its type \"%s\" is not in allowlist", columnName, tableName, columnType),
					StartPosition: advisor.ConvertANTLRLineToPosition(checker.baseLine + ctx.GetStart().GetLine()),
				})
			}
		}
	}
	checker.tablesNewColumns.set(tableName, columnName, columnType)
}

func (checker *indexPrimaryKeyTypeAllowlistChecker) checkConstraintDef(tableName string, ctx mysql.ITableConstraintDefContext) {
	if ctx.GetType_().GetTokenType() != mysql.MySQLParserPRIMARY_SYMBOL {
		return
	}
	if ctx.KeyListVariants() == nil {
		return
	}
	columnList := mysqlparser.NormalizeKeyListVariants(ctx.KeyListVariants())

	for _, columnName := range columnList {
		columnType, err := checker.getPKColumnType(tableName, columnName)
		if err != nil {
			continue
		}
		columnType = strings.ToLower(columnType)
		if _, exists := checker.allowlist[columnType]; !exists {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:        checker.level,
				Code:          advisor.IndexPKType.Int32(),
				Title:         checker.title,
				Content:       fmt.Sprintf("The column `%s` in table `%s` is one of the primary key, but its type \"%s\" is not in allowlist", columnName, tableName, columnType),
				StartPosition: advisor.ConvertANTLRLineToPosition(checker.baseLine + ctx.GetStart().GetLine()),
			})
		}
	}
}

// getPKColumnType gets the column type string from v.tablesNewColumns or catalog, returns empty string and non-nil error if cannot find the column in given table.
func (checker *indexPrimaryKeyTypeAllowlistChecker) getPKColumnType(tableName string, columnName string) (string, error) {
	if columnType, ok := checker.tablesNewColumns.get(tableName, columnName); ok {
		return columnType, nil
	}
	column := checker.catalog.Origin.FindColumn(&catalog.ColumnFind{
		TableName:  tableName,
		ColumnName: columnName,
	})
	if column != nil {
		return column.Type(), nil
	}
	return "", errors.Errorf("cannot find the type of `%s`.`%s`", tableName, columnName)
}
