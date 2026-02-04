package authdto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type PostAuthLoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r PostAuthLoginRequestDTO) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.Email,
			validation.Required.Error("email_required"),
			is.Email.Error("email_format_invalid"),
		),
		validation.Field(
			&r.Password,
			validation.Required.Error("password_required"),
			validation.Length(8, 100).Error("password_length_invalid"),
		),
	)
}

type PostAuthTokenRefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (r PostAuthTokenRefreshRequestDTO) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.RefreshToken,
			validation.Required.Error("refresh_token_required"),
		),
	)
}

type PostAuthLogoutRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

func (r PostAuthLogoutRequestDTO) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.RefreshToken,
			validation.Required.Error("refresh_token_required"),
		),
	)
}
