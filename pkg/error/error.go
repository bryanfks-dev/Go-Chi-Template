package apperror

import "errors"

var (
	ErrRouteNotFound      = errors.New("route_not_found")
	ErrMethodNotAllowed   = errors.New("method_not_allowed")
	ErrMissingRequestBody = errors.New("missing_request_body")
	ErrInvalidRequestBody = errors.New("invalid_request_body")
	ErrInternalServer     = errors.New("internal_server_error")
	ErrValidationFailed   = errors.New("validation_failed")
)
