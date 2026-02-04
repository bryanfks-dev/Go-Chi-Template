package authrepository

import "skeleton/pkg/logger"

type AuthRepository struct {
	logger *logger.Logger
}

func NewAuthRepository(logger *logger.Logger) *AuthRepository {
	return &AuthRepository{
		logger: logger,
	}
}
