package authdomain

import "errors"

var (
	ErrInvalidRefreshToken       = errors.New("invalid_refresh_token")
	ErrGenerateAccessTokenFailed = errors.New("generate_access_token_failed")
	ErrInvalidCredentials        = errors.New("invalid_credentials")
	ErrGenerateTokenFailed       = errors.New("generate_token_failed")
)
