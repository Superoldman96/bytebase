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
	_ advisor.Advisor = (*TableRequirePkAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLTableRequirePK, &TableRequirePkAdvisor{})
}

// TableRequirePkAdvisor is the advisor checking for table require primary key..
type TableRequirePkAdvisor struct {
}

// Check checks for table require primary key..
func (*TableRequirePkAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := checkCtx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &tableRequirePkChecker{
		level:                      level,
		title:                      string(checkCtx.Rule.Type),
		currentNormalizedTableName: "",
		currentConstraintAction:    currentConstraintActionNone,
		tableHasPrimaryKey:         make(map[string]bool),
		tableOriginalName:          make(map[string]string),
		tableLine:                  make(map[string]int),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// tableRequirePkChecker is the listener for table require primary key.
type tableRequirePkChecker struct {
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	currentNormalizedTableName string
	currentConstraintAction    currentConstraintAction

	// tableHasPrimaryKey is a map from normalized table name to whether the table has primary key.
	tableHasPrimaryKey map[string]bool
	// tableOriginalName is a map from normalized table name to the original table name.
	tableOriginalName map[string]string
	// tableLine is a map from normalized table name to the line number of the table.
	tableLine map[string]int

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *tableRequirePkChecker) generateAdvice() ([]*storepb.Advice, error) {
	for tableName, hasPK := range l.tableHasPrimaryKey {
		if !hasPK {
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    advisor.TableNoPK.Int32(),
				Title:   l.title,
				Content: fmt.Sprintf("Table %s requires PRIMARY KEY.", l.tableOriginalName[tableName]),
				StartPosition: &storepb.Position{
					Line: int32(l.tableLine[tableName]),
				},
			})
		}
	}

	return l.adviceList, nil
}

// EnterCreate_table will be called when entering "CREATE TABLE" statement.
func (l *tableRequirePkChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	tableName := ctx.Table_name()
	if tableName == nil {
		return
	}
	normalizedTableName := tsqlparser.NormalizeTSQLTableName(tableName, "" /* fallbackDatabase */, "dbo" /* fallbackSchema */, false /* caseSensitive */)

	l.tableHasPrimaryKey[normalizedTableName] = false
	l.tableOriginalName[normalizedTableName] = tableName.GetText()
	l.tableLine[normalizedTableName] = tableName.GetStart().GetLine()

	l.currentNormalizedTableName = normalizedTableName
	l.currentConstraintAction = currentConstraintActionAdd
}

func (l *tableRequirePkChecker) ExitCreate_table(*parser.Create_tableContext) {
	l.currentNormalizedTableName = ""
	l.currentConstraintAction = currentConstraintActionNone
}

// EnterColumn_def_table_constraints will be called when entering "column_def_table_constraints" rule.
func (l *tableRequirePkChecker) EnterColumn_def_table_constraints(ctx *parser.Column_def_table_constraintsContext) {
	if l.currentNormalizedTableName == "" {
		return
	}

	allColumnDefTableConstraints := ctx.AllColumn_def_table_constraint()
	for _, columnDefTableConstraint := range allColumnDefTableConstraints {
		if v := columnDefTableConstraint.Column_definition(); v != nil {
			allColumnDefinitionElements := v.AllColumn_definition_element()
			for _, columnDefinitionElement := range allColumnDefinitionElements {
				if v := columnDefinitionElement.Column_constraint(); v != nil {
					if v.PRIMARY() != nil {
						if l.currentConstraintAction == currentConstraintActionAdd {
							l.tableHasPrimaryKey[l.currentNormalizedTableName] = true
						}
						return
					}
				}
			}
		} else if v := columnDefTableConstraint.Table_constraint(); v != nil {
			if v.PRIMARY() != nil {
				if l.currentConstraintAction == currentConstraintActionAdd {
					l.tableHasPrimaryKey[l.currentNormalizedTableName] = true
				}
				return
			}
		}
	}
}

func (l *tableRequirePkChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	tableName := ctx.Table_name(0)
	if tableName == nil {
		return
	}
	normalizedTableName := tsqlparser.NormalizeTSQLTableName(tableName, "" /* fallbackDatabase */, "dbo" /* fallbackSchema */, false /* caseSensitive */)
	if ctx.ADD() != nil && ctx.Column_def_table_constraints() != nil {
		l.currentNormalizedTableName = normalizedTableName
		l.currentConstraintAction = currentConstraintActionAdd
	} else if ctx.DROP() != nil && ctx.CONSTRAINT() != nil && ctx.GetConstraint() != nil {
		l.currentNormalizedTableName = normalizedTableName
		l.currentConstraintAction = currentConstraintActionDrop
	}
}

func (l *tableRequirePkChecker) ExitAlter_table(*parser.Alter_tableContext) {
	l.currentNormalizedTableName = ""
	l.currentConstraintAction = currentConstraintActionNone
}
