//go:build wireinject
// +build wireinject

package di

import (
	"skeleton/infra/ent"
	authdelivery "skeleton/internal/api/auth/delivery"
	errordelivery "skeleton/internal/api/error/delivery"
	etcdelivery "skeleton/internal/api/etc/delivery"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"

	"github.com/google/wire"
)

type Container struct {
	EtcHandler   *etcdelivery.EtcHandler
	ErrorHandler *errordelivery.ErrorHandler
	AuthHandler  *authdelivery.AuthHandler
}

func NewContainer(
	logger *logger.Logger,
	db *ent.Client,
	sec *security.Security,
) *Container {
	wire.Build(
		DeliverySet,
		wire.Struct(new(Container), "*"),
	)
	return &Container{}
}
