package di

import (
	authrepository "skeleton/internal/api/auth/repository"
	userrepository "skeleton/internal/api/user/repository"

	"github.com/google/wire"
)

type RepositoryContainer struct {
	UserRepo userrepository.UserRepository
}

var RepositorySet = wire.NewSet(
	authrepository.NewAuthRepository,
	userrepository.NewUserRepository,
	wire.Struct(new(RepositoryContainer), "*"),
)
