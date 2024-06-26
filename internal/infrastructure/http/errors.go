package http

import "fmt"

type HttpClientError struct {
	Message string
	url     string
}

// Error implements error.
func (h HttpClientError) Error() string {
	return fmt.Sprintf("http client error: %s - %s", h.url, h.Message)
}

func NewHttpClientError(message string, url string) error {
	return HttpClientError{Message: message, url: url}
}
