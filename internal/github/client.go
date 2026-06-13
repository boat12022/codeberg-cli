package github

import "net/http"

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
