package tidb

// Framework code is generated by the generator.

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*CollationAllowlistAdvisor)(nil)
	_ ast.Visitor     = (*collationAllowlistChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLCollationAllowlist, &CollationAllowlistAdvisor{})
}

// CollationAllowlistAdvisor is the advisor checking for collation allowlist.
type CollationAllowlistAdvisor struct {
}

// Check checks for collation allowlist.
func (*CollationAllowlistAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
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
		checker.allowlist[strings.ToLower(collation)] = true
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
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
	line       int
	allowlist  map[string]bool
}

// Enter implements the ast.Visitor interface.
func (checker *collationAllowlistChecker) Enter(in ast.Node) (ast.Node, bool) {
	code := advisor.Ok
	var disabledCollation string
	line := checker.line
	switch node := in.(type) {
	case *ast.CreateDatabaseStmt:
		collation := getDatabaseCollation(node.Options)
		if _, exists := checker.allowlist[collation]; collation != "" && !exists {
			code = advisor.DisabledCollation
			disabledCollation = collation
		}
	case *ast.CreateTableStmt:
		collation := getTableCollation(node.Options)
		if _, exist := checker.allowlist[collation]; collation != "" && !exist {
			code = advisor.DisabledCollation
			disabledCollation = collation
			break
		}
		for _, column := range node.Cols {
			collation := getColumnCollation(column)
			if _, exist := checker.allowlist[collation]; collation != "" && !exist {
				code = advisor.DisabledCollation
				disabledCollation = collation
				line = column.OriginTextPosition()
				break
			}
		}
	case *ast.AlterDatabaseStmt:
		collation := getDatabaseCollation(node.Options)
		if _, exist := checker.allowlist[collation]; collation != "" && !exist {
			code = advisor.DisabledCollation
			disabledCollation = collation
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableOption:
				collation := getTableCollation(spec.Options)
				if _, exist := checker.allowlist[collation]; collation != "" && !exist {
					code = advisor.DisabledCollation
					disabledCollation = collation
				}
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					collation := getColumnCollation(column)
					if _, exist := checker.allowlist[collation]; collation != "" && !exist {
						code = advisor.DisabledCollation
						disabledCollation = collation
						break
					}
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				collation := getColumnCollation(spec.NewColumns[0])
				if _, exist := checker.allowlist[collation]; collation != "" && !exist {
					code = advisor.DisabledCollation
					disabledCollation = collation
				}
			}
		}
	}

	if code != advisor.Ok {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:  checker.level,
			Code:    code.Int32(),
			Title:   checker.title,
			Content: fmt.Sprintf("\"%s\" used disabled collation '%s'", checker.text, disabledCollation),
			StartPosition: &storepb.Position{
				Line: int32(line),
			},
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*collationAllowlistChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func getDatabaseCollation(optionList []*ast.DatabaseOption) string {
	for _, option := range optionList {
		if option.Tp == ast.DatabaseOptionCollate {
			return strings.ToLower(option.Value)
		}
	}
	return ""
}

func getTableCollation(optionList []*ast.TableOption) string {
	for _, option := range optionList {
		if option.Tp == ast.TableOptionCollate {
			return strings.ToLower(option.StrValue)
		}
	}
	return ""
}

func getColumnCollation(column *ast.ColumnDef) string {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionCollate {
			return strings.ToLower(option.StrValue)
		}
	}
	return ""
}
