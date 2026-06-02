package apperror

import "errors"

var (
	ErrRouteNotFound      = errors.New("ROUTE_NOT_FOUND")
	ErrMethodNotAllowed   = errors.New("METHOD_NOT_ALLOWED")
	ErrMissingRequestBody = errors.New("MISSING_REQUEST_BODY")
	ErrInvalidRequestBody = errors.New("INVALID_REQUEST_BODY")
	ErrInternalServer     = errors.New("INTERNAL_SERVER_ERROR")
	ErrValidationFailed   = errors.New("VALIDATION_FAILED")
	ErrAuthSessionError   = errors.New("AUTH_SESSION_ERROR")
)
