package pushshift

import (
	"net/http"
	"net/url"
	"time"
)

const (
	submissionEndpoint string = "https://api.pushshift.io/reddit/search/submission"
	commentsEndpoint   string = "https://api.pushshift.io/reddit/search/comment"
	sseEndpoint        string = "http://stream.pushshift.io"
	// HTTPTimeout is the timeout in seconds for the HTTP client
	HTTPTimeout int = 5
	// SortAsc is a helper constant to sort ascending
	SortAsc string = "asc"
	// SortDesc is a helper constant to sort ascending
	SortDesc string = "desc"
)

// NewClient creates a new client
func NewClient(userAgent string) *Client {
	httpClient := newHTTPClient(HTTPTimeout)
	return &Client{
		http:      httpClient,
		userAgent: userAgent,
	}
}

// NewClientWithProxy creates a new client that uses a proxy
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

// Client is the base pushshift HTTP client
type Client struct {
	http      *http.Client
	userAgent string
}
