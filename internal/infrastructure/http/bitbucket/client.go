package bitbucket

import (
	domain "codebleu/internal/domain/gitrepo"
	"net/http"
)

type client struct {
	httpClient  *http.Client
	workspace   string
	repoSlug    string
	accessToken string
}

func NewClient(workspace string, repoSlug string, accessToken string) domain.Repository {
	httpClient := &http.Client{}
	return &client{
		httpClient:  httpClient,
		workspace:   workspace,
		repoSlug:    repoSlug,
		accessToken: accessToken,
	}
}
