package codeberg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"codeberg.org/13thab/codeberg-cli/internal/models"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://codeberg.org/api/v1",
		HTTP:    &http.Client{},
	}
}

func (c *Client) GetUser(username string) (*models.User, error) {
	url := c.BaseURL + "/users/" + username
	req, err := c.HTTP.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) GetRepos(username string) ([]models.Repository, error) {
	url := c.BaseURL + "/users/" + username + "/repos"
	req, err := c.HTTP.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var repos []models.Repository
	if err := json.NewDecoder(req.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (c *Client) GetRepo(username string, repoName string) (*models.Repository, error) {
	url := c.BaseURL + "/repos/" + username + "/" + repoName
	req, err := c.HTTP.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var repo models.Repository
	if err := json.NewDecoder(req.Body).Decode(&repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (c *Client) GetIssues(username string, repoName string, state string, limit int) ([]models.Issue, error) {
	url := c.BaseURL + "/repos/" + username + "/" + repoName + "/issues"
	if state != "" {
		url += "?state=" + state
	}
	if limit > 0 {
		url += "&limit=" + strconv.Itoa(limit)
	}
	req, err := c.HTTP.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var issues []models.Issue
	if err := json.NewDecoder(req.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}
