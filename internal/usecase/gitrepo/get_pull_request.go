package gitrepo

import (
	domain "codebleu/internal/domain/gitrepo"
	u "codebleu/internal/usecase"
	"context"
)

type getPullRequest struct {
	repository domain.Repository
}

// Invoke implements usecase.UseCase.
func (u *getPullRequest) Invoke(ctx context.Context, input string) (*domain.PullRequest, error) {
	return u.repository.GetPullRequest(ctx, input)
}

func GetPullRequest(
	repository domain.Repository,
) u.UseCase[string, *domain.PullRequest] {
	return &getPullRequest{
		repository: repository,
	}
}
