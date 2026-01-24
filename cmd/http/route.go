package main

import (
	"net/http"
	etcapi "skeleton/internal/api/etc"
	etcdelivery "skeleton/internal/api/etc/delivery"
	"skeleton/pkg/config"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/server"
	"skeleton/pkg/utils"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func _registerAPIRoutes(srv server.Server) {
	apiRouter := chi.NewRouter()
	srv.Router.Mount("/api", apiRouter)

}

func _registerPublicRoutes(srv server.Server) {
	publicRouter := chi.NewRouter()
	srv.Router.Mount("/public", publicRouter)

	etcapi.RegisterRoutes(publicRouter, etcdelivery.NewEtcHandler())
}

func _registerErrorRoutes(srv server.Server) {
	srv.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteErrorJSONResponse(
			w,
			http.StatusNotFound,
			basedto.NewErrorHTTPResponse("Route not found"),
		)
	})

	srv.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteErrorJSONResponse(
			w,
			http.StatusMethodNotAllowed,
			basedto.NewErrorHTTPResponse("Method not allowed"),
		)
	})
}

func registerRoutes(srv server.Server, env config.Environment) {
	_registerAPIRoutes(srv)
	_registerPublicRoutes(srv)
	_registerErrorRoutes(srv)

	if env == config.EnvironmentDevelopment {
		srv.Router.Mount("/swagger/", httpSwagger.WrapHandler)
		srv.Router.Mount("/debug", middleware.Profiler())
	}
}
