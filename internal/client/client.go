package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	mainPageURL = "https://skazkiwsem.fun/"
	timeout     = 30 * time.Second
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) MainPage() (io.ReadCloser, error) {
	resp, err := c.httpClient.Get(mainPageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get main page: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
