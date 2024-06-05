package mysql

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementMaximumLimitValueAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLStatementMaximumLimitValue, &StatementMaximumLimitValueAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLStatementMaximumLimitValue, &StatementMaximumLimitValueAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLStatementMaximumLimitValue, &StatementMaximumLimitValueAdvisor{})
}

type StatementMaximumLimitValueAdvisor struct {
}

func (*StatementMaximumLimitValueAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &statementMaximumLimitValueChecker{
		level:         level,
		title:         string(ctx.Rule.Type),
		limitMaxValue: payload.Number,
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
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

type statementMaximumLimitValueChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine      int
	adviceList    []*storepb.Advice
	level         storepb.Advice_Status
	title         string
	text          string
	isSelect      bool
	limitMaxValue int
}

func (checker *statementMaximumLimitValueChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}

func (checker *statementMaximumLimitValueChecker) EnterSelectStatement(ctx *mysql.SelectStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	checker.isSelect = true
}

func (checker *statementMaximumLimitValueChecker) ExitSelectStatement(ctx *mysql.SelectStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	checker.isSelect = false
}

func (checker *statementMaximumLimitValueChecker) EnterLimitClause(ctx *mysql.LimitClauseContext) {
	if !checker.isSelect {
		return
	}
	if ctx.LIMIT_SYMBOL() == nil {
		return
	}

	limitOptions := ctx.LimitOptions()
	for _, limitOption := range limitOptions.AllLimitOption() {
		limitValue, err := strconv.Atoi(limitOption.GetText())
		if err != nil {
			// Ignore invalid limit value and continue.
			continue
		}

		if limitValue > checker.limitMaxValue {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.StatementExceedMaximumLimitValue.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("The limit value %d exceeds the maximum allowed value %d", limitValue, checker.limitMaxValue),
				StartPosition: &storepb.Position{
					Line: int32(checker.baseLine + ctx.GetStart().GetLine()),
				},
			})
		}
	}
}
