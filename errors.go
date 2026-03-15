package jules

import (
	"errors"
	"fmt"
	"net/http"
)

// APIError represents a structured error response from the Jules API.
type APIError struct {
	// HTTPStatusCode is the HTTP status code of the response.
	HTTPStatusCode int `json:"-"`
	// Code is the API error code from the response body.
	Code int `json:"code"`
	// Message is the human-readable error message.
	Message string `json:"message"`
	// Status is the canonical error status string (e.g., "INVALID_ARGUMENT").
	Status string `json:"status"`
}

// Error returns a human-readable representation of the API error.
func (e *APIError) Error() string {
	return fmt.Sprintf("jules api: %d %s: %s", e.HTTPStatusCode, e.Status, e.Message)
}

// errorResponse is the JSON envelope for API error responses.
type errorResponse struct {
	Error *APIError `json:"error"`
}

// IsNotFound reports whether the error is, or wraps, a 404 Not Found API response.
func IsNotFound(err error) bool {
	var e *APIError
	if errors.As(err, &e) {
		return e.HTTPStatusCode == http.StatusNotFound
	}
	return false
}

// IsRateLimited reports whether the error is, or wraps, a 429 Too Many Requests
// API response.
func IsRateLimited(err error) bool {
	var e *APIError
	if errors.As(err, &e) {
		return e.HTTPStatusCode == http.StatusTooManyRequests
	}
	return false
}

// IsUnauthorized reports whether the error is, or wraps, a 401 Unauthorized
// API response.
func IsUnauthorized(err error) bool {
	var e *APIError
	if errors.As(err, &e) {
		return e.HTTPStatusCode == http.StatusUnauthorized
	}
	return false
}
