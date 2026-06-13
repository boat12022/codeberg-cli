package codeberg

import (
	"encoding/json"
	"fmt"
	"net/http"

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
