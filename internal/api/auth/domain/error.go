package authdomain

import "errors"

var (
	ErrInvalidToken            = errors.New("invalid_token")
	ErrInvalidCredentials      = errors.New("invalid_credentials")
	ErrGenerateTokenFailed     = errors.New("generate_token_failed")
	ErrGetAuthSessionFailed    = errors.New("get_auth_session_failed")
	ErrCreateAuthSessionFailed = errors.New("create_auth_session_failed")
	ErrDeleteAuthSessionFailed = errors.New("delete_auth_session_failed")
)
