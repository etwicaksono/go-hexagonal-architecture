package errors

import "errors"

var (
	ErrNoData               = errors.New("no data in result")
	ErrFailedToHashPassword = errors.New("failed to hash password")
)
