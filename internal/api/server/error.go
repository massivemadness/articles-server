package server

import "errors"

type HttpError struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    int    `json:"error_code"`
}

var (
	ErrDecode   = errors.New("decode_error")
	ErrNotFound = errors.New("not_found")
	ErrUnknown  = errors.New("unknown_error")
)
