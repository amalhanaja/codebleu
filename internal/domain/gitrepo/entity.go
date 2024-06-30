package gitrepo

type PostPullRequestCommentInput struct {
	PullRequestId string
	CommitHash    string
	Comment       string
}

type PullRequest struct {
	Id          string
	Title       string
	Description string
	DiffPatch   string
	CommitHash  string
}
