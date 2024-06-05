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
	_ advisor.Advisor = (*ColumnAutoIncrementInitialValueAdvisor)(nil)
	_ ast.Visitor     = (*columnAutoIncrementInitialValueChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLAutoIncrementColumnInitialValue, &ColumnAutoIncrementInitialValueAdvisor{})
}

// ColumnAutoIncrementInitialValueAdvisor is the advisor checking for auto-increment column initial value.
type ColumnAutoIncrementInitialValueAdvisor struct {
}

// Check checks for auto-increment column initial value.
func (*ColumnAutoIncrementInitialValueAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &columnAutoIncrementInitialValueChecker{
		level: level,
		title: string(ctx.Rule.Type),
		value: payload.Number,
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

type columnAutoIncrementInitialValueChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
	value      int
}

// Enter implements the ast.Visitor interface.
func (checker *columnAutoIncrementInitialValueChecker) Enter(in ast.Node) (ast.Node, bool) {
	if createTable, ok := in.(*ast.CreateTableStmt); ok {
		for _, option := range createTable.Options {
			if option.Tp == ast.TableOptionAutoIncrement {
				if option.UintValue != uint64(checker.value) {
					checker.adviceList = append(checker.adviceList, &storepb.Advice{
						Status:  checker.level,
						Code:    advisor.AutoIncrementColumnInitialValueNotMatch.Int32(),
						Title:   checker.title,
						Content: fmt.Sprintf("The initial auto-increment value in table `%s` is %v, which doesn't equal %v", createTable.Table.Name.O, option.UintValue, checker.value),
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

// Leave implements the ast.Visitor interface.
func (*columnAutoIncrementInitialValueChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
