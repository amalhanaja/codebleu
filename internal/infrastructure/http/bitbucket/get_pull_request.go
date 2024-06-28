package bitbucket

import (
	domain "codebleu/internal/domain/gitrepo"
	infraHttp "codebleu/internal/infrastructure/http"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// GetPullRequest implements bitbucket.Repository.
func (c *client) GetPullRequest(ctx context.Context, id string) (*domain.PullRequest, error) {
	pullRequestResponse, err := c.getPullRequest(ctx, id)
	if err != nil {
		return nil, err
	}
	diff, err := c.getPullRequestDiff(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.PullRequest{
		Id:          strconv.Itoa(int(pullRequestResponse.ID)),
		Title:       pullRequestResponse.Title,
		Description: pullRequestResponse.Description,
		DiffPatch:   diff,
	}, nil
}

// GetPullRequest implements Client.
func (c *client) getPullRequest(ctx context.Context, id string) (*PullRequestResponse, error) {
	var response *PullRequestResponse
	err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/pullrequests/%s", id), nil, &response)
	return response, err
}

func (c *client) getPullRequestDiff(ctx context.Context, id string) (string, error) {
	path := fmt.Sprintf("/pullrequests/%s/diff", id)
	response, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "nil", errors.Join(infraHttp.NewHttpClientError("failed read response", path), err)
	}
	return string(body), nil
}
