package etcapi

import (
	etcdelivery "skeleton/internal/api/etc/delivery"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, etcHandler *etcdelivery.EtcHandler) {
	router.Get("/health", etcHandler.HealthCheck)
}
