package authapi

import (
	authdelivery "skeleton/internal/api/auth/delivery"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, authHandler *authdelivery.AuthHandler) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)

		r.Route("/token", func(r chi.Router) {
			r.Post("/refresh", authHandler.RefreshAccessToken)
		})
	})
}
