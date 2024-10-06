package client

import (
	"bytes"
	"fmt"
	"io"
	"kidstales/internal/cache"
	"kidstales/internal/config"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	cache      *cache.Cache
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: config.ClientTimeout,
		},
		cache: cache.NewHttpRequestCache(config.CacheDir),
	}
}

func (c *Client) GetWithCache(path string) (io.ReadCloser, error) {
	cacheKey := &http.Request{
		RequestURI: path,
		URL:        &url.URL{},
	}

	data, found := c.cache.Get(cacheKey)
	if found {
		log.Printf("return cached %s", path)
		return io.NopCloser(bytes.NewBuffer(data)), nil
	}

	log.Printf("do request %s", path)

	bodyReader, err := c.Get(path)
	if err != nil {
		return nil, err
	}

	defer bodyReader.Close()

	data, err = io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	c.cache.Put(cacheKey, data)

	return io.NopCloser(bytes.NewBuffer(data)), nil
}

func (c *Client) Get(path string) (io.ReadCloser, error) {
	uri := config.Host + path

	resp, err := c.httpClient.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to get page %s: %w", uri, err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("get %s: unexpected status code %d", uri, resp.StatusCode)
	}

	return resp.Body, nil
}
