package tidb

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnDisallowDropInIndexAdvisor)(nil)
	_ ast.Visitor     = (*columnDisallowDropInIndexChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLColumnDisallowDropInIndex, &ColumnDisallowDropInIndexAdvisor{})
}

// ColumnDisallowDropInIndexAdvisor is the advisor checking for disallow DROP COLUMN in index.
type ColumnDisallowDropInIndexAdvisor struct {
}

// Check checks for disallow Drop COLUMN in index statement.
func (*ColumnDisallowDropInIndexAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	checker := &columnDisallowDropInIndexChecker{
		level:   level,
		title:   string(checkCtx.Rule.Type),
		tables:  make(tableState),
		catalog: checkCtx.Catalog,
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.adviceList, nil
}

type columnDisallowDropInIndexChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	tables     tableState // the variable mean whether the column in index.
	catalog    *catalog.Finder
	line       int
}

func (checker *columnDisallowDropInIndexChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		checker.addIndexColumn(node)
	case *ast.AlterTableStmt:
		return checker.dropColumn(node)
	}
	return in, false
}

func (*columnDisallowDropInIndexChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (checker *columnDisallowDropInIndexChecker) dropColumn(in ast.Node) (ast.Node, bool) {
	if node, ok := in.(*ast.AlterTableStmt); ok {
		for _, spec := range node.Specs {
			if spec.Tp == ast.AlterTableDropColumn {
				table := node.Table.Name.O

				index := checker.catalog.Origin.Index(&catalog.TableIndexFind{
					// In MySQL, the SchemaName is "".
					SchemaName: "",
					TableName:  table,
				})

				if index != nil {
					if checker.tables[table] == nil {
						checker.tables[table] = make(columnSet)
					}
					for _, indexColumn := range *index {
						for _, column := range indexColumn.ExpressionList() {
							checker.tables[table][column] = true
						}
					}
				}

				colName := spec.OldColumnName.Name.String()
				if !checker.canDrop(table, colName) {
					checker.adviceList = append(checker.adviceList, &storepb.Advice{
						Status:  checker.level,
						Code:    advisor.DropIndexColumn.Int32(),
						Title:   checker.title,
						Content: fmt.Sprintf("`%s`.`%s` cannot drop index column", table, colName),
						StartPosition: &storepb.Position{
							Line: int32(checker.line),
						},
					})
				}
			}
		}
	}
	return in, false
}

func (checker *columnDisallowDropInIndexChecker) addIndexColumn(in ast.Node) {
	if node, ok := in.(*ast.CreateTableStmt); ok {
		for _, spec := range node.Constraints {
			if spec.Tp == ast.ConstraintIndex {
				for _, key := range spec.Keys {
					table := node.Table.Name.O
					if checker.tables[table] == nil {
						checker.tables[table] = make(columnSet)
					}
					checker.tables[table][key.Column.Name.O] = true
				}
			}
		}
	}
}

func (checker *columnDisallowDropInIndexChecker) canDrop(table string, column string) bool {
	if _, ok := checker.tables[table][column]; ok {
		return false
	}
	return true
}
