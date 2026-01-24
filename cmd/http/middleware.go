package main

import (
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/server"
	"skeleton/pkg/server/middleware"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func registerMiddlewares(
	srv server.Server,
	logger *logger.Logger,
	env config.Environment,
) {
	srv.Router.Use(chimiddleware.RequestID)
	srv.Router.Use(middleware.ZapRequestLoggerMiddleware(logger))
	srv.Router.Use(middleware.CORSMiddleware(env))
}
