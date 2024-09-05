package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/injector"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/grpc"
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

	/*
		Server initialization
	*/
	// Rest app initialization
	restApp := injector.RestProvider(ctx)

	// Grpc app initialization
	grpcHandler := injector.GrpcHandlerProvider(ctx)
	grpcApp := grpc.NewGrpcAdapter(
		ctx,
		fmt.Sprintf("%s:%d", cfg.App.GrpcHost, cfg.App.GrpcPort),
		grpcHandler,
	)

	/*
		Start server
	*/
	// Run fiber rest server
	go func() {
		slog.InfoContext(ctx, "Starting rest server...")
		err := restApp.Listen(fmt.Sprintf("%s:%d", cfg.App.RestHost, cfg.App.RestPort))
		if err != nil {
			slog.ErrorContext(ctx, "Failed to start rest server", slog.String(entity.Error, err.Error()))
		}
	}()

	// Run grpc server
	go func() {
		err = grpcApp.Run()
		if err != nil {
			slog.ErrorContext(ctx, "Failed to start grpc server", slog.String(entity.Error, err.Error()))
			shutdown <- err
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
