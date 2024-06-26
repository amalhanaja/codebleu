package repowebhooks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/webhooks/v6/bitbucket"
)

func ParsePullRequestEvent(webhook *bitbucket.Webhook, r *http.Request) (string, error) {
	result, err := webhook.Parse(r, bitbucket.PullRequestCreatedEvent, bitbucket.PullRequestUpdatedEvent)
	if err != nil {
		return "", err
	}
	switch payload := result.(type) {
	case bitbucket.PullRequestCreatedPayload:
		return strconv.Itoa(int(payload.PullRequest.ID)), nil
	case bitbucket.PullRequestUpdatedPayload:
		return strconv.Itoa(int(payload.PullRequest.ID)), nil
	default:
		log.Println("event is not registered")
		return "", nil
	}
}
