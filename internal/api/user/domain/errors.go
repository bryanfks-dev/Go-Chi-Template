package userdomain

import "errors"

var (
	ErrUserNotFound  = errors.New("user_not_found")
	ErrGetUserFailed = errors.New("get_user_failed")
)
