package example_message_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
)

func (e exampleMessageApp) GetMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error) {
	messages, err := e.core.GetMultimediaMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error on getting multimedia message", slog.String(constants.Error, err.Error()))
		return nil, err
	}

	return messages, nil
}
