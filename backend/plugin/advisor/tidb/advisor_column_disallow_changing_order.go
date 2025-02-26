package tidb

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnDisallowChangingOrderAdvisor)(nil)
	_ ast.Visitor     = (*columnDisallowChangingOrderChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLColumnDisallowChangingOrder, &ColumnDisallowChangingOrderAdvisor{})
}

// ColumnDisallowChangingOrderAdvisor is the advisor checking for disallow changing column order.
type ColumnDisallowChangingOrderAdvisor struct {
}

// Check checks for disallow changing column order.
func (*ColumnDisallowChangingOrderAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnDisallowChangingOrderChecker{
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

type columnDisallowChangingOrderChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
}

// Enter implements the ast.Visitor interface.
func (checker *columnDisallowChangingOrderChecker) Enter(in ast.Node) (ast.Node, bool) {
	if node, ok := in.(*ast.AlterTableStmt); ok {
		for _, spec := range node.Specs {
			if (spec.Tp == ast.AlterTableChangeColumn || spec.Tp == ast.AlterTableModifyColumn) &&
				spec.Position.Tp != ast.ColumnPositionNone {
				checker.adviceList = append(checker.adviceList, &storepb.Advice{
					Status:  checker.level,
					Code:    advisor.ChangeColumnOrder.Int32(),
					Title:   checker.title,
					Content: fmt.Sprintf("\"%s\" changes column order", checker.text),
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
func (*columnDisallowChangingOrderChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
