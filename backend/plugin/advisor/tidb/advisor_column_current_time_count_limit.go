package tidb

// Framework code is generated by the generator.

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/mysql"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

const (
	maxDefaultCurrentTimeColumCount   = 2
	maxOnUpdateCurrentTimeColumnCount = 1
)

var (
	_ advisor.Advisor = (*ColumnCurrentTimeCountLimitAdvisor)(nil)
	_ ast.Visitor     = (*columnCurrentTimeCountLimitChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLCurrentTimeColumnCountLimit, &ColumnCurrentTimeCountLimitAdvisor{})
}

// ColumnCurrentTimeCountLimitAdvisor is the advisor checking for current time column count limit.
type ColumnCurrentTimeCountLimitAdvisor struct {
}

// Check checks for current time column count limit.
func (*ColumnCurrentTimeCountLimitAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnCurrentTimeCountLimitChecker{
		level:    level,
		title:    string(ctx.Rule.Type),
		tableSet: make(map[string]tableData),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.generateAdvice(), nil
}

type columnCurrentTimeCountLimitChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
	tableSet   map[string]tableData
}

type tableData struct {
	tableName                string
	defaultCurrentTimeCount  int
	onUpdateCurrentTimeCount int
	line                     int
}

func (checker *columnCurrentTimeCountLimitChecker) generateAdvice() []*storepb.Advice {
	var tableList []tableData
	for _, table := range checker.tableSet {
		tableList = append(tableList, table)
	}
	sort.Slice(tableList, func(i, j int) bool {
		return tableList[i].line < tableList[j].line
	})
	for _, table := range tableList {
		if table.defaultCurrentTimeCount > maxDefaultCurrentTimeColumCount {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.DefaultCurrentTimeColumnCountExceedsLimit.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("Table `%s` has %d DEFAULT CURRENT_TIMESTAMP() columns. The count greater than %d.", table.tableName, table.defaultCurrentTimeCount, maxDefaultCurrentTimeColumCount),
				StartPosition: &storepb.Position{
					Line: int32(table.line),
				},
			})
		}
		if table.onUpdateCurrentTimeCount > maxOnUpdateCurrentTimeColumnCount {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.OnUpdateCurrentTimeColumnCountExceedsLimit.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("Table `%s` has %d ON UPDATE CURRENT_TIMESTAMP() columns. The count greater than %d.", table.tableName, table.onUpdateCurrentTimeCount, maxOnUpdateCurrentTimeColumnCount),
				StartPosition: &storepb.Position{
					Line: int32(table.line),
				},
			})
		}
	}
	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  storepb.Advice_SUCCESS,
			Code:    advisor.Ok.Int32(),
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList
}

func (checker *columnCurrentTimeCountLimitChecker) count(tableName string, column *ast.ColumnDef, line int) {
	switch column.Tp.GetType() {
	case mysql.TypeDatetime, mysql.TypeTimestamp:
		if isDefaultCurrentTime(column) {
			table, exists := checker.tableSet[tableName]
			if !exists {
				table = tableData{
					tableName: tableName,
				}
			}
			table.defaultCurrentTimeCount++
			table.line = line
			checker.tableSet[tableName] = table
		}
		if isOnUpdateCurrentTime(column) {
			table, exists := checker.tableSet[tableName]
			if !exists {
				table = tableData{
					tableName: tableName,
				}
			}
			table.onUpdateCurrentTimeCount++
			table.line = line
			checker.tableSet[tableName] = table
		}
	}
}

// Enter implements the ast.Visitor interface.
func (checker *columnCurrentTimeCountLimitChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		tableName := node.Table.Name.O
		for _, column := range node.Cols {
			checker.count(tableName, column, node.OriginTextPosition())
		}
	case *ast.AlterTableStmt:
		tableName := node.Table.Name.O
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					checker.count(tableName, column, node.OriginTextPosition())
				}
			case ast.AlterTableModifyColumn, ast.AlterTableChangeColumn:
				checker.count(tableName, spec.NewColumns[0], node.OriginTextPosition())
			}
		}
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnCurrentTimeCountLimitChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func isOnUpdateCurrentTime(column *ast.ColumnDef) bool {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionOnUpdate {
			if function, ok := option.Expr.(*ast.FuncCallExpr); ok && isCurrentTime(function.FnName.L) {
				return true
			}
		}
	}
	return false
}

func isDefaultCurrentTime(column *ast.ColumnDef) bool {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionDefaultValue {
			if function, ok := option.Expr.(*ast.FuncCallExpr); ok && isCurrentTime(function.FnName.L) {
				return true
			}
		}
	}
	return false
}

func isCurrentTime(name string) bool {
	switch strings.ToLower(name) {
	// Any of the synonyms for CURRENT_TIMESTAMP have the same meaning as CURRENT_TIMESTAMP.
	// These are CURRENT_TIMESTAMP(), NOW(), LOCALTIME, LOCALTIME(), LOCALTIMESTAMP, and LOCALTIMESTAMP().
	// See https://dev.mysql.com/doc/refman/8.0/en/timestamp-initialization.html.
	case "current_timestamp", "now", "localtime", "localtimestamp":
		return true
	}
	return false
}
