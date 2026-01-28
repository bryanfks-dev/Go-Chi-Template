package di

import (
	userrepository "skeleton/internal/api/user/repository"

	"github.com/google/wire"
)

type RepositoryContainer struct {
	UserRepository *userrepository.UserRepository
}

var RepositorySet = wire.NewSet(
	userrepository.NewUserRepository,
	wire.Struct(new(RepositoryContainer), "*"),
)
