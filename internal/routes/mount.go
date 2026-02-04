package routes

import (
	"fmt"
	"net/http"
	authapi "skeleton/internal/api/auth"
	errordelivery "skeleton/internal/api/error/delivery"
	etcapi "skeleton/internal/api/etc"
	"skeleton/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpswagger "github.com/swaggo/http-swagger"
)

func (r *Route) MountRoutes() {
	r.mountAssetRoutes()
	r.mountErrorRoutes(r.container.ErrorHandler)
	r.mountAPIRoutes()
	r.mountPublicAPIRoutes()

	if r.env == config.EnvironmentDevelopment {
		r.mountDocsRoutes()
		r.srv.Router.Mount("/debug", middleware.Profiler())
	}
}

func (r *Route) mountAPIRoutes() {
	apiRouter := chi.NewRouter()
	authapi.RegisterRoutes(apiRouter, r.container.AuthHandler)

	r.srv.Router.Mount("/api", apiRouter)
}

func (r *Route) mountPublicAPIRoutes() {
	publicRouter := chi.NewRouter()
	etcapi.RegisterRoutes(publicRouter, r.container.EtcHandler)

	r.srv.Router.Mount("/public/api", publicRouter)
}

func (r *Route) mountAssetRoutes() {
	fs := http.FileServer(http.Dir("./public"))
	r.srv.Router.Handle("/assets/", http.StripPrefix("/assets/", fs))
}

func (r *Route) mountErrorRoutes(handler *errordelivery.ErrorHandler) {
	r.srv.Router.NotFound(handler.NotFound)
	r.srv.Router.MethodNotAllowed(handler.MethodNotAllowed)
}

func (r *Route) mountDocsRoutes() {
	fmt := fmt.Sprintf(
		"API Docs will be available at %s%s",
		r.srv.Address(),
		APIDocsRoute,
	)
	r.logger.Info(fmt)

	r.srv.Router.Mount(APIDocsRoute, httpswagger.WrapHandler)
}
