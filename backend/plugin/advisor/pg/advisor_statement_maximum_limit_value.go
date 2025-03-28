package pg

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementMaximumLimitValueAdvisor)(nil)
	_ ast.Visitor     = (*statementMaximumLimitValueChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementMaximumLimitValue, &StatementMaximumLimitValueAdvisor{})
}

// StatementAddCheckNotValidAdvisor is the advisor checking for to add check not valid.
type StatementMaximumLimitValueAdvisor struct {
}

// Check checks for to add check not valid.
func (*StatementMaximumLimitValueAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &statementMaximumLimitValueChecker{
		level:         level,
		title:         string(checkCtx.Rule.Type),
		limitMaxValue: payload.Number,
	}

	for _, stmt := range stmtList {
		checker.line = stmt.LastLine()
		ast.Walk(checker, stmt)
	}

	return checker.adviceList, nil
}

type statementMaximumLimitValueChecker struct {
	adviceList    []*storepb.Advice
	level         storepb.Advice_Status
	title         string
	line          int
	limitMaxValue int
}

// Visit implements ast.Visitor interface.
func (checker *statementMaximumLimitValueChecker) Visit(in ast.Node) ast.Visitor {
	if node, ok := in.(*ast.SelectStmt); ok {
		if node.Limit != nil && int(*node.Limit) > checker.limitMaxValue {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:        checker.level,
				Code:          advisor.StatementExceedMaximumLimitValue.Int32(),
				Title:         checker.title,
				Content:       fmt.Sprintf("The limit value %d exceeds the maximum allowed value %d", *node.Limit, checker.limitMaxValue),
				StartPosition: advisor.ConvertANTLRLineToPosition(checker.line),
			})
		}
	}

	return checker
}
