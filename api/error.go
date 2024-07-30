package api

import "errors"

type HttpError struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}

var (
	ErrNotFound = errors.New("article_not_found")
	ErrUnknown  = errors.New("unknown_error")
)
