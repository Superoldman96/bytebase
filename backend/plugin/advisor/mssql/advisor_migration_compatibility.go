// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*MigrationCompatibilityAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLMigrationCompatibility, &MigrationCompatibilityAdvisor{})
}

// MigrationCompatibilityAdvisor is the advisor checking for migration compatibility..
type MigrationCompatibilityAdvisor struct {
}

// Check checks for migration compatibility..
func (*MigrationCompatibilityAdvisor) Check(ctx advisor.Context) ([]*storepb.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &migrationCompatibilityChecker{
		level:                              level,
		title:                              string(ctx.Rule.Type),
		currentDatabase:                    ctx.CurrentDatabase,
		normalizedNewCreateTableNameMap:    make(map[string]any),
		normalizedNewCreateSchemaNameMap:   make(map[string]any),
		normalizedNewCreateDatabaseNameMap: make(map[string]any),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// migrationCompatibilityChecker is the listener for migration compatibility.
type migrationCompatibilityChecker struct {
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	// normalizedLastCreateTableNameMap contain the last created table name in normalized format.
	normalizedNewCreateTableNameMap map[string]any
	// normalizedLastCreateSchemaNameMap contain the last created schema name in normalized format.
	normalizedNewCreateSchemaNameMap map[string]any
	// normalizedNewCreateDatabaseNameMap contain the new created database name in normalized format.
	normalizedNewCreateDatabaseNameMap map[string]any

	// currentDatabase is the current database name.
	currentDatabase string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *migrationCompatibilityChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

func (l *migrationCompatibilityChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	tableName := ctx.Table_name()
	if tableName == nil || tableName.GetTable() == nil {
		return
	}
	normalizedTableName := tsqlparser.NormalizeTSQLTableName(tableName, l.currentDatabase, "dbo", false)
	l.normalizedNewCreateTableNameMap[normalizedTableName] = any(nil)
}

func (l *migrationCompatibilityChecker) EnterCreate_schema(ctx *parser.Create_schemaContext) {
	var schemaName string
	if v := ctx.GetSchema_name(); v != nil {
		_, schemaName = tsqlparser.NormalizeTSQLIdentifier(v)
	} else {
		_, schemaName = tsqlparser.NormalizeTSQLIdentifier(ctx.GetOwner_name())
	}

	normalizedDatabaseSchemaName := fmt.Sprintf("%s.%s", l.currentDatabase, schemaName)
	l.normalizedNewCreateSchemaNameMap[normalizedDatabaseSchemaName] = any(nil)
}

func (l *migrationCompatibilityChecker) EnterCreate_database(ctx *parser.Create_databaseContext) {
	_, databaseName := tsqlparser.NormalizeTSQLIdentifier(ctx.GetDatabase())
	l.normalizedNewCreateDatabaseNameMap[databaseName] = any(nil)
}

func (l *migrationCompatibilityChecker) EnterDrop_table(ctx *parser.Drop_tableContext) {
	allTableNames := ctx.AllTable_name()
	for _, tableName := range allTableNames {
		if tableName == nil || tableName.GetTable() == nil {
			continue
		}
		normalizedTableName := tsqlparser.NormalizeTSQLTableName(tableName, l.currentDatabase, "dbo", false)
		if _, ok := l.normalizedNewCreateTableNameMap[normalizedTableName]; !ok {
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    advisor.CompatibilityDropSchema.Int32(),
				Title:   l.title,
				Content: fmt.Sprintf("Drop table %s may cause incompatibility with the existing data and code", normalizedTableName),
				StartPosition: &storepb.Position{
					Line: int32(ctx.GetStart().GetLine()),
				},
			})
		}
		delete(l.normalizedNewCreateTableNameMap, normalizedTableName)
	}
}

func (l *migrationCompatibilityChecker) EnterDrop_schema(ctx *parser.Drop_schemaContext) {
	schemaName := ctx.GetSchema_name()
	_, normalizedSchemaName := tsqlparser.NormalizeTSQLIdentifier(schemaName)
	normalizedSchemaName = fmt.Sprintf("%s.%s", l.currentDatabase, normalizedSchemaName)
	if _, ok := l.normalizedNewCreateSchemaNameMap[normalizedSchemaName]; !ok {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.CompatibilityDropSchema.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Drop schema %s may cause incompatibility with the existing data and code", normalizedSchemaName),
			StartPosition: &storepb.Position{
				Line: int32(ctx.GetStart().GetLine()),
			},
		})
	}
	delete(l.normalizedNewCreateSchemaNameMap, normalizedSchemaName)
}

func (l *migrationCompatibilityChecker) EnterDrop_database(ctx *parser.Drop_databaseContext) {
	databaseName := ctx.GetDatabase_name_or_database_snapshot_name()
	_, normalizedDatabaseName := tsqlparser.NormalizeTSQLIdentifier(databaseName)
	if _, ok := l.normalizedNewCreateDatabaseNameMap[normalizedDatabaseName]; !ok {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.CompatibilityDropSchema.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Drop database %s may cause incompatibility with the existing data and code", normalizedDatabaseName),
			StartPosition: &storepb.Position{
				Line: int32(ctx.GetStart().GetLine()),
			},
		})
	}
	delete(l.normalizedNewCreateDatabaseNameMap, normalizedDatabaseName)
}

