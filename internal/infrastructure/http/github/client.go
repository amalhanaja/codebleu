package github

import (
	"codebleu/internal/domain/gitrepo"
	"net/http"
)

type client struct {
	accessToken string
	owner       string
	repoSlug    string
	httpClient  *http.Client
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
