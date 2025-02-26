package pg

// Framework code is generated by the generator.

import (
	"context"
	"encoding/json"

	pgquery "github.com/pganalyze/pg_query_go/v5"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementDisallowOnDelCascadeAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLStatementDisallowOnDelCascade, &StatementDisallowOnDelCascadeAdvisor{})
}

// StatementDisallowOnDelCascadeAdvisor is the advisor checking the disallow cascade.
type StatementDisallowOnDelCascadeAdvisor struct {
}

// Check checks for DML dry run.
func (*StatementDisallowOnDelCascadeAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmt := checkCtx.Statements
	if stmt == "" {
		return nil, nil
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	jsonText, err := pgquery.ParseToJSON(stmt)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse statement to JSON")
	}

	var jsonData map[string]any
	if err := json.Unmarshal([]byte(jsonText), &jsonData); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal JSON")
	}

	cascadeLocations := cascadeNumRecursive(jsonData, 0, isOnDelCascade)
	cascadePositions := convertLocationsToPositions(stmt, cascadeLocations)

	var adviceList []*storepb.Advice
	for _, p := range cascadePositions {
		adviceList = append(adviceList, &storepb.Advice{
			Status:  level,
			Title:   string(checkCtx.Rule.Type),
			Content: "The CASCADE option is not permitted for ON DELETE clauses",
			Code:    advisor.StatementDisallowCascade.Int32(),
			StartPosition: &storepb.Position{
				Line:   int32(p.line + 1),
				Column: int32(p.column + 1),
			},
		})
	}

	return adviceList, nil
}

func isOnDelCascade(json map[string]any) bool {
	return json["fk_del_action"] == "c"
}
