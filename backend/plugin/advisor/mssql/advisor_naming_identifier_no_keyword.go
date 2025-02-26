// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*NamingIdentifierNoKeywordAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLIdentifierNamingNoKeyword, &NamingIdentifierNoKeywordAdvisor{})
}

// NamingIdentifierNoKeywordAdvisor is the advisor checking for identifier naming convention without keyword..
type NamingIdentifierNoKeywordAdvisor struct {
}

// Check checks for identifier naming convention without keyword..
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
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingIdentifierNoKeywordChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterId_ is called when production id_ is entered.
func (l *namingIdentifierNoKeywordChecker) EnterId_(ctx *parser.Id_Context) {
	if ctx == nil {
		return
	}

	parent := ctx.GetParent()
	switch parent.(type) {
	case *parser.Column_definitionContext:
	case *parser.Table_constraintContext:
	case *parser.Create_schemaContext:
	case *parser.Create_databaseContext:
	case *parser.Create_indexContext:
	case *parser.Table_nameContext:
	default:
		return
	}
	if ctx.GetText() == "" {
		return
	}

	_, normalizedID := tsqlparser.NormalizeTSQLIdentifier(ctx)
	if tsqlparser.IsTSQLReservedKeyword(normalizedID, false) {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.NameIsKeywordIdentifier.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Identifier [%s] is a keyword identifier and should be avoided.", normalizedID),
			StartPosition: &storepb.Position{
				Line: int32(ctx.GetStart().GetLine()),
			},
		})
	}
}
