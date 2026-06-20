package models

type Repository struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description"`
	Language        string   `json:"language"`
	HTMLURL         string   `json:"html_url"`
	CloneURL        string   `json:"clone_url"`
	StarsCount      int      `json:"stars_count"`
	ForksCount      int      `json:"forks_count"`
	WatchersCount   int      `json:"watchers_count"`
	OpenIssuesCount int      `json:"open_issues_count"`
	DefaultBranch   string   `json:"default_branch"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	Private         bool     `json:"private"`
	Archived        bool     `json:"archived"`
	Topics          []string `json:"topics"`
	Owner           Owner    `json:"owner"`
}

type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Login    string `json:"login"`
	HTMLURL  string `json:"html_url"`
}
