package gitrepo

type PostPullRequestCommentInput struct {
	PullRequestId string
	Comment       string
}

type PullRequest struct {
	Id          string
	Title       string
	Description string
	DiffPatch   string
}
