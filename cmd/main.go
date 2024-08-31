package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	shutdown := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	//Initiate config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to load cfg", slog.String(entity.Error, err.Error()))
	}

	//Initiate logger
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.String("service", cfg.App.Name),
		slog.String("with-release", cfg.App.Version),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	// Fiber app initialization
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  cfg.App.IdleTimeout,
		WriteTimeout: cfg.App.WriteTimeout,
		ReadTimeout:  cfg.App.ReadTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			status := fiber.ErrInternalServerError.Message
			message := entity.Error

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				status = utils.StatusMessage(e.Code)

				if cfg.App.Env != "production" {
					message = e.Error()
				}
			}

			return ctx.Status(code).JSON(model.Response{
				Code:    code,
				Status:  status,
				Message: message,
			})
		},
	})

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
