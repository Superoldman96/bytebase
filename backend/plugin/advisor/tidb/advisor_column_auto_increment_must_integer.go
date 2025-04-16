package tidb

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/mysql"
	"github.com/pingcap/tidb/pkg/parser/types"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnAutoIncrementMustIntegerAdvisor)(nil)
	_ ast.Visitor     = (*columnAutoIncrementMustIntegerChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLAutoIncrementColumnMustInteger, &ColumnAutoIncrementMustIntegerAdvisor{})
}

// ColumnAutoIncrementMustIntegerAdvisor is the advisor checking for auto-increment column type.
type ColumnAutoIncrementMustIntegerAdvisor struct {
}

// Check checks for auto-increment column type.
func (*ColumnAutoIncrementMustIntegerAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnAutoIncrementMustIntegerChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.adviceList, nil
}

type columnAutoIncrementMustIntegerChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
}

type columnData struct {
	table  string
	column string
	line   int
}

// Enter implements the ast.Visitor interface.
func (checker *columnAutoIncrementMustIntegerChecker) Enter(in ast.Node) (ast.Node, bool) {
	var columnList []columnData
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.Cols {
			if !autoIncrementColumnIsInteger(column) {
				columnList = append(columnList, columnData{
					table:  node.Table.Name.O,
					column: column.Name.Name.O,
					line:   column.OriginTextPosition(),
				})
			}
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					if !autoIncrementColumnIsInteger(column) {
						columnList = append(columnList, columnData{
							table:  node.Table.Name.O,
							column: column.Name.Name.O,
							line:   node.OriginTextPosition(),
						})
					}
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				if !autoIncrementColumnIsInteger(spec.NewColumns[0]) {
					columnList = append(columnList, columnData{
						table:  node.Table.Name.O,
						column: spec.NewColumns[0].Name.Name.O,
						line:   node.OriginTextPosition(),
					})
				}
			}
		}
	}

	for _, column := range columnList {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.AutoIncrementColumnNotInteger.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("Auto-increment column `%s`.`%s` requires integer type", column.table, column.column),
			StartPosition: common.ConvertANTLRLineToPosition(checker.line),
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnAutoIncrementMustIntegerChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func autoIncrementColumnIsInteger(column *ast.ColumnDef) bool {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionAutoIncrement && !isInteger(column.Tp) {
			return false
		}
	}
	return true
}

func isInteger(tp *types.FieldType) bool {
	switch tp.GetType() {
	case mysql.TypeTiny, mysql.TypeShort, mysql.TypeInt24, mysql.TypeLong, mysql.TypeLonglong:
		return true
	default:
		return false
	}
}
