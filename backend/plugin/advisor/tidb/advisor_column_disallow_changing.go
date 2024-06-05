package tidb

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnDisallowChangingAdvisor)(nil)
	_ ast.Visitor     = (*columnDisallowChangingChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLColumnDisallowChanging, &ColumnDisallowChangingAdvisor{})
}

// ColumnDisallowChangingAdvisor is the advisor checking for disallow CHANGE COLUMN statement.
type ColumnDisallowChangingAdvisor struct {
}

// Check checks for disallow CHANGE COLUMN statement.
func (*ColumnDisallowChangingAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnDisallowChangingChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  storepb.Advice_SUCCESS,
			Code:    advisor.Ok.Int32(),
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type columnDisallowChangingChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
}

// Enter implements the ast.Visitor interface.
func (checker *columnDisallowChangingChecker) Enter(in ast.Node) (ast.Node, bool) {
	if node, ok := in.(*ast.AlterTableStmt); ok {
		for _, spec := range node.Specs {
			if spec.Tp == ast.AlterTableChangeColumn {
				checker.adviceList = append(checker.adviceList, &storepb.Advice{
					Status:  checker.level,
					Code:    advisor.UseChangeColumnStatement.Int32(),
					Title:   checker.title,
					Content: fmt.Sprintf("\"%s\" contains CHANGE COLUMN statement", checker.text),
					StartPosition: &storepb.Position{
						Line: int32(checker.line),
					},
				})
				break
			}
		}
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnDisallowChangingChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
