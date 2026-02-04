package routes

import (
	authapi "skeleton/internal/api/auth"

	"github.com/go-chi/chi/v5"
)

func (r *Route) mountAPIRoutes() {
	apiRouter := chi.NewRouter()
	authapi.RegisterRoutes(apiRouter, r.container.AuthHandler)

	r.srv.Router.Mount("/api", apiRouter)
}
