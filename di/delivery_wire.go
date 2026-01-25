//go:build wireinject
// +build wireinject

package di

import (
	etcdelivery "skeleton/internal/api/etc/delivery"

	"github.com/google/wire"
)

type DeliveryDependency struct {
	EtcHandler *etcdelivery.EtcHandler
}

func NewDeliveryDependency(
	etcHandler *etcdelivery.EtcHandler,
) *DeliveryDependency {
	return &DeliveryDependency{
		EtcHandler: etcHandler,
	}
}

func InitDeliveryDeps() *DeliveryDependency {
	wire.Build(NewDeliveryDependency, etcdelivery.NewEtcHandler)
	return &DeliveryDependency{}
}
