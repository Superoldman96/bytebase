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
	_ advisor.Advisor = (*WhereRequirementForUpdateDeleteAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLWhereRequirementForUpdateDelete, &WhereRequirementForUpdateDeleteAdvisor{})
}

// WhereRequirementForUpdateDeleteAdvisor is the advisor checking for WHERE clause requirement for UPDATE/DELETE statements.
type WhereRequirementForUpdateDeleteAdvisor struct {
}

// Check checks for WHERE clause requirement.
func (*WhereRequirementForUpdateDeleteAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &whereRequirementForUpdateDeleteChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// whereRequirementForUpdateDeleteChecker is the listener for WHERE clause requirement for SELECT statements.
type whereRequirementForUpdateDeleteChecker struct {
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *whereRequirementForUpdateDeleteChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterDelete_statement is called when production delete_statement is entered.
func (l *whereRequirementForUpdateDeleteChecker) EnterDelete_statement(ctx *parser.Delete_statementContext) {
	if ctx.WHERE() == nil {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.StatementNoWhere.Int32(),
			Title:         l.title,
			Content:       "WHERE clause is required for DELETE statement.",
			StartPosition: common.ConvertANTLRLineToPosition(ctx.GetStart().GetLine()),
		})
	}
}

// EnterUpdate_statement is called when production update_statement is entered.
func (l *whereRequirementForUpdateDeleteChecker) EnterUpdate_statement(ctx *parser.Update_statementContext) {
	if ctx.WHERE() == nil {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.StatementNoWhere.Int32(),
			Title:         l.title,
			Content:       "WHERE clause is required for UPDATE statement.",
			StartPosition: common.ConvertANTLRLineToPosition(ctx.GetStart().GetLine()),
		})
	}
}
