package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	StatusCode int
	StatusText string
	Detail     string
	Stack      string
}

type ErrorResponse struct {
	StatusText string `json:"status_text"`
	Detail     string `json:"detail"`
}
type SuccessResponse struct {
	StatusText string `json:"status_text"`
}

func New(statusCode int, detail string, stack string) error {
	return &Error{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Detail:     detail,
		Stack:      stack,
	}
}
func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s %s \nStackTrace:\n%s", e.StatusCode, e.StatusText, e.Detail, e.Stack)
}
func Is(err1 error, err2 error) bool {
	return errors.Is(err1, err2)
}
