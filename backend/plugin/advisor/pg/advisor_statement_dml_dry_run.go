package pg

// Framework code is generated by the generator.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementDmlDryRunAdvisor)(nil)
	_ ast.Visitor     = (*statementDmlDryRunChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
}

// StatementDmlDryRunAdvisor is the advisor checking for DML dry run.
type StatementDmlDryRunAdvisor struct {
}

// Check checks for DML dry run.
func (*StatementDmlDryRunAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementDmlDryRunChecker{
		level:  level,
		title:  string(ctx.Rule.Type),
		driver: ctx.Driver,
		ctx:    ctx.Context,
	}

	if checker.driver != nil {
		for _, stmt := range stmtList {
			ast.Walk(checker, stmt)
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
	return checker.adviceList, nil
}

type statementDmlDryRunChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	driver     *sql.DB
	ctx        context.Context
}

// Visit implements ast.Visitor interface.
func (checker *statementDmlDryRunChecker) Visit(in ast.Node) ast.Visitor {
	switch node := in.(type) {
	case *ast.InsertStmt, *ast.UpdateStmt, *ast.DeleteStmt:
		if _, err := advisor.Query(checker.ctx, checker.driver, fmt.Sprintf("EXPLAIN %s", node.Text())); err != nil {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.StatementDMLDryRunFailed.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("\"%s\" dry runs failed: %s", node.Text(), err.Error()),
				StartPosition: &storepb.Position{
					Line: int32(node.LastLine()),
				},
			})
		}
	}

	return checker
}
