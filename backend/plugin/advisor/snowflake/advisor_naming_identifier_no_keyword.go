// Package snowflake is the advisor for snowflake database.
package snowflake

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/snowsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	snowsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/snowflake"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*NamingIdentifierNoKeywordAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_SNOWFLAKE, advisor.SnowflakeIdentifierNamingNoKeyword, &NamingIdentifierNoKeywordAdvisor{})
}

// NamingIdentifierNoKeywordAdvisor is the advisor checking for identifier naming convention without keyword.
type NamingIdentifierNoKeywordAdvisor struct {
}

// Check checks for identifier naming convention without keyword.
func (*NamingIdentifierNoKeywordAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &namingIdentifierNoKeywordChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// namingIdentifierNoKeywordChecker is the listener for identifier naming convention without keyword.
type namingIdentifierNoKeywordChecker struct {
	*parser.BaseSnowflakeParserListener

	level storepb.Advice_Status
	title string

	// currentOriginalTableName is the original table name in the statement.
	currentOriginalTableName string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingIdentifierNoKeywordChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterCreate_table is called when production create_table is entered.
func (l *namingIdentifierNoKeywordChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	l.currentOriginalTableName = ctx.Object_name().GetText()
}

// ExitCreate_table is called when production create_table is exited.
func (l *namingIdentifierNoKeywordChecker) ExitCreate_table(*parser.Create_tableContext) {
	l.currentOriginalTableName = ""
}

// EnterCreate_table_as_select is called when production create_table_as_select is entered.
func (l *namingIdentifierNoKeywordChecker) EnterCreate_table_as_select(ctx *parser.Create_table_as_selectContext) {
	l.currentOriginalTableName = ctx.Object_name().GetText()
}

// ExitCreate_table_as_select is called when production create_table_as_select is exited.
func (l *namingIdentifierNoKeywordChecker) ExitCreate_table_as_select(*parser.Create_table_as_selectContext) {
	l.currentOriginalTableName = ""
}

// EnterColumn_decl_item_list is called when production column_decl_item_list is entered.
func (l *namingIdentifierNoKeywordChecker) EnterColumn_decl_item_list(ctx *parser.Column_decl_item_listContext) {
	if l.currentOriginalTableName == "" {
		return
	}

	allItems := ctx.AllColumn_decl_item()
	if len(allItems) == 0 {
		return
	}

	for _, item := range allItems {
		if fullColDecl := item.Full_col_decl(); fullColDecl != nil {
			originalID := fullColDecl.Col_decl().Column_name().Id_()
			originalColName := snowsqlparser.NormalizeSnowSQLObjectNamePart(originalID)
			if snowsqlparser.IsSnowflakeKeyword(originalColName, false) {
				l.adviceList = append(l.adviceList, &storepb.Advice{
					Status:        l.level,
					Code:          advisor.NameIsKeywordIdentifier.Int32(),
					Title:         l.title,
					Content:       fmt.Sprintf("Identifier %s is a keyword and should be avoided", originalID.GetText()),
					StartPosition: common.ConvertANTLRLineToPosition(ctx.GetStart().GetLine()),
				})
			}
		}
	}
}

// EnterAlter_table is called when production alter_table is entered.
func (l *namingIdentifierNoKeywordChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	if ctx.Table_column_action() == nil || ctx.Table_column_action().RENAME() == nil {
		return
	}
	l.currentOriginalTableName = ctx.Object_name(0).GetText()
	renameToID := ctx.Table_column_action().Column_name(1).Id_()
	renameToColName := snowsqlparser.NormalizeSnowSQLObjectNamePart(renameToID)
	if snowsqlparser.IsSnowflakeKeyword(renameToColName, false) {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.NameIsKeywordIdentifier.Int32(),
			Title:         l.title,
			Content:       fmt.Sprintf("Identifier %s is a keyword and should be avoided", renameToID.GetText()),
			StartPosition: common.ConvertANTLRLineToPosition(renameToID.GetStart().GetLine()),
		})
	}
}
