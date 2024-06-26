package codeassistant

import (
	"codebleu/internal/usecase"
	"net/http"

	"github.com/go-playground/webhooks/v6/bitbucket"
)

type Handler interface {
	BitbucketWebhook(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	reviewAndCommentPullRequest usecase.UseCase[string, interface{}]
	bitbucketWebhook            *bitbucket.Webhook
}

func NewHandler(reviewAndCommentPullRequest usecase.UseCase[string, interface{}], bitbucketWebhook *bitbucket.Webhook) Handler {
	return &handlerImpl{
		reviewAndCommentPullRequest: reviewAndCommentPullRequest,
		bitbucketWebhook:            bitbucketWebhook,
	}
}
