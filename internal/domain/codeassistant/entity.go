package codeassistant

import (
	"codebleu/internal/domain/gitrepo"
)

type PullRequestReviewInput struct {
	PullRequest       *gitrepo.PullRequest
	SystemInstruction string
}

type ReviewAndCommentPullRequestInput struct {
	PullRequestId     string
	SystemInstruction string
}

type ReviewResult struct {
	Path    string `json:"path"`
	Comment string `json:"comment_in_markdown"`
}
