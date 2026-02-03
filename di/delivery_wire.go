//go:build wireinject
// +build wireinject

package di

import (
	errordelivery "skeleton/internal/api/error/delivery"
	etcdelivery "skeleton/internal/api/etc/delivery"

	"github.com/google/wire"
)

type DeliveryContainer struct {
	ErrorHandler *errordelivery.ErrorHandler
	EtcHandler   *etcdelivery.EtcHandler
}

func NewDeliveryContainer() *DeliveryContainer {
	wire.Build(
		etcdelivery.NewEtcHandler,
		errordelivery.NewErrorHandler,
		wire.Struct(new(DeliveryContainer), "*"),
	)
	return &DeliveryContainer{}
}
