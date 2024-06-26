package gitrepo

import (
	domain "codebleu/internal/domain/gitrepo"
	u "codebleu/internal/usecase"
	"context"
)

type postPullRequestCommentUseCase struct {
	repository domain.Repository
}

// Invoke implements usecase.UseCase.
func (p *postPullRequestCommentUseCase) Invoke(ctx context.Context, input domain.PostPullRequestCommentInput) (interface{}, error) {
	err := p.repository.PostPullRequestComment(ctx, input)
	return nil, err
}

func PostPullRequestComment(
	repository domain.Repository,
) u.UseCase[domain.PostPullRequestCommentInput, interface{}] {
	return &postPullRequestCommentUseCase{
		repository: repository,
	}
}
