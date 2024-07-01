package github

import (
	"codebleu/internal/domain/gitrepo"
	infraHttp "codebleu/internal/infrastructure/http"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// PostPullRequestComment implements gitrepo.Repository.
func (c *client) PostPullRequestComment(ctx context.Context, input gitrepo.PostPullRequestCommentInput) error {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%s/comments", c.getBaseUrl(), c.owner, c.repoSlug, input.PullRequestId)
	println(input.CommitHash, input.Comment)
	payload := &PostPullRequestCommentRequest{
		Body:        input.Comment,
		CommitId:    input.CommitHash,
		Path:        input.Path,
		SubjectType: "file",
	}
	jsonPayload, err := c.buildRequestPayload(payload)
	if err != nil {
		return err
	}
	req, err := c.buildRequest(ctx, http.MethodPost, url, jsonPayload)
	if err != nil {
		return err
	}
	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode >= 300 {
		body, err := io.ReadAll(httpResponse.Body)
		httpError := infraHttp.NewHttpClientError(fmt.Sprintf("request failed: httpstatus = %d\n%s", httpResponse.StatusCode, string(body)), url)
		return errors.Join(httpError, err)
	}
	return nil
}
