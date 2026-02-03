package routes

import (
	etcapi "skeleton/internal/api/etc"

	"github.com/go-chi/chi/v5"
)

func (r *Route) mountPublicRoutes() {
	publicRouter := chi.NewRouter()
	etcapi.RegisterRoutes(publicRouter, r.container.EtcHandler)

	r.srv.Router.Mount("/public", publicRouter)
}
