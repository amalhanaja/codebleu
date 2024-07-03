package codeassistant

import (
	domain "codebleu/internal/domain/codeassistant"
	gitRepoDomain "codebleu/internal/domain/gitrepo"
	"codebleu/internal/usecase"
	"context"
	"errors"
)

type reviewAndCommentPullRequest struct {
	getPullRequest         usecase.UseCase[string, *gitRepoDomain.PullRequest]
	reviewPullRequest      usecase.UseCase[domain.PullRequestReviewInput, []domain.ReviewResult]
	postPullRequestComment usecase.UseCase[gitRepoDomain.PostPullRequestCommentInput, interface{}]
}

// Invoke implements usecase.UseCase.
func (r *reviewAndCommentPullRequest) Invoke(ctx context.Context, input domain.ReviewAndCommentPullRequestInput) (interface{}, error) {
	pullRequest, err := r.getPullRequest.Invoke(ctx, input.PullRequestId)
	if err != nil {
		return nil, err
	}
	reviewResults, err := r.reviewPullRequest.Invoke(ctx, domain.PullRequestReviewInput{
		PullRequest:       pullRequest,
		SystemInstruction: input.SystemInstruction,
	})
	if err != nil {
		return nil, err
	}
	for _, reviewResult := range reviewResults {
		_, postCommentErr := r.postPullRequestComment.Invoke(ctx, gitRepoDomain.PostPullRequestCommentInput{
			PullRequestId: input.PullRequestId,
			CommitHash:    pullRequest.CommitHash,
			Path:          reviewResult.Path,
			Comment:       reviewResult.Comment,
		})
		if postCommentErr != nil {
			if err == nil {
				err = errors.New("failed post comment")
			}
			err = errors.Join(err, postCommentErr)
		}
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func ReviewAndCommentPullRequest(
	getPullRequest usecase.UseCase[string, *gitRepoDomain.PullRequest],
	reviewPullRequest usecase.UseCase[domain.PullRequestReviewInput, []domain.ReviewResult],
	postPullRequestComment usecase.UseCase[gitRepoDomain.PostPullRequestCommentInput, interface{}],
) usecase.UseCase[domain.ReviewAndCommentPullRequestInput, interface{}] {
	return &reviewAndCommentPullRequest{
		reviewPullRequest:      reviewPullRequest,
		postPullRequestComment: postPullRequestComment,
		getPullRequest:         getPullRequest,
	}
}
