package codeassistant

import "codebleu/internal/domain/gitrepo"

type PullRequestReviewInput struct {
	PullRequest  *gitrepo.PullRequest
	ReviewPrompt string
}
