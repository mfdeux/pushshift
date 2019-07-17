package pushshift

import (
	"net/http"
	"net/url"
	"time"
)

const (
	submissionEndpoint string = "https://api.pushshift.io/reddit/submission/search"
	commentsEndpoint   string = "https://api.pushshift.io/reddit/comments/search"
	sseEndpoint        string = "http://stream.pushshift.io"
	HTTPTimeout        int    = 5
	// SortAsc is a helper constant to sort ascending
	SortAsc string = "asc"
	// SortDesc is a helper constant to sort ascending
	SortDesc string = "desc"
)

func NewClient(userAgent string) *Client {
	httpClient := newHTTPClient(HTTPTimeout)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}
}

func NewClientWithProxy(userAgent string, proxyURL string) (*Client, error) {
	httpClient, err := newHTTPClientWithProxy(HTTPTimeout, proxyURL)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}, err
}

func newHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func newHTTPClientWithProxy(timeout int, proxyURL string) (*http.Client, error) {
	parsedProxyURL, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(parsedProxyURL),
	}
	netClient := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: transport,
	}
	return netClient, nil
}

type Client struct {
	http      *http.Client
	userAgent string
}

func (c *Client) makeURL(endpoint, query string) string {
	return ""
}
