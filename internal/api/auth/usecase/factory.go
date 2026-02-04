package authusecase

import (
	"skeleton/infra/ent"
	userrepository "skeleton/internal/api/user/repository"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"
)

type AuthUsecase struct {
	logger   *logger.Logger
	sec      *security.Security
	db       *ent.Client
	userRepo *userrepository.UserRepository
}

func NewAuthUsecase(
	logger *logger.Logger,
	sec *security.Security,
	db *ent.Client,
	userRepo *userrepository.UserRepository,
) *AuthUsecase {
	return &AuthUsecase{
		logger:   logger,
		sec:      sec,
		db:       db,
		userRepo: userRepo,
	}
}
