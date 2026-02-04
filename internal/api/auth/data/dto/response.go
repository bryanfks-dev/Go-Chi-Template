package authdto

import (
	"skeleton/infra/ent"
	userdto "skeleton/internal/api/user/data/dto"
)

type PostAuthTokenRefreshResponseDTO struct {
	AccessToken string `json:"access_token"`
	CSRFToken   string `json:"csrf_token,omitempty"`
	XSRFToken   string `json:"xsrf_token"`
}

func NewPostAuthTokenRefreshResponseDTO(
	accessToken string,
	csrfToken string,
	xsrfToken string,
) *PostAuthTokenRefreshResponseDTO {
	return &PostAuthTokenRefreshResponseDTO{
		AccessToken: accessToken,
		CSRFToken:   csrfToken,
		XSRFToken:   xsrfToken,
	}
}

type PostAuthLoginResponseDTO struct {
	User         *userdto.UserDTO `json:"user"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token,omitempty"`
	CSRFToken    string           `json:"csrf_token,omitempty"`
	XSRFToken    string           `json:"xsrf_token"`
}

func NewPostAuthLoginResponseDTO(
	user *ent.User,
	accessToken string,
	refreshToken string,
	csrfToken string,
	xsrfToken string,
) *PostAuthLoginResponseDTO {
	userDto := userdto.NewUserDTO(user)
	return &PostAuthLoginResponseDTO{
		User:         userDto,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CSRFToken:    csrfToken,
		XSRFToken:    xsrfToken,
	}
}
