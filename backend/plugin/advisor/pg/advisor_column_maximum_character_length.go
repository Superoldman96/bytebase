package pg

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnMaximumCharacterLengthAdvisor)(nil)
	_ ast.Visitor     = (*columnMaximumCharacterLengthChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLColumnMaximumCharacterLength, &ColumnMaximumCharacterLengthAdvisor{})
}

// ColumnMaximumCharacterLengthAdvisor is the advisor checking for maximum character length.
type ColumnMaximumCharacterLengthAdvisor struct {
}

// Check checks for maximum character length.
func (*ColumnMaximumCharacterLengthAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &columnMaximumCharacterLengthChecker{
		level:   level,
		title:   string(checkCtx.Rule.Type),
		maximum: payload.Number,
	}

	if payload.Number > 0 {
		for _, stmt := range stmtList {
			ast.Walk(checker, stmt)
		}
	}

	return checker.adviceList, nil
}

type columnMaximumCharacterLengthChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	maximum    int
}

// Visit implements ast.Visitor interface.
func (checker *columnMaximumCharacterLengthChecker) Visit(in ast.Node) ast.Visitor {
	var tableName, columnName string
	var line int
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.ColumnList {
			charLength := getCharLength(column)
			if charLength > checker.maximum {
				tableName = normalizeTableName(node.Name, "")
				columnName = column.ColumnName
				line = column.LastLine()
				break
			}
		}
	case *ast.AlterTableStmt:
		for _, item := range node.AlterItemList {
			switch itemNode := item.(type) {
			case *ast.AddColumnListStmt:
				for _, column := range itemNode.ColumnList {
					charLength := getCharLength(column)
					if charLength > checker.maximum {
						tableName = normalizeTableName(node.Table, "")
						columnName = column.ColumnName
						line = itemNode.LastLine()
					}
				}
			case *ast.AlterColumnTypeStmt:
				if char, ok := itemNode.Type.(*ast.Character); ok {
					if char.Size > checker.maximum {
						tableName = normalizeTableName(node.Table, "")
						columnName = itemNode.ColumnName
						line = itemNode.LastLine()
					}
				}
			}
			if tableName != "" {
				break
			}
		}
	}

	if tableName != "" {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  checker.level,
			Code:    advisor.CharLengthExceedsLimit.Int32(),
			Title:   checker.title,
			Content: fmt.Sprintf(`The length of the CHAR column %q in table %s is bigger than %d, please use VARCHAR instead`, columnName, tableName, checker.maximum),
			StartPosition: &storepb.Position{
				Line: int32(line),
			},
		})
	}

	return checker
}

func getCharLength(column *ast.ColumnDef) int {
	if char, ok := column.Type.(*ast.Character); ok {
		return char.Size
	}
	return 0
}
