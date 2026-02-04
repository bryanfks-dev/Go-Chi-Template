package authdto

import (
	"skeleton/infra/ent"
	userdto "skeleton/internal/api/user/data/dto"
)

type PostAuthTokenRefreshResponseDTO struct {
	AccessToken string `json:"access_token"`
}

func NewPostAuthTokenRefreshResponseDTO(
	accessToken string,
) *PostAuthTokenRefreshResponseDTO {
	return &PostAuthTokenRefreshResponseDTO{
		AccessToken: accessToken,
	}
}

type PostAuthLoginResponseDTO struct {
	User         *userdto.UserDTO `json:"user"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
}

func NewPostAuthLoginResponseDTO(
	user *ent.User,
	accessToken string,
	refreshToken string,
) *PostAuthLoginResponseDTO {
	userDto := userdto.NewUserDTO(user)
	return &PostAuthLoginResponseDTO{
		User:         userDto,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
