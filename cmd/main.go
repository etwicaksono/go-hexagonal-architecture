package main

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/injector"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	shutdown := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Load config
	cfg := config.LoadConfig()

	err := injector.LoggerInit()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize logger", slog.String(entity.Error, err.Error()))
		stop()
	}

	// Fiber app initialization
	fiberApp := injector.FiberProvider()

	// Run fiber rest server
	go func() {
		slog.InfoContext(ctx, "Starting fiber rest server...")
		err := fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.App.RestHost, cfg.App.RestPort))
		if err != nil {
			slog.ErrorContext(ctx, "Failed to start fiber rest server", slog.String(entity.Error, err.Error()))
		}
	}()

	select {
	case err = <-shutdown:
		// Wait throw error
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}
}
