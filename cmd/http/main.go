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

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

const (
	shutdownServerTimeout = 10 * time.Second
	databasePingTimeout   = 5 * time.Second
)

// @Title Skeleton API
// @Version 1.0
func main() {
	app := application.NewApplication()
	ctx, cancel := context.WithTimeout(
		context.Background(),
		databasePingTimeout,
	)
	defer func() {
		app.Logger.Info("Executing post shutdown tasks...")
		app.CleanUp()
	}()

	ctx, cancel = signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	srv := server.NewServer(app.Cfg.Server, app.Logger, app.Env)
	srv.Router.Use(chimiddleware.Recoverer)
	srv.Router.Use(chimiddleware.RequestID)
	srv.Router.Use(middleware.ZapRequestLoggerMiddleware(app.Logger))
	srv.Router.Use(middleware.CORSMiddleware(app.Env))

	container := di.NewContainer(app.Logger, app.Db.MasterClient, app.Sec)
	routes := routes.NewRoute(
		srv,
		app.Env,
		app.Logger,
		container,
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
