package trongrid

import "errors"

var (
	// ErrEmpty represents an empty result error
	ErrEmpty = errors.New("empty result")

	// API specific errors
	ErrInvalidRequest    = errors.New("invalid request parameters")
	ErrInvalidAddress    = errors.New("invalid tron address")
	ErrInvalidTimeRange  = errors.New("invalid time range")
	ErrRateLimitExceeded = errors.New("API rate limit exceeded")
	ErrUnauthorized      = errors.New("unauthorized API access")
	ErrNetworkError      = errors.New("network communication error")
	ErrServerError       = errors.New("trongrid server error")

	// Validation errors
	ErrMissingAddress   = errors.New("address is required")
	ErrInvalidLimit     = errors.New("invalid limit value")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

// APIError represents a detailed API error
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new APIError
func NewAPIError(code int, message string, err error) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
