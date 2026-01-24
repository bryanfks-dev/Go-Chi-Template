package server

import (
	"net/http"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router      *chi.Mux
	server      *http.Server
	config      config.ServerProperties
	logger      *logger.Logger
	environment config.Environment
}

func NewServer(
	cfg config.ServerProperties,
	logger *logger.Logger,
	env config.Environment,
) Server {
	addr := "0.0.0.0:" + strconv.Itoa(cfg.Port)
	router := chi.NewRouter()

	return Server{
		logger: logger,
		config: cfg,
		Router: router,
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		environment: env,
	}
}
