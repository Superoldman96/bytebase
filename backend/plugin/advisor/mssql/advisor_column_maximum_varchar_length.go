// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/common/log"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
)

var (
	_ advisor.Advisor = (*ColumnMaximumVarcharLengthAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLColumnMaximumVarcharLength, &ColumnMaximumVarcharLengthAdvisor{})
}

// ColumnMaximumVarcharLengthAdvisor is the advisor checking for maximum varchar length..
type ColumnMaximumVarcharLengthAdvisor struct {
}

// Check checks for maximum varchar length.
func (*ColumnMaximumVarcharLengthAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}

	listener := &columnMaximumVarcharLengthChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
		checkTypeString: map[string]any{
			"varchar":  nil,
			"nvarchar": nil,
			"char":     nil,
			"nchar":    nil,
		},
		maximum: payload.Number,
	}

	if listener.maximum > 0 {
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	}

	return listener.generateAdvice()
}

// columnMaximumVarcharLengthChecker is the listener for maximum varchar length.
type columnMaximumVarcharLengthChecker struct {
	*parser.BaseTSqlParserListener

	level           storepb.Advice_Status
	title           string
	checkTypeString map[string]any
	maximum         int

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *columnMaximumVarcharLengthChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

func (l *columnMaximumVarcharLengthChecker) EnterData_type(ctx *parser.Data_typeContext) {
	currentLength := 0
	line := ctx.GetStart().GetLine()
	if ctx.MAX() != nil && (ctx.VARCHAR() != nil || ctx.NVARCHAR() != nil) {
		// https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16&redirectedfrom=MSDN
		currentLength = math.MaxInt32 // 2 ^ 31 - 1
		line = ctx.MAX().GetSymbol().GetLine()
	} else if ctx.GetExt_type() != nil && ctx.GetScale() != nil && ctx.GetPrec() == nil && ctx.GetInc() == nil {
		_, normalizedTypeString := tsqlparser.NormalizeTSQLIdentifier(ctx.GetExt_type())
		if _, ok := l.checkTypeString[normalizedTypeString]; !ok {
			return
		}
		length, err := strconv.Atoi(ctx.GetScale().GetText())
		if err != nil {
			slog.Error("failed to convert scale to int", log.BBError(err))
		}
		currentLength = length
		line = ctx.GetScale().GetLine()
	} else if ctx.GetUnscaled_type() != nil {
		_, normalizedTypeString := tsqlparser.NormalizeTSQLIdentifier(ctx.GetUnscaled_type())
		if _, ok := l.checkTypeString[normalizedTypeString]; !ok {
			return
		}
		line = ctx.GetUnscaled_type().GetStart().GetLine()
	}
	if currentLength > l.maximum {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:        l.level,
			Code:          advisor.VarcharLengthExceedsLimit.Int32(),
			Title:         l.title,
			Content:       fmt.Sprintf("The maximum varchar length is %d.", l.maximum),
			StartPosition: common.ConvertANTLRLineToPosition(line),
		})
	}
}
