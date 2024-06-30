package github

import (
	"bytes"
	infraHttp "codebleu/internal/infrastructure/http"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *client) getBaseUrl() string {
	return "https://api.github.com"
}

func (c *client) buildRequest(ctx context.Context, method string, requestUrl string, bodyPayload io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, requestUrl, bodyPayload)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("failed construct request", requestUrl), err)
	}
	if bodyPayload != nil {
		request.Header.Add("Accept", "application/vnd.github+json")
		request.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return request, nil
}

func (c *client) buildRequestPayload(payload interface{}) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonPayload), nil
}
