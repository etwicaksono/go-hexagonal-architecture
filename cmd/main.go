package main

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	shutdown := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	_, err := config.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to load cfg", slog.String(entity.Error, err.Error()))
	}

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
