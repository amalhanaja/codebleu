package codeassistant

type PullRequestReviewInput struct {
	PullRequestId string
	DiffPatch     string
	Description   string
	Title         string
}
