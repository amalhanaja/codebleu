package bitbucket

type PostPullRequestCommentRequest struct {
	Content *PullRequestCommentContent `json:"content"`
}

type PullRequestCommentContent struct {
	Raw string `json:"raw"`
}

type PullRequestResponse struct {
	CommentCount        int64                   `json:"comment_count"`
	TaskCount           int64                   `json:"task_count"`
	Type                string                  `json:"type"`
	ID                  int64                   `json:"id"`
	Title               string                  `json:"title"`
	Description         string                  `json:"description"`
	Rendered            Rendered                `json:"rendered"`
	State               string                  `json:"state"`
	MergeCommitResponse *MergeCommitResponse    `json:"merge_commit"`
	CloseSourceBranch   bool                    `json:"close_source_branch"`
	ClosedBy            interface{}             `json:"closed_by"`
	Author              *AuthorResponse         `json:"author"`
	Reason              string                  `json:"reason"`
	CreatedOn           string                  `json:"created_on"`
	UpdatedOn           string                  `json:"updated_on"`
	Destination         *AddressResponse        `json:"destination"`
	Source              *AddressResponse        `json:"source"`
	Reviewers           []interface{}           `json:"reviewers"`
	Participants        []interface{}           `json:"participants"`
	Links               map[string]LinkResponse `json:"links"`
	Summary             *SummaryResponse        `json:"summary"`
}

type MergeCommitResponse struct {
	Hash string `json:"hash"`
}

type AuthorResponse struct {
	DisplayName string              `json:"display_name"`
	Links       AuthorLinksResponse `json:"links"`
	Type        string              `json:"type"`
	UUID        string              `json:"uuid"`
	AccountID   string              `json:"account_id"`
	Nickname    string              `json:"nickname"`
}

type AuthorLinksResponse struct {
	Self   LinkResponse `json:"self"`
	Avatar LinkResponse `json:"avatar"`
	HTML   LinkResponse `json:"html"`
}

type LinkResponse struct {
	Href string `json:"href"`
}

type AddressResponse struct {
	Branch     BranchResponse     `json:"branch"`
	Commit     CommitResponse     `json:"commit"`
	Repository RepositoryResponse `json:"repository"`
}

type BranchResponse struct {
	Name string `json:"name"`
}

type CommitResponse struct {
	Hash  string              `json:"hash"`
	Links CommitLinksResponse `json:"links"`
	Type  string              `json:"type"`
}

type CommitLinksResponse struct {
	Self LinkResponse `json:"self"`
	HTML LinkResponse `json:"html"`
}

type RepositoryResponse struct {
	Type     string              `json:"type"`
	FullName string              `json:"full_name"`
	Links    AuthorLinksResponse `json:"links"`
	Name     string              `json:"name"`
	UUID     string              `json:"uuid"`
}

type Rendered struct {
	Title       *SummaryResponse `json:"title"`
	Description *SummaryResponse `json:"description"`
}

type SummaryResponse struct {
	Type   string `json:"type"`
	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
}
