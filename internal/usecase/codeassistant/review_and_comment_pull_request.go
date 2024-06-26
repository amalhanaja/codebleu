package codeassistant

import (
	domain "codebleu/internal/domain/codeassistant"
	gitRepoDomain "codebleu/internal/domain/gitrepo"
	"codebleu/internal/usecase"
	"context"
)

type reviewAndCommentPullRequest struct {
	getPullRequest         usecase.UseCase[string, *gitRepoDomain.PullRequest]
	reviewPullRequest      usecase.UseCase[domain.PullRequestReviewInput, string]
	postPullRequestComment usecase.UseCase[gitRepoDomain.PostPullRequestCommentInput, interface{}]
}

// Invoke implements usecase.UseCase.
func (r *reviewAndCommentPullRequest) Invoke(ctx context.Context, input string) (interface{}, error) {
	pullRequest, err := r.getPullRequest.Invoke(ctx, input)
	if err != nil {
		return nil, err
	}
	reviewResult, err := r.reviewPullRequest.Invoke(ctx, domain.PullRequestReviewInput{
		PullRequestId: pullRequest.Id,
		DiffPatch:     pullRequest.DiffPatch,
		Title:         pullRequest.Title,
		Description:   pullRequest.Description,
	})
	if err != nil {
		return nil, err
	}
	postPullRequestCommentInput := gitRepoDomain.PostPullRequestCommentInput{
		PullRequestId: input,
		Comment:       reviewResult,
	}
	return r.postPullRequestComment.Invoke(ctx, postPullRequestCommentInput)
}

func ReviewAndCommentPullRequest(
	getPullRequest usecase.UseCase[string, *gitRepoDomain.PullRequest],
	reviewPullRequest usecase.UseCase[domain.PullRequestReviewInput, string],
	postPullRequestComment usecase.UseCase[gitRepoDomain.PostPullRequestCommentInput, interface{}],
) usecase.UseCase[string, interface{}] {
	return &reviewAndCommentPullRequest{
		reviewPullRequest:      reviewPullRequest,
		postPullRequestComment: postPullRequestComment,
		getPullRequest:         getPullRequest,
	}
}
