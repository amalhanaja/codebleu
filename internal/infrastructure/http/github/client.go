package github

import (
	"codebleu/internal/domain/gitrepo"
	"context"
	"fmt"
	"net/http"
)

type client struct {
	accessToken string
	owner       string
	repoSlug    string
	httpClient  *http.Client
}

// PostPullRequestComment implements gitrepo.Repository.
func (c *client) PostPullRequestComment(ctx context.Context, input gitrepo.PostPullRequestCommentInput) error {
	fmt.Println(input.Comment)
	return nil
}

func NewClient(
	owner string,
	repoSlug string,
	accessToken string,
) gitrepo.Repository {
	httpClient := &http.Client{}
	return &client{
		owner:       owner,
		repoSlug:    repoSlug,
		accessToken: accessToken,
		httpClient:  httpClient,
	}
}
