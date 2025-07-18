// Package oracle is the advisor for oracle database.
package oracle

import (
	"context"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/plsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
)

var (
	_ advisor.Advisor = (*SelectNoSelectAllAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_ORACLE, advisor.OracleNoSelectAll, &SelectNoSelectAllAdvisor{})
	advisor.Register(storepb.Engine_DM, advisor.OracleNoSelectAll, &SelectNoSelectAllAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE_ORACLE, advisor.OracleNoSelectAll, &SelectNoSelectAllAdvisor{})
}

// SelectNoSelectAllAdvisor is the advisor checking for no select all.
type SelectNoSelectAllAdvisor struct {
}

// Check checks for no select all.
func (*SelectNoSelectAllAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &selectNoSelectAllListener{
		level:           level,
		title:           string(checkCtx.Rule.Type),
		currentDatabase: checkCtx.CurrentDatabase,
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// selectNoSelectAllListener is the listener for no select all.
type selectNoSelectAllListener struct {
	*parser.BasePlSqlParserListener

	level           storepb.Advice_Status
	title           string
	currentDatabase string
	adviceList      []*storepb.Advice
}

func (l *selectNoSelectAllListener) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterSelected_list is called when production selected_list is entered.
func (l *selectNoSelectAllListener) EnterSelected_list(ctx *parser.Selected_listContext) {
	if ctx.ASTERISK() != nil {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.StatementSelectAll.Int32(),
			Title:         l.title,
			Content:       "Avoid using SELECT *.",
			StartPosition: common.ConvertANTLRLineToPosition(ctx.GetStart().GetLine()),
		})
	}
}
