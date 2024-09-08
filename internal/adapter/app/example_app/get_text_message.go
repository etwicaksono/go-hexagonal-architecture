package example_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (e exampleApp) GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	messages, err := e.core.GetTextMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error on getting text message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	return messages, nil
}
