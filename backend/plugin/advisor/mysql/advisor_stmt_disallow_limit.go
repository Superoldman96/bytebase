package mysql

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	mysql "github.com/bytebase/mysql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementDisallowLimitAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLStatementDisallowLimit, &StatementDisallowLimitAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLStatementDisallowLimit, &StatementDisallowLimitAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLStatementDisallowLimit, &StatementDisallowLimitAdvisor{})
}

// StatementDisallowLimitAdvisor is the advisor checking for no LIMIT clause in INSERT/UPDATE statement.
type StatementDisallowLimitAdvisor struct {
}

// Check checks for no LIMIT clause in INSERT/UPDATE statement.
func (*StatementDisallowLimitAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parser result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementDisallowLimitChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type statementDisallowLimitChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine     int
	isInsertStmt bool
	adviceList   []*storepb.Advice
	level        storepb.Advice_Status
	title        string
	text         string
	line         int
}

func (checker *statementDisallowLimitChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (checker *statementDisallowLimitChecker) EnterDeleteStatement(ctx *mysql.DeleteStatementContext) {
	if ctx.SimpleLimitClause() != nil && ctx.SimpleLimitClause().LIMIT_SYMBOL() != nil {
		checker.handleLimitClause(advisor.DeleteUseLimit, ctx.GetStart().GetLine())
	}
}

// EnterUpdateStatement is called when production updateStatement is entered.
func (checker *statementDisallowLimitChecker) EnterUpdateStatement(ctx *mysql.UpdateStatementContext) {
	if ctx.SimpleLimitClause() != nil && ctx.SimpleLimitClause().LIMIT_SYMBOL() != nil {
		checker.handleLimitClause(advisor.UpdateUseLimit, ctx.GetStart().GetLine())
	}
}

// EnterInsertStatement is called when production insertStatement is entered.
func (checker *statementDisallowLimitChecker) EnterInsertStatement(_ *mysql.InsertStatementContext) {
	checker.isInsertStmt = true
}

// ExitInsertStatement is called when production insertStatement is exited.
func (checker *statementDisallowLimitChecker) ExitInsertStatement(_ *mysql.InsertStatementContext) {
	checker.isInsertStmt = false
}

// EnterQueryExpression is called when production queryExpression is entered.
func (checker *statementDisallowLimitChecker) EnterQueryExpression(ctx *mysql.QueryExpressionContext) {
	if !checker.isInsertStmt {
		return
	}
	if ctx.LimitClause() != nil && ctx.LimitClause().LIMIT_SYMBOL() != nil {
		checker.handleLimitClause(advisor.InsertUseLimit, ctx.GetStart().GetLine())
	}
}

func (checker *statementDisallowLimitChecker) handleLimitClause(code advisor.Code, lineNumber int) {
	checker.adviceList = append(checker.adviceList, &storepb.Advice{
		Status:  checker.level,
		Code:    code.Int32(),
		Title:   checker.title,
		Content: fmt.Sprintf("LIMIT clause is forbidden in INSERT, UPDATE and DELETE statement, but \"%s\" uses", checker.text),
		StartPosition: &storepb.Position{
			Line: int32(checker.line + lineNumber),
		},
	})
}
