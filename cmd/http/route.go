package main

import (
	"fmt"
	"net/http"
	"skeleton/di"
	etcapi "skeleton/internal/api/etc"
	"skeleton/pkg/config"
	"skeleton/pkg/constant"
	apperror "skeleton/pkg/error"
	"skeleton/pkg/logger"
	"skeleton/pkg/server"
	"skeleton/pkg/utils"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func _registerAPIRoutes(
	srv *server.Server,
	deliveryContainer *di.DeliveryContainer,
) {
	if srv == nil {
		panic("srv cannot be nil")
	}

	if deliveryContainer == nil {
		panic("deliveryContainer cannot be nil")
	}

	apiRouter := chi.NewRouter()
	srv.Router.Mount("/api", apiRouter)
}

func _registerPublicRoutes(
	srv *server.Server,
	deliveryContainer *di.DeliveryContainer,
) {
	if srv == nil {
		panic("srv cannot be nil")
	}

	if deliveryContainer == nil {
		panic("deliveryContainer cannot be nil")
	}

	publicRouter := chi.NewRouter()
	srv.Router.Mount("/public", publicRouter)

	etcapi.RegisterRoutes(publicRouter, deliveryContainer.EtcHandler)
}

func _registerErrorRoutes(srv *server.Server) {
	if srv == nil {
		panic("srv cannot be nil")
	}

	srv.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteErrorJSONResponse(
			w,
			http.StatusNotFound,
			apperror.ErrRouteNotFound,
		)
	})

	srv.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteErrorJSONResponse(
			w,
			http.StatusMethodNotAllowed,
			apperror.ErrMethodNotAllowed,
		)
	})
}

func registerRoutes(
	srv *server.Server,
	logger *logger.Logger,
	env config.Environment,
) {
	if srv == nil {
		panic("srv cannot be nil")
	}

	if logger == nil {
		panic("logger cannot be nil")
	}

	deliveryContainer := di.NewDeliveryContainer()
	_registerAPIRoutes(srv, deliveryContainer)
	_registerPublicRoutes(srv, deliveryContainer)
	_registerErrorRoutes(srv)

	if env == config.EnvironmentDevelopment {
		fmt := fmt.Sprintf(
			"API Docs will be available at %s%s",
			srv.Address(),
			constant.APIDocsRoute,
		)
		logger.Info(fmt)

		srv.Router.Mount(constant.APIDocsRoute, httpSwagger.WrapHandler)
		srv.Router.Mount("/debug", middleware.Profiler())
	}
}
