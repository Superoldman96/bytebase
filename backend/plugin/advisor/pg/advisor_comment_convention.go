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
	_ advisor.Advisor = (*CommentConventionAdvisor)(nil)
	_ ast.Visitor     = (*commentConventionChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_POSTGRES, advisor.PostgreSQLCommentConvention, &CommentConventionAdvisor{})
}

// CommentConventionAdvisor is the advisor checking for comment convention.
type CommentConventionAdvisor struct {
}

// Check checks for comment convention.
func (*CommentConventionAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
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
	checker := &commentConventionChecker{
		level:     level,
		title:     string(checkCtx.Rule.Type),
		maxLength: payload.Number,
	}

	for _, stmt := range stmtList {
		ast.Walk(checker, stmt)
	}

	return checker.adviceList, nil
}

type commentConventionChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	maxLength  int
}

// Visit implements the ast.Visitor interface.
func (checker *commentConventionChecker) Visit(node ast.Node) ast.Visitor {
	type commentData struct {
		comment string
		line    int
	}
	var comment commentData

	if n, ok := node.(*ast.CommentStmt); ok {
		comment = commentData{
			comment: n.Comment,
			line:    n.LastLine(),
		}
	}

	if checker.maxLength > 0 && len(comment.comment) > checker.maxLength {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.CommentTooLong.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("The length of comment should be within %d characters", checker.maxLength),
			StartPosition: advisor.ConvertANTLRLineToPosition(comment.line),
		})
	}

	return checker
}
