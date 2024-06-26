package bitbucket

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
	return fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s", c.workspace, c.repoSlug)
}

func (c *client) do(ctx context.Context, method string, path string, payload interface{}) (*http.Response, error) {
	requestUrl := c.getBaseUrl() + path
	bodyPayload, err := c.buildRequestPayload(payload)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("failed marshal payload", requestUrl), err)
	}
	request, err := c.buildRequest(ctx, method, requestUrl, bodyPayload)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(request)
}

func (c *client) doRequest(ctx context.Context, method string, path string, payload interface{}, out interface{}) error {
	requestUrl := c.getBaseUrl() + path
	response, err := c.do(ctx, method, path, payload)
	if err != nil {
		return errors.Join(infraHttp.NewHttpClientError("request failed", requestUrl), err)
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 {
		body, err := io.ReadAll(response.Body)
		httpError := infraHttp.NewHttpClientError(fmt.Sprintf("request failed: httpstatus = %d\n%s", response.StatusCode, string(body)), requestUrl)
		return errors.Join(httpError, err)
	}
	err = json.NewDecoder(response.Body).Decode(out)
	if err != nil {
		return errors.Join(infraHttp.NewHttpClientError("failed decode repsonse", requestUrl))
	}
	return nil
}

func (c *client) buildRequest(ctx context.Context, method string, requestUrl string, bodyPayload io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, requestUrl, bodyPayload)
	if err != nil {
		return nil, errors.Join(infraHttp.NewHttpClientError("failed construct request", requestUrl), err)
	}
	if bodyPayload != nil {
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Accept", "application/json")
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
