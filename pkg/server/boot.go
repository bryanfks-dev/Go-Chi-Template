package server

import (
	"context"
	"fmt"
	"net/http"
	"skeleton/pkg/config"
)

func (s *Server) Start() {
	format := fmt.Sprintf("Starting HTTP server on %s", s.server.Addr)
	s.logger.Info(format)

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic("HTTP server failed to start: " + err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		panic("HTTP server forced to shutdown: " + err.Error())
	}
	s.logger.Info("HTTP server stopped gracefully")

	s.logger.Info("Executing post shutdown tasks")
	if s.environment == config.EnvironmentProduction {
		if err := s.logger.Sync(); err != nil {
			panic("Failed to flush logs during shutdown: " + err.Error())
		}
	}

	s.logger.Info("Graceful shutdown complete")
}
