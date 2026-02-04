package authdelivery

import authusecase "skeleton/internal/api/auth/usecase"

type AuthHandler struct {
	authUc *authusecase.AuthUsecase
}

func NewAuthHandler(
	authUc *authusecase.AuthUsecase,
) *AuthHandler {
	return &AuthHandler{
		authUc: authUc,
	}
}
