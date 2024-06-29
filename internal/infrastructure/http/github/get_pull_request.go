package github

import (
	"codebleu/internal/domain/gitrepo"
	infraHttp "codebleu/internal/infrastructure/http"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GetPullRequest implements gitrepo.Repository.
func (c *client) GetPullRequest(ctx context.Context, id string) (*gitrepo.PullRequest, error) {
	pullRequest, err := c.getPullRequest(ctx, id)
	if err != nil {
		return nil, err
	}
	diff, err := c.getPullRequestDiff(ctx, id)
	if err != nil {
		return nil, err
	}
	return &gitrepo.PullRequest{
		Id:          id,
		Title:       pullRequest.Title,
		Description: pullRequest.Body,
		DiffPatch:   diff,
	}, nil
}

func (c *client) getPullRequestDiff(ctx context.Context, id string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%s", c.getBaseUrl(), c.owner, c.repoSlug, id)
	req, err := c.buildRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.github.v3.diff")
	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return "", infraHttp.NewHttpClientError(fmt.Sprintf("response failed %s", string(body)), url)
	}
	return string(body), nil
}

func (c *client) getPullRequest(ctx context.Context, id string) (*PullRequestResponse, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%s", c.getBaseUrl(), c.owner, c.repoSlug, id)
	req, err := c.buildRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	var responseBody *PullRequestResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&responseBody)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("failed decode repsonse", url))
	}
	return responseBody, nil
}
