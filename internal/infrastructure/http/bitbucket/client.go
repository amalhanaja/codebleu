package bitbucket

import (
	domain "codebleu/internal/domain/gitrepo"
	"context"
	"fmt"
	"net/http"
)

type client struct {
	httpClient  http.Client
	workspace   string
	repoSlug    string
	accessToken string
}

// GetPullRequest implements Client.
func (c *client) getPullRequest(ctx context.Context, id string) (*PullRequestResponse, error) {
	var response *PullRequestResponse
	err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/pullrequests/%s", id), nil, &response)
	return response, err
}

func NewClient(workspace string, repoSlug string, accessToken string) domain.Repository {
	httpClient := http.Client{}
	return &client{
		httpClient:  httpClient,
		workspace:   workspace,
		repoSlug:    repoSlug,
		accessToken: accessToken,
	}
}
