package gitrepo

import (
	"codebleu/internal/domain/gitrepo"
	"codebleu/internal/infrastructure/http/bitbucket"
	"codebleu/pkg/env"
	"fmt"
)

func NewRepository(provider string) (gitrepo.Repository, error) {
	switch provider {
	case "bitbucket":
		return bitbucket.NewClient(
			env.MustString("BITBUCKET_WORKSPACE"),
			env.MustString("BITBUCKET_REPO_SLUG"),
			env.MustString("BITBUCKET_ACCESS_TOKEN"),
		), nil
	default:
		return nil, fmt.Errorf("unsupported remote repository provider %s", provider)
	}
}