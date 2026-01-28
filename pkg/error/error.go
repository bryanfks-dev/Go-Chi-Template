package apperror

import "errors"

var (
	ErrRouteNotFound    = errors.New("ROUTE_NOT_FOUND")
	ErrMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")
	ErrValidationFailed = errors.New("VALIDATION_ERROR")
)
