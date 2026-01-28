package di

import "github.com/google/wire"

type UsecaseContainer struct{}

var UsecaseSet = wire.NewSet(
	RepositorySet,
	wire.Struct(new(UsecaseContainer), "*"),
)
