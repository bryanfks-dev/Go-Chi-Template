package routes

import (
	"skeleton/di"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/server"
)

type Route struct {
	srv       *server.Server
	env       config.Environment
	logger    *logger.Logger
	container *di.DeliveryContainer
}

func NewRoute(
	srv *server.Server,
	env config.Environment,
	logger *logger.Logger,
	container *di.DeliveryContainer,
) *Route {
	return &Route{
		srv:       srv,
		env:       env,
		logger:    logger,
		container: container,
	}
}
