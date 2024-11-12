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
	var dbClient *entity.DbClient
	switch cfg.Db.Protocol {
	case "mongodb":
		{
			mongoDb := infrastructure.NewMongoDb(ctx, cfg)
			err = mongoDb.Connect()
			if err != nil {
				slog.ErrorContext(ctx, "Failed to connect to MongoDB", slog.String(entity.Error, err.Error()))
				return
			}
			defer mongoDb.Disconnect()
			dbClient = mongoDb.GetClient()
		}
	default:
		slog.ErrorContext(ctx, "Unsupported database protocol, supported protocol: [mongo]", slog.String("protocol", cfg.Db.Protocol))
		return
	}

	redis := infrastructure.NewRedis(ctx, cfg)
	redis.Connect()
	defer redis.Disconnect()

	/*
	   Server initialization
	*/
	// Rest app initialization
	restApp := injector.RestProvider(ctx, dbClient, redis.GetClient())

	// Grpc app initialization
	grpcHandler := injector.GrpcHandlerProvider(ctx, dbClient)
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
