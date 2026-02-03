package server

import (
	"context"
	"fmt"
	"net/http"
)

func (s *Server) Start() {
	format := fmt.Sprintf("Starting HTTP server on %s", s.Address())
	s.logger.Info(format)

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic("HTTP server failed to start: " + err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("HTTP server forced to shutdown: " + err.Error())
		return
	}
	s.logger.Info("HTTP server stopped gracefully")
}
