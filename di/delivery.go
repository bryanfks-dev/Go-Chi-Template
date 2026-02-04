package di

import (
	authdelivery "skeleton/internal/api/auth/delivery"
	errordelivery "skeleton/internal/api/error/delivery"
	etcdelivery "skeleton/internal/api/etc/delivery"

	"github.com/google/wire"
)

type DeliveryContainer struct {
	EtcHandler   *etcdelivery.EtcHandler
	ErrorHandler *errordelivery.ErrorHandler
	AuthHandler  *authdelivery.AuthHandler
}

var DeliverySet = wire.NewSet(
	UsecaseSet,
	etcdelivery.NewEtcHandler,
	errordelivery.NewErrorHandler,
	authdelivery.NewAuthHandler,
	wire.Struct(new(DeliveryContainer), "*"),
)
