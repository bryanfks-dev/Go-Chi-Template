// Function to start the HTTP server
// Usage: go run ./cmd/http

package main

import (
	"context"
	"os"
	"os/signal"
	"skeleton/pkg/application"
	"skeleton/pkg/server"
	"syscall"
	"time"

	_ "skeleton/docs"
	_ "skeleton/pkg/data/dto"
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

	srv := server.NewServer(app.Config.Server, app.Logger, *app.Environment)
	registerMiddlewares(srv, app.Logger, *app.Environment)
	registerRoutes(srv, *app.Environment)

	go srv.Start()
	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		shutdownServerTimeout,
	)
	defer shutdownCancel()
	srv.Shutdown(shutdownCtx)
}
