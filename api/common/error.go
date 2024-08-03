package common

import "errors"

type HttpError struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}

var (
	ErrInvalid  = errors.New("invalid_argument")
	ErrNotFound = errors.New("not_found")
	ErrUnknown  = errors.New("unknown_error")
)