func (l *migrationCompatibilityChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	handleTableName := ctx.Table_name(0)
	normalizedHandleTableName := tsqlparser.NormalizeTSQLTableName(handleTableName, l.currentDatabase, "dbo", false)
	if _, ok := l.normalizedNewCreateTableNameMap[normalizedHandleTableName]; ok {
		return
	}

	if ctx.DROP() != nil && ctx.COLUMN() != nil {
		allDropColumns := ctx.AllId_()
		var allNormalizedDropColumnNames []string
		for _, dropColumn := range allDropColumns {
			_, normalizedDropColumnName := tsqlparser.NormalizeTSQLIdentifier(dropColumn)
			allNormalizedDropColumnNames = append(allNormalizedDropColumnNames, normalizedDropColumnName)
		}
		placeholder := strings.Join(allNormalizedDropColumnNames, ", ")
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.CompatibilityDropSchema.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Drop column %s may cause incompatibility with the existing data and code", placeholder),
			StartPosition: &storepb.Position{
				Line: int32(ctx.COLUMN().GetSymbol().GetLine()),
			},
		})
		return
	}
	if len(ctx.AllALTER()) == 2 && ctx.COLUMN() != nil {
		normalizedColumnName := ""
		if ctx.Column_definition() != nil {
			_, normalizedColumnName = tsqlparser.NormalizeTSQLIdentifier(ctx.Column_definition().Id_())
		} else if ctx.Column_modifier() != nil {
			_, normalizedColumnName = tsqlparser.NormalizeTSQLIdentifier(ctx.Column_modifier().Id_())
		}

		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.CompatibilityAlterColumn.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Alter COLUMN %s may cause incompatibility with the existing data and code", normalizedColumnName),
			StartPosition: &storepb.Position{
				Line: int32(ctx.COLUMN().GetSymbol().GetLine()),
			},
		})
		return
	}
	if v := ctx.Column_def_table_constraints(); v != nil {
		allColumnDefTableConstraints := v.AllColumn_def_table_constraint()
		for _, columnDefTableConstraint := range allColumnDefTableConstraints {
			code := advisor.Ok
			operation := ""
			tableConstraint := columnDefTableConstraint.Table_constraint()
			if tableConstraint == nil {
				continue
			}
			if tableConstraint.PRIMARY() != nil {
				code = advisor.CompatibilityAddPrimaryKey
				operation = "Add PRIMARY KEY"
			}
			if tableConstraint.UNIQUE() != nil {
				code = advisor.CompatibilityAddUniqueKey
				operation = "Add UNIQUE KEY"
			}
			if tableConstraint.Check_constraint() != nil {
				code = advisor.CompatibilityAddCheck
				operation = "Add CHECK"
			}
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    code.Int32(),
				Title:   l.title,
				Content: fmt.Sprintf("%s may cause incompatibility with the existing data and code", operation),
				StartPosition: &storepb.Position{
					Line: int32(ctx.GetStart().GetLine()),
				},
			})
		}
		return
	}
	if ctx.WITH() != nil && ctx.NOCHECK() != nil {
		if ctx.FOREIGN() != nil {
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    advisor.CompatibilityAddForeignKey.Int32(),
				Title:   l.title,
				Content: "Add FOREIGN KEY WITH NO CHECK may cause incompatibility with the existing data and code",
				StartPosition: &storepb.Position{
					Line: int32(ctx.FOREIGN().GetSymbol().GetLine()),
				},
			})
			return
		}
		if len(ctx.AllCHECK()) == 1 {
			l.adviceList = append(l.adviceList, &storepb.Advice{
				Status:  l.level,
				Code:    advisor.CompatibilityAddForeignKey.Int32(),
				Title:   l.title,
				Content: "Add CHECK WITH NO CHECK may cause incompatibility with the existing data and code",
				StartPosition: &storepb.Position{
					Line: int32(ctx.CHECK(0).GetSymbol().GetLine()),
				},
			})
			return
		}
	}
}

// EnterExecute_body is called when production execute_body is entered.
func (l *migrationCompatibilityChecker) EnterExecute_body(ctx *parser.Execute_bodyContext) {
	if ctx.Func_proc_name_server_database_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema().GetSchema() != nil {
		return
	}

	v := ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema().GetProcedure()
	_, normalizedProcedureName := tsqlparser.NormalizeTSQLIdentifier(v)
	if normalizedProcedureName != "sp_rename" {
		return
	}

	unnamedArguments := tsqlparser.FlattenExecuteStatementArgExecuteStatementArgUnnamed(ctx.Execute_statement_arg())

	firstArgument := unnamedArguments[0]
	if firstArgument == nil {
		return
	}
	if firstArgument.Execute_parameter() == nil {
		return
	}
	if firstArgument.Execute_parameter().Constant() == nil {
		return
	}
	if firstArgument.Execute_parameter().Constant().STRING() == nil {
		return
	}
	l.adviceList = append(l.adviceList, &storepb.Advice{
		Status:  l.level,
		Code:    advisor.CompatibilityRenameTable.Int32(),
		Title:   l.title,
		Content: "sp_rename may cause incompatibility with the existing data and code, and break scripts and stored procedures.",
		StartPosition: &storepb.Position{
			Line: int32(ctx.GetStart().GetLine()),
		},
	})
}
