package example_message_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"log/slog"
)

func (e exampleMessageCore) GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	messages, err := e.db.FindAllTextMessage(ctx)
	if err != nil && !errors.Is(err, errors2.ErrNoData) {
		slog.ErrorContext(ctx, "Failed to find all text message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	return messages, nil
}
