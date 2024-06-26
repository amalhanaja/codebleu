package gitrepo

import "context"

type Repository interface {
	GetPullRequest(ctx context.Context, id string) (*PullRequest, error)
	PostPullRequestComment(ctx context.Context, input PostPullRequestCommentInput) error
}
