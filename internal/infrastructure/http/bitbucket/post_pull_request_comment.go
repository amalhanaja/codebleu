package bitbucket

import (
	domain "codebleu/internal/domain/gitrepo"
	"context"
	"fmt"
	"net/http"
)

// PostPullRequestComment implements bitbucket.Repository.
func (c *client) PostPullRequestComment(ctx context.Context, input domain.PostPullRequestCommentInput) error {
	var response map[string]interface{}
	payload := &PostPullRequestCommentRequest{
		Content: &PullRequestCommentContent{
			Raw: input.Comment,
		},
		Inline: &PullRequestCommentInline{
			Path: input.Path,
		},
	}
	if err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/pullrequests/%s/comments", input.PullRequestId), payload, &response); err != nil {
		return err
	}
	return nil
}
