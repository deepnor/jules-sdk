package jules

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Jules API.
	DefaultBaseURL = "https://jules.googleapis.com/v1alpha"
)

// Client is a low-level HTTP client for the Jules REST API.
type Client struct {
	client  *http.Client
	BaseURL string
	apiKey  string

	// Services used for talking to different parts of the Jules API.
	Sessions   *SessionsService
	Activities *ActivitiesService
	Sources    *SourcesService
}

type service struct {
	client *Client
}

// SessionsService provides methods for interacting with the Jules sessions API.
type SessionsService service

// ActivitiesService provides methods for interacting with the Jules activities API.
type ActivitiesService service

// SourcesService provides methods for interacting with the Jules sources API.
type SourcesService service

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithHTTPClient sets a custom http.Client for the Client.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.client = hc
	}
}

// WithBaseURL overrides the default Jules API base URL.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithTransport sets a custom Transport for the underlying HTTP client.
func WithTransport(t http.RoundTripper) ClientOption {
	return func(c *Client) {
		c.client.Transport = t
	}
}

// WithTimeout sets a custom timeout for the underlying HTTP client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

// defaultHTTPClient constructs a hardened HTTP client with robust
// connection pooling and timeouts to prevent connection socket exhaustion.
func defaultHTTPClient() *http.Client {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Transport: t,
	}
}

// NewClient creates a new Client with the given API key and options.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		client:  defaultHTTPClient(),
		BaseURL: DefaultBaseURL,
		apiKey:  apiKey,
	}
	for _, opt := range opts {
		opt(c)
	}

	c.Sessions = &SessionsService{client: c}
	c.Activities = &ActivitiesService{client: c}
	c.Sources = &SourcesService{client: c}

	return c
}

// Do executes an HTTP request against the Jules API.
//
// It automatically injects the x-goog-api-key authentication header and sets
// Content-Type to application/json when a request body is provided. Non-2xx
// responses are parsed into an *APIError.
//
// If response is non-nil, the response body is JSON-decoded into it.
func (c *Client) Do(ctx context.Context, method, path string, body, response any) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("jules api: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	url := c.BaseURL + path

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("jules api: create request: %w", err)
	}

	req.Header.Set("x-goog-api-key", c.apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("jules api: execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("jules api: read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseErrorResponse(resp.StatusCode, respBody)
	}

	if response != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, response); err != nil {
			return fmt.Errorf("jules api: unmarshal response: %w", err)
		}
	}

	return nil
}

// parseErrorResponse parses a non-2xx response body into an *APIError.
func parseErrorResponse(statusCode int, body []byte) error {
	var errResp errorResponse
	if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error != nil {
		errResp.Error.HTTPStatusCode = statusCode
		return errResp.Error
	}
	// If we cannot parse the error body, return a generic error.
	return &APIError{
		HTTPStatusCode: statusCode,
		Code:           statusCode,
		Message:        string(body),
		Status:         http.StatusText(statusCode),
	}
}

// appendPagination helper to centralize page size and page token formatting.
func appendPagination(q url.Values, pageSize int, pageToken string) {
	if pageSize > 0 {
		q.Set("pageSize", strconv.Itoa(pageSize))
	}
	if pageToken != "" {
		q.Set("pageToken", pageToken)
	}
}
