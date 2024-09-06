package example_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (e exampleApp) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error {
	err := e.core.SendTextMessage(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on sending text message", slog.String("error", err.Error()))
		return err
	}

	return nil
}
