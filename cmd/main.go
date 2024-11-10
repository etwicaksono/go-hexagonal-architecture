package main

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/infrastructure"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"log/slog"
	"os"
	"os/signal"

	"github.com/etwicaksono/go-hexagonal-architecture/injector"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
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
	   Infrastructure initialization
	*/
	mongoDb := infrastructure.NewMongo(ctx, cfg) // TODO: adjust so it can use other database
	err = mongoDb.Connect()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to connect to MongoDB", slog.String(entity.Error, err.Error()))
		return
	}
	defer mongoDb.Disconnect()

	redis := infrastructure.NewRedis(ctx, cfg)
	redis.Connect()
	defer redis.Disconnect()

	/*
	   Server initialization
	*/
	// Rest app initialization
	restApp := injector.RestProvider(ctx, mongoDb.GetClient(), redis.GetClient())

	// Grpc app initialization
	grpcHandler := injector.GrpcHandlerProvider(ctx, mongoDb.GetClient())
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
			shutdown <- err
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
		slog.ErrorContext(ctx, "Server crashed", slog.String(entity.Error, err.Error()))
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		slog.InfoContext(ctx, "Shutting down server...")
		stop()
	}
}
