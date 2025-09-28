package utils

import (
	"fmt"
	"net/http"
)

// APIError represents an API error
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// NewAPIError creates a new API error
func NewAPIError(code int, message string, details ...string) *APIError {
	err := &APIError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// Common API errors
var (
	ErrBadRequest          = NewAPIError(http.StatusBadRequest, "Bad request")
	ErrUnauthorized        = NewAPIError(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden          = NewAPIError(http.StatusForbidden, "Forbidden")
	ErrNotFound           = NewAPIError(http.StatusNotFound, "Not found")
	ErrInternalServer     = NewAPIError(http.StatusInternalServerError, "Internal server error")
	ErrServiceUnavailable = NewAPIError(http.StatusServiceUnavailable, "Service unavailable")
	ErrTimeout            = NewAPIError(http.StatusRequestTimeout, "Request timeout")
)

// IsRetryableError checks if an error is retryable
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Check if it's an API error
	if apiErr, ok := err.(*APIError); ok {
		switch apiErr.Code {
		case http.StatusInternalServerError,
			http.StatusServiceUnavailable,
			http.StatusRequestTimeout,
			http.StatusBadGateway,
			http.StatusGatewayTimeout:
			return true
		default:
			return false
		}
	}

	// Check for network errors (you might want to add more specific checks)
	return true
}

// WrapError wraps an error with additional context
func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}