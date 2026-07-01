package models

type PullRequest struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`

	Draft        bool `json:"draft"`
	Mergeable    bool `json:"mergeable"`
	Merged       bool `json:"merged"`
	Comments     int  `json:"comments"`
	Additions    int  `json:"additions"`
	Deletions    int  `json:"deletions"`
	ChangedFiles int  `json:"changed_files"`

	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ClosedAt  string `json:"closed_at"`
	MergedAt  string `json:"merged_at"`

	User User `json:"user"`

	Base PullRequestBranch `json:"base"`
	Head PullRequestBranch `json:"head"`
}

type PullRequestBranch struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	SHA   string `json:"sha"`

	Repo Repository `json:"repo"`
}
