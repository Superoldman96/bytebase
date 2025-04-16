package tidb

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*IndexKeyNumberLimitAdvisor)(nil)
	_ ast.Visitor     = (*indexKeyNumberLimitChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLIndexKeyNumberLimit, &IndexKeyNumberLimitAdvisor{})
}

// IndexKeyNumberLimitAdvisor is the advisor checking for index key number limit.
type IndexKeyNumberLimitAdvisor struct {
}

// Check checks for index key number limit.
func (*IndexKeyNumberLimitAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &indexKeyNumberLimitChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
		max:   payload.Number,
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.adviceList, nil
}

type indexKeyNumberLimitChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
	line       int
	max        int
}

type indexData struct {
	table string
	index string
	line  int
}

// Enter implements the ast.Visitor interface.
func (checker *indexKeyNumberLimitChecker) Enter(in ast.Node) (ast.Node, bool) {
	var indexList []indexData

	appendIndexItem := func(table, index string, line int) {
		indexList = append(indexList, indexData{
			table: table,
			index: index,
			line:  line,
		})
	}

	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, constraint := range node.Constraints {
			if checker.max > 0 && indexKeyNumber(constraint) > checker.max {
				appendIndexItem(node.Table.Name.O, constraint.Name, constraint.OriginTextPosition())
			}
		}
	case *ast.CreateIndexStmt:
		if checker.max > 0 && len(node.IndexPartSpecifications) > checker.max {
			appendIndexItem(node.Table.Name.O, node.IndexName, checker.line)
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			if spec.Tp == ast.AlterTableAddConstraint {
				if checker.max > 0 && indexKeyNumber(spec.Constraint) > checker.max {
					appendIndexItem(node.Table.Name.O, spec.Constraint.Name, checker.line)
				}
			}
		}
	}

	for _, index := range indexList {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.IndexKeyNumberExceedsLimit.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("The number of index `%s` in table `%s` should be not greater than %d", index.index, index.table, checker.max),
			StartPosition: common.ConvertANTLRLineToPosition(index.line),
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*indexKeyNumberLimitChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func indexKeyNumber(constraint *ast.Constraint) int {
	switch constraint.Tp {
	case ast.ConstraintIndex,
		ast.ConstraintPrimaryKey,
		ast.ConstraintUniq,
		ast.ConstraintUniqKey,
		ast.ConstraintUniqIndex,
		ast.ConstraintForeignKey:
		return len(constraint.Keys)
	default:
		return 0
	}
}
