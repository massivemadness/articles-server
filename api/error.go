package api

import "errors"

var (
	ErrNotFound = errors.New("article_not_found")
	ErrUnknown  = errors.New("unknown_error")
)
