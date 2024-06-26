package bitbucket

import (
	domain "codebleu/internal/domain/gitrepo"
	infraHttp "codebleu/internal/infrastructure/http"
	"context"
	"errors"
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
	diffLink := pullRequestResponse.Links["diff"].Href
	reqest, err := c.buildRequest(ctx, http.MethodGet, diffLink, nil)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.httpClient.Do(reqest)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("response failed", diffLink), err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != 200 {
		return nil, errors.Join(infraHttp.NewHttpClientError("response not succeed", diffLink), err)
	}
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("failed read response", diffLink), err)
	}
	return &domain.PullRequest{
		Id:          strconv.Itoa(int(pullRequestResponse.ID)),
		Title:       pullRequestResponse.Title,
		Description: pullRequestResponse.Description,
		DiffPatch:   string(body),
	}, nil
}
