package task

type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	LoginName         string `json:"login_name"`
	SourceID          int    `json:"source_id"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	AvatarURL         string `json:"avatar_url"`
	HtmlURL           string `json:"html_url"`
	Language          string `json:"language"`
	IsAdmin           bool   `json:"is_admin"`
	LastLogin         string `json:"last_login"`
	CreatedAt         string `json:"created_at"`
	Restricted        bool   `json:"restricted"`
	Active            bool   `json:"active"`
	ProhibitLogin     bool   `json:"prohibit_login"`
	Location          string `json:"location"`
	Pronouns          string `json:"pronouns"`
	Website           string `json:"website"`
	Description       string `json:"description"`
	Visibility        string `json:"visibility"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	StarredReposCount int    `json:"starred_repos_count"`
	Username          string `json:"username"`
}
