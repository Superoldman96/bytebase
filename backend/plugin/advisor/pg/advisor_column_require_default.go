package pg

// Framework code is generated by the generator.

import (
	"context"
	"fmt"
	"sort"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnRequireDefaultAdvisor)(nil)
	_ ast.Visitor     = (*columnRequireDefaultChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLRequireColumnDefault, &ColumnRequireDefaultAdvisor{})
}

// ColumnRequireDefaultAdvisor is the advisor checking for column default requirement.
type ColumnRequireDefaultAdvisor struct {
}

// Check checks for column default requirement.
func (*ColumnRequireDefaultAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnRequireDefaultChecker{
		level:     level,
		title:     string(checkCtx.Rule.Type),
		catalog:   checkCtx.Catalog,
		columnSet: make(map[string]columnData),
	}

	if checker.catalog.Final.Usable() {
		for _, stmt := range stmtList {
			ast.Walk(checker, stmt)
		}
	}

	return checker.generateAdvice(), nil
}

type columnRequireDefaultChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	columnSet  map[string]columnData
	catalog    *catalog.Finder
}

type columnData struct {
	schema string
	table  string
	name   string
	line   int
}

func (checker *columnRequireDefaultChecker) generateAdvice() []*storepb.Advice {
	var columnList []columnData
	for _, column := range checker.columnSet {
		columnList = append(columnList, column)
	}
	sort.Slice(columnList, func(i, j int) bool {
		return columnList[i].line < columnList[j].line
	})

	for _, column := range columnList {
		columnInfo := checker.catalog.Final.FindColumn(&catalog.ColumnFind{
			SchemaName: column.schema,
			TableName:  column.table,
			ColumnName: column.name,
		})
		if columnInfo != nil && !columnInfo.HasDefault() {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:        checker.level,
				Code:          advisor.NoDefault.Int32(),
				Title:         checker.title,
				Content:       fmt.Sprintf("Column %q.%q in schema %q doesn't have DEFAULT", column.table, column.name, column.schema),
				StartPosition: advisor.ConvertANTLRLineToPosition(column.line),
			})
		}
	}

	return checker.adviceList
}

func (checker *columnRequireDefaultChecker) addColumn(schema string, table string, column string, line int) {
	if schema == "" {
		schema = "public"
	}

	checker.columnSet[fmt.Sprintf("%s.%s.%s", schema, table, column)] = columnData{
		schema: schema,
		table:  table,
		name:   column,
		line:   line,
	}
}

// Visit implements ast.Visitor interface.
func (checker *columnRequireDefaultChecker) Visit(in ast.Node) ast.Visitor {
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.ColumnList {
			checker.addColumn(node.Name.Schema, node.Name.Name, column.ColumnName, column.LastLine())
		}
	case *ast.AlterTableStmt:
		for _, item := range node.AlterItemList {
			if addColumn, ok := item.(*ast.AddColumnListStmt); ok {
				for _, column := range addColumn.ColumnList {
					checker.addColumn(node.Table.Schema, node.Table.Name, column.ColumnName, node.LastLine())
				}
			}
		}
	}

	return checker
}
