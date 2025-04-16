// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"context"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*WhereRequirementForSelectAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLWhereRequirementForSelect, &WhereRequirementForSelectAdvisor{})
}

// WhereRequirementForSelectAdvisor is the advisor checking for WHERE clause requirement for SELECT statements.
type WhereRequirementForSelectAdvisor struct {
}

// Check checks for WHERE clause requirement.
func (*WhereRequirementForSelectAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &whereRequirementForSelectChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// whereRequirementForSelectChecker is the listener for WHERE clause requirement for SELECT statements.
type whereRequirementForSelectChecker struct {
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *whereRequirementForSelectChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterQuery_Specification is called when production query_Specification is entered.
func (l *whereRequirementForSelectChecker) EnterQuery_specification(ctx *parser.Query_specificationContext) {
	// Allow SELECT queries without a FROM clause to proceed, e.g. SELECT 1.
	if ctx.WHERE() == nil && ctx.From_table_sources() != nil {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.StatementNoWhere.Int32(),
			Title:         l.title,
			Content:       "WHERE clause is required for SELETE statement.",
			StartPosition: common.ConvertANTLRLineToPosition(ctx.GetStart().GetLine()),
		})
	}
}
