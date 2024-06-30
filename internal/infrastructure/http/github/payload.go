package github

type PullRequestResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Head  struct {
		Sha string `json:"sha"`
	} `json:"Head"`
}

type PostPullRequestCommentRequest struct {
	Body     string `json:"body"`
	CommitId string `json:"commit_id"`
	Path     string `json:"path"`
	Line     int    `json:"line"`
}
