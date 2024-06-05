package pg

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*CollationAllowlistAdvisor)(nil)
	_ ast.Visitor     = (*collationAllowlistChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLCollationAllowlist, &CollationAllowlistAdvisor{})
}

// CollationAllowlistAdvisor is the advisor checking for collation allowlist.
type CollationAllowlistAdvisor struct {
}

// Check checks for collation allowlist.
func (*CollationAllowlistAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalStringArrayTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &collationAllowlistChecker{
		level:     level,
		title:     string(ctx.Rule.Type),
		allowlist: make(map[string]bool),
	}
	for _, collation := range payload.List {
		checker.allowlist[collation] = true
	}

	for _, stmt := range stmtList {
		checker.text = advisor.NormalizeStatement(stmt.Text())
		ast.Walk(checker, stmt)
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

type collationAllowlistChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	allowlist  map[string]bool
}

// Visit implements ast.Visitor interface.
func (checker *collationAllowlistChecker) Visit(in ast.Node) ast.Visitor {
	code := advisor.Ok
	line := 0
	disabledCollation := ""
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.ColumnList {
			if column.Collation != nil {
				if _, exists := checker.allowlist[column.Collation.Name]; !exists {
					code = advisor.DisabledCollation
					line = column.LastLine()
					disabledCollation = column.Collation.Name
					break
				}
			}
		}
	case *ast.AlterTableStmt:
		for _, item := range node.AlterItemList {
			switch itemNode := item.(type) {
			case *ast.AddColumnListStmt:
				for _, column := range itemNode.ColumnList {
					if column.Collation != nil {
						if _, exists := checker.allowlist[column.Collation.Name]; !exists {
							code = advisor.DisabledCollation
							line = node.LastLine()
							disabledCollation = column.Collation.Name
							break
						}
					}
				}
			case *ast.AlterColumnTypeStmt:
				if itemNode.Collation != nil {
					if _, exists := checker.allowlist[itemNode.Collation.Name]; !exists {
						code = advisor.DisabledCollation
						line = node.LastLine()
						disabledCollation = itemNode.Collation.Name
						break
					}
				}
			}

			if code != advisor.Ok {
				break
			}
		}
	}

	if code != advisor.Ok {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  checker.level,
			Code:    code.Int32(),
			Title:   checker.title,
			Content: fmt.Sprintf("Use disabled collation \"%s\", related statement \"%s\"", disabledCollation, checker.text),
			StartPosition: &storepb.Position{
				Line: int32(line),
			},
		})
	}

	return checker
}
