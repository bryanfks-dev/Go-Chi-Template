//go:build wireinject
// +build wireinject

package di

import (
	etcdelivery "skeleton/internal/api/etc/delivery"

	"github.com/google/wire"
)

type DeliveryContainer struct {
	EtcHandler *etcdelivery.EtcHandler
}

func NewDeliveryContainer() *DeliveryContainer {
	wire.Build(
		etcdelivery.NewEtcHandler,
		wire.Struct(new(DeliveryContainer), "*"),
	)
	return &DeliveryContainer{}
}
