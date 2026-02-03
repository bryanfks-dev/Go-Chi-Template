package routes

import (
	"skeleton/pkg/config"

	"github.com/go-chi/chi/v5/middleware"
)

func (r *Route) MountRoutes() {
	r.mountErrorRoutes(r.container.ErrorHandler)
	r.mountAPIRoutes()
	r.mountPublicRoutes()

	if r.env == config.EnvironmentDevelopment {
		r.mountDocsRoutes()
		r.srv.Router.Mount("/debug", middleware.Profiler())
	}
}
