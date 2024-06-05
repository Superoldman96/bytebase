// Package snowflake is the advisor for snowflake database.
package snowflake

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/snowsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	snowsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/snowflake"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*TableNoForeignKeyAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_SNOWFLAKE, advisor.SnowflakeTableNoFK, &TableNoForeignKeyAdvisor{})
}

// TableNoForeignKeyAdvisor is the advisor checking for table disallow foreign key.
type TableNoForeignKeyAdvisor struct {
}

// Check checks for table disallow foreign key.
func (*TableNoForeignKeyAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &tableNoForeignKeyChecker{
		level:                      level,
		title:                      string(ctx.Rule.Type),
		currentConstraintAction:    currentConstraintActionNone,
		currentNormalizedTableName: "",
		tableForeignKeyTimes:       make(map[string]int),
		tableOriginalName:          make(map[string]string),
		tableLine:                  make(map[string]int),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// tableNoForeignKeyChecker is the listener for table disallow foreign key.
type tableNoForeignKeyChecker struct {
	*parser.BaseSnowflakeParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice

	currentConstraintAction currentConstraintAction
	// currentNormalizedTableName is the current table name, and it is normalized.
	// It should be set then entering create_table, alter_table and so on,
	// and should be reset then exiting them.
	currentNormalizedTableName string

	// tableForeignKeyTimes is a map of normalized table name to the times of FOREIGN KEY.
	tableForeignKeyTimes map[string]int
	// tableOriginalName is a map of normalized table name to original table name.
	// The key of the tableOriginalName is the superset of the key of the tableHasForeignKey.
	tableOriginalName map[string]string
	// tableLine is a map of normalized table name to the line number of the table.
	// The key of the tableLine is the superset of the key of the tableHasForeignKey.
	tableLine map[string]int
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *tableNoForeignKeyChecker) generateAdvice() ([]*storepb.Advice, error) {
	for tableName, times := range l.tableForeignKeyTimes {
		if times > 0 {
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    advisor.TableHasFK.Int32(),
				Title:   l.title,
				Content: fmt.Sprintf("FOREIGN KEY is not allowed in the table %s.", l.tableOriginalName[tableName]),
				StartPosition: &storepb.Position{
					Line: int32(l.tableLine[tableName]),
				},
			})
		}
	}
	if len(l.adviceList) == 0 {
		return []*storepb.Advice{
			{
				Status:  storepb.Advice_SUCCESS,
				Code:    advisor.Ok.Int32(),
				Title:   "OK",
				Content: "",
			},
		}, nil
	}
	return l.adviceList, nil
}

// EnterCreate_table is called when production create_table is entered.
func (l *tableNoForeignKeyChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	originalTableName := ctx.Object_name()
	normalizedTableName := snowsqlparser.NormalizeSnowSQLObjectName(originalTableName, "", "PUBLIC")

	l.tableForeignKeyTimes[normalizedTableName] = 0
	l.tableOriginalName[normalizedTableName] = originalTableName.GetText()
	l.tableLine[normalizedTableName] = ctx.GetStart().GetLine()
	l.currentNormalizedTableName = normalizedTableName
	l.currentConstraintAction = currentConstraintActionAdd
}

// ExitCreate_table is called when production create_table is exited.
func (l *tableNoForeignKeyChecker) ExitCreate_table(*parser.Create_tableContext) {
	l.currentNormalizedTableName = ""
	l.currentConstraintAction = currentConstraintActionNone
}

// EnterInline_constraint is called when production inline_constraint is entered.
func (l *tableNoForeignKeyChecker) EnterInline_constraint(ctx *parser.Inline_constraintContext) {
	if ctx.REFERENCES() == nil || l.currentNormalizedTableName == "" {
		return
	}
	l.tableForeignKeyTimes[l.currentNormalizedTableName]++
}

// EnterOut_of_line_constraint is called when production out_of_line_constraint is entered.
func (l *tableNoForeignKeyChecker) EnterOut_of_line_constraint(ctx *parser.Out_of_line_constraintContext) {
	if ctx.REFERENCES() == nil || l.currentNormalizedTableName == "" || l.currentConstraintAction == currentConstraintActionNone {
		return
	}
	if l.currentConstraintAction == currentConstraintActionAdd {
		l.tableForeignKeyTimes[l.currentNormalizedTableName]++
		l.tableLine[l.currentNormalizedTableName] = ctx.GetStart().GetLine()
	} else if l.currentConstraintAction == currentConstraintActionDrop {
		if times, ok := l.tableForeignKeyTimes[l.currentNormalizedTableName]; ok && times > 0 {
			l.tableForeignKeyTimes[l.currentNormalizedTableName]--
		}
	}
}

// EnterConstraint_action is called when production constraint_action is entered.
func (l *tableNoForeignKeyChecker) EnterConstraint_action(ctx *parser.Constraint_actionContext) {
	if l.currentNormalizedTableName == "" {
		return
	}
	if ctx.DROP() != nil && ctx.FOREIGN() != nil {
		if times, ok := l.tableForeignKeyTimes[l.currentNormalizedTableName]; ok && times > 0 {
			l.tableForeignKeyTimes[l.currentNormalizedTableName]--
		}
		return
	}
	if ctx.ADD() != nil {
		l.currentConstraintAction = currentConstraintActionAdd
		return
	}
}

// EnterAlter_table is called when production alter_table is entered.
func (l *tableNoForeignKeyChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	if ctx.Constraint_action() == nil {
		return
	}
	originalTableName := ctx.Object_name(0)
	normalizedTableName := snowsqlparser.NormalizeSnowSQLObjectName(originalTableName, "", "PUBLIC")

	l.currentNormalizedTableName = normalizedTableName
	l.tableOriginalName[normalizedTableName] = originalTableName.GetText()
}

// ExitAlter_table is called when production alter_table is exited.
func (l *tableNoForeignKeyChecker) ExitAlter_table(*parser.Alter_tableContext) {
	l.currentNormalizedTableName = ""
	l.currentConstraintAction = currentConstraintActionNone
}
