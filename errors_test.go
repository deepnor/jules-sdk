package jules

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

// apiErr is a helper to create an *APIError with the given HTTP status code.
func apiErr(statusCode int) *APIError {
	return &APIError{
		HTTPStatusCode: statusCode,
		Code:           statusCode,
		Message:        http.StatusText(statusCode),
		Status:         http.StatusText(statusCode),
	}
}

// --- IsUnauthorized ---

func TestIsUnauthorized_WithUnauthorizedAPIError(t *testing.T) {
	err := apiErr(http.StatusUnauthorized)
	if !IsUnauthorized(err) {
		t.Errorf("IsUnauthorized(%v) = false, want true", err)
	}
}

func TestIsUnauthorized_Wrapped(t *testing.T) {
	err := fmt.Errorf("operation failed: %w", apiErr(http.StatusUnauthorized))
	if !IsUnauthorized(err) {
		t.Errorf("IsUnauthorized(wrapped 401) = false, want true")
	}
}

func TestIsUnauthorized_WithOtherStatusCode(t *testing.T) {
	otherCodes := []int{
		http.StatusOK,
		http.StatusNotFound,
		http.StatusTooManyRequests,
		http.StatusForbidden,
		http.StatusInternalServerError,
	}
	for _, code := range otherCodes {
		if IsUnauthorized(apiErr(code)) {
			t.Errorf("IsUnauthorized(apiErr(%d)) = true, want false", code)
		}
	}
}

func TestIsUnauthorized_WithNilError(t *testing.T) {
	if IsUnauthorized(nil) {
		t.Error("IsUnauthorized(nil) = true, want false")
	}
}

func TestIsUnauthorized_WithNonAPIError(t *testing.T) {
	err := errors.New("a plain error")
	if IsUnauthorized(err) {
		t.Errorf("IsUnauthorized(%v) = true, want false", err)
	}
}

// --- IsNotFound ---

func TestIsNotFound_WithNotFoundAPIError(t *testing.T) {
	err := apiErr(http.StatusNotFound)
	if !IsNotFound(err) {
		t.Errorf("IsNotFound(%v) = false, want true", err)
	}
}

func TestIsNotFound_Wrapped(t *testing.T) {
	err := fmt.Errorf("operation failed: %w", apiErr(http.StatusNotFound))
	if !IsNotFound(err) {
		t.Errorf("IsNotFound(wrapped 404) = false, want true")
	}
}

func TestIsNotFound_WithOtherStatusCode(t *testing.T) {
	otherCodes := []int{
		http.StatusOK,
		http.StatusUnauthorized,
		http.StatusTooManyRequests,
		http.StatusForbidden,
		http.StatusInternalServerError,
	}
	for _, code := range otherCodes {
		if IsNotFound(apiErr(code)) {
			t.Errorf("IsNotFound(apiErr(%d)) = true, want false", code)
		}
	}
}

func TestIsNotFound_WithNilError(t *testing.T) {
	if IsNotFound(nil) {
		t.Error("IsNotFound(nil) = true, want false")
	}
}

func TestIsNotFound_WithNonAPIError(t *testing.T) {
	err := errors.New("a plain error")
	if IsNotFound(err) {
		t.Errorf("IsNotFound(%v) = true, want false", err)
	}
}

// --- IsRateLimited ---

func TestIsRateLimited_WithTooManyRequestsAPIError(t *testing.T) {
	err := apiErr(http.StatusTooManyRequests)
	if !IsRateLimited(err) {
		t.Errorf("IsRateLimited(%v) = false, want true", err)
	}
}

func TestIsRateLimited_Wrapped(t *testing.T) {
	err := fmt.Errorf("operation failed: %w", apiErr(http.StatusTooManyRequests))
	if !IsRateLimited(err) {
		t.Errorf("IsRateLimited(wrapped 429) = false, want true")
	}
}

func TestIsRateLimited_WithOtherStatusCode(t *testing.T) {
	otherCodes := []int{
		http.StatusOK,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusForbidden,
		http.StatusInternalServerError,
	}
	for _, code := range otherCodes {
		if IsRateLimited(apiErr(code)) {
			t.Errorf("IsRateLimited(apiErr(%d)) = true, want false", code)
		}
	}
}

func TestIsRateLimited_WithNilError(t *testing.T) {
	if IsRateLimited(nil) {
		t.Error("IsRateLimited(nil) = true, want false")
	}
}

func TestIsRateLimited_WithNonAPIError(t *testing.T) {
	err := errors.New("a plain error")
	if IsRateLimited(err) {
		t.Errorf("IsRateLimited(%v) = true, want false", err)
	}
}

// --- APIError.Error() ---

func TestAPIError_Error_Format(t *testing.T) {
	e := &APIError{
		HTTPStatusCode: 401,
		Code:           16,
		Message:        "invalid credentials",
		Status:         "UNAUTHENTICATED",
	}
	want := "jules api: 401 UNAUTHENTICATED: invalid credentials"
	if got := e.Error(); got != want {
		t.Errorf("APIError.Error() = %q, want %q", got, want)
	}
}
