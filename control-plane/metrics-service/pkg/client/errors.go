package client

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewAPIErrorFromResponse(resp *http.Response, err error) *APIError {
	statusCode := http.StatusInternalServerError
	if resp != nil {
		statusCode = resp.StatusCode
	}

	msg := "unknown error"
	if err != nil {
		msg = err.Error()
	}

	return &APIError{
		StatusCode: statusCode,
		Message:    msg,
	}
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, http.StatusText(e.StatusCode))
}
