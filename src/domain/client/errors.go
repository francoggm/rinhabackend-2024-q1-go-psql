package client

import "errors"

var (
	ErrNotFound          = errors.New("user not found")
	ErrInsufficientLimit = errors.New("insufficient limit")
)
