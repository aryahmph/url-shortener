package error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	StatusCode int               `json:"code"`
	Message    string            `json:"message"`
	Errors     map[string]string `json:"errors"`
}

var _ error = &ClientError{}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.StatusCode, e.Message)
}

func NewBadRequestError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Errors:     nil,
	}
}

func NewNotFoundError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Errors:     nil,
	}
}

func NewForbiddenError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
		Errors:     nil,
	}
}

func NewUnauthorizedError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
		Errors:     nil,
	}
}

func NewInternalServerError(msg string) *ClientError {
	return &ClientError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
		Errors:     nil,
	}
}
