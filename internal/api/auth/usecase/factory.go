package authusecase

import (
	"skeleton/infra/ent"
	authrepository "skeleton/internal/api/auth/repository"
	userrepository "skeleton/internal/api/user/repository"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"
)

type AuthUsecase struct {
	logger   *logger.Logger
	sec      *security.Security
	db       *ent.Client
	userRepo *userrepository.UserRepository
	authRepo *authrepository.AuthRepository
}

func NewAuthUsecase(
	logger *logger.Logger,
	sec *security.Security,
	db *ent.Client,
	userRepo *userrepository.UserRepository,
	authRepo *authrepository.AuthRepository,
) *AuthUsecase {
	return &AuthUsecase{
		logger:   logger,
		sec:      sec,
		db:       db,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}
