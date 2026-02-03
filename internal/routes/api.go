package routes

import "github.com/go-chi/chi/v5"

func (r Route) mountAPIRoutes() {
	apiRouter := chi.NewRouter()
	r.srv.Router.Mount("/api", apiRouter)
}
