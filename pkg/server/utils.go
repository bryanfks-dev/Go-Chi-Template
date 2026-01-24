package server

import (
	"os"
	"os/signal"
	"syscall"
)

func setupGracefulShutdown(onShutdownSignal func()) <-chan any {
	shutdown := make(chan any)
	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
