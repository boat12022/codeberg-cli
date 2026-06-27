package models

type Issue struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`

	HTMLURL   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ClosedAt  string `json:"closed_at"`

	Comments int `json:"comments"`

	User       User            `json:"user"`
	Repository IssueRepository `json:"repository"`
	Labels     []Label         `json:"labels"`
}

type IssueRepository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	FullName string `json:"full_name"`
}

type Label struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}
