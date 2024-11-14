package example_message_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
)

func (e exampleMessageApp) GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	messages, err := e.core.GetTextMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error on getting text message", slog.String(constants.Error, err.Error()))
		return nil, err
	}

	return messages, nil
}
