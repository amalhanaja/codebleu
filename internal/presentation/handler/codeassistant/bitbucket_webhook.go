package codeassistant

import (
	"codebleu/internal/infrastructure/thirdparty/repowebhooks"
	"context"
	"fmt"
	"net/http"
)

// BitbucketWebhook implements Handler.
func (h *handlerImpl) BitbucketWebhook(w http.ResponseWriter, r *http.Request) {
	pullRequestId, err := repowebhooks.ParsePullRequestEvent(h.bitbucketWebhook, r)
	if err != nil {
		fmt.Println("Failed get Pull Request ID")
		w.Write([]byte("Failed get Pull Request ID"))
		return
	}
	if _, err := h.reviewAndCommentPullRequest.Invoke(context.Background(), pullRequestId); err != nil {
		fmt.Println("Review Fail")
		w.Write([]byte("Review Fail"))
		return
	}
	fmt.Println("Success : " + pullRequestId)
	w.WriteHeader(200)
}
