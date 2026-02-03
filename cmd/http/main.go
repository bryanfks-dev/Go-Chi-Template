// Function to start the HTTP server
// Usage: go run ./cmd/http

package main

import (
	"context"
	"os"
	"os/signal"
	"skeleton/di"
	"skeleton/internal/routes"
	"skeleton/pkg/application"
	"skeleton/pkg/server"
	"skeleton/pkg/server/middleware"
	"syscall"
	"time"

	_ "skeleton/docs"
	_ "skeleton/pkg/data/dto"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

const (
	shutdownServerTimeout = 10 * time.Second
)

// @Title Skeleton API
// @Version 1.0
func main() {
	app := application.NewApplication()
	app.Database.Migrate()
	defer func() {
		app.Logger.Info("Executing post shutdown tasks...")
		app.CleanUp()
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	srv := server.NewServer(app.Config.Server, app.Logger, app.Environment)
	srv.Router.Use(chimiddleware.RequestID)
	srv.Router.Use(middleware.ZapRequestLoggerMiddleware(app.Logger))
	srv.Router.Use(middleware.CORSMiddleware(app.Environment))

	deliveryContainer := di.NewDeliveryContainer()
	routes := routes.NewRoute(
		srv,
		app.Environment,
		app.Logger,
		deliveryContainer,
	)
	routes.MountRoutes()

	go srv.Start()
	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		shutdownServerTimeout,
	)
	defer shutdownCancel()
	srv.Shutdown(shutdownCtx)
}
