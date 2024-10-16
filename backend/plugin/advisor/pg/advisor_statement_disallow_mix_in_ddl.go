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
	_ advisor.Advisor = (*StatementDisallowMixInDDLAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementDisallowMixInDDL, &StatementDisallowMixInDDLAdvisor{})
}

// StatementDisallowMixInDDLAdvisor is the advisor checking for disallow mix DDL and DML.
type StatementDisallowMixInDDLAdvisor struct {
}

// Check checks for disallow mix DDL and DML.
func (*StatementDisallowMixInDDLAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	switch ctx.ChangeType {
	case storepb.PlanCheckRunConfig_DDL, storepb.PlanCheckRunConfig_SDL, storepb.PlanCheckRunConfig_DDL_GHOST:
	default:
		return nil, nil
	}
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, errors.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	title := string(ctx.Rule.Type)

	var adviceList []*storepb.Advice
	for _, stmt := range stmtList {
		if _, ok := stmt.(ast.DMLNode); ok {
			adviceList = append(adviceList, &storepb.Advice{
				Status:  level,
				Title:   title,
				Content: fmt.Sprintf("Alter schema can only run DDL, \"%s\" is not DDL", stmt.Text()),
				Code:    advisor.StatementDisallowMixDDLDML.Int32(),
				StartPosition: &storepb.Position{
					Line: int32(stmt.LastLine()),
				},
			})
		}
	}

	return adviceList, nil
}
