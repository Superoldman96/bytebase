// Package snowflake is the advisor for snowflake database.
package snowflake

import (
	"context"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/snowsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*WhereRequireForSelectAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_SNOWFLAKE, advisor.SnowflakeWhereRequirementForSelect, &WhereRequireForSelectAdvisor{})
}

// WhereRequireForSelectAdvisor is the advisor checking for WHERE clause requirement for SELECT statement.
type WhereRequireForSelectAdvisor struct {
}

// Check checks for WHERE clause requirement.
func (*WhereRequireForSelectAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &whereRequireForSelectChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// whereRequireForSelectChecker is the listener for WHERE clause requirement.
type whereRequireForSelectChecker struct {
	*parser.BaseSnowflakeParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *whereRequireForSelectChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterQuery_statement is called when production query_statement is entered.
func (l *whereRequireForSelectChecker) EnterQuery_statement(ctx *parser.Query_statementContext) {
	if ctx.Select_statement() == nil {
		return
	}
	optional := ctx.Select_statement().Select_optional_clauses()
	if optional == nil {
		return
	}
	// Allow SELECT queries without a FROM clause to proceed, e.g. SELECT 1.
	if optional.Where_clause() == nil && optional.From_clause() != nil {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.StatementNoWhere.Int32(),
			Title:   l.title,
			Content: "WHERE clause is required for SELECT statement.",
			StartPosition: &storepb.Position{
				Line: int32(ctx.GetStart().GetLine()),
			},
		})
	}
}
