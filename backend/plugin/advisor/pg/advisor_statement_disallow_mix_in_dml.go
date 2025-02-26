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
	_ advisor.Advisor = (*StatementDisallowMixInDMLAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementDisallowMixInDML, &StatementDisallowMixInDMLAdvisor{})
}

// StatementDisallowMixInDMLAdvisor is the advisor checking for disallow mix DDL and DML.
type StatementDisallowMixInDMLAdvisor struct {
}

// Check checks for disallow mix DDL and DML.
func (*StatementDisallowMixInDMLAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	switch checkCtx.ChangeType {
	case storepb.PlanCheckRunConfig_DML:
	default:
		return nil, nil
	}
	stmtList, ok := checkCtx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	title := string(checkCtx.Rule.Type)

	var adviceList []*storepb.Advice
	for _, stmt := range stmtList {
		if _, ok := stmt.(ast.DDLNode); ok {
			adviceList = append(adviceList, &storepb.Advice{
				Status:  level,
				Title:   title,
				Content: fmt.Sprintf("Data change can only run DML, \"%s\" is not DML", stmt.Text()),
				Code:    advisor.StatementDisallowMixDDLDML.Int32(),
				StartPosition: &storepb.Position{
					Line: int32(stmt.LastLine()),
				},
			})
		}
	}

	return adviceList, nil
}
