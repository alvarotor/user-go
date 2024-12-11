package models

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserNotAuthenticated = errors.New("user not authenticated")
)
