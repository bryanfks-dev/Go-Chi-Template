package userrepository

import "skeleton/pkg/logger"

type UserRepository struct {
	logger *logger.Logger
}

func NewUserRepository(logger *logger.Logger) *UserRepository {
	return &UserRepository{
		logger: logger,
	}
}
