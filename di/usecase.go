package di

import (
	authusecase "skeleton/internal/api/auth/usecase"

	"github.com/google/wire"
)

type UsecaseContainer struct {
	AuthUC *authusecase.AuthUsecase
}

var UsecaseSet = wire.NewSet(
	RepositorySet,
	authusecase.NewAuthUsecase,
	wire.Struct(new(UsecaseContainer), "*"),
)
