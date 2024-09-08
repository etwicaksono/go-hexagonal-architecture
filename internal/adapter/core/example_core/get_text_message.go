package example_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (e exampleCore) GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	messages, err := e.db.FindAllTextMessage(ctx)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		slog.ErrorContext(ctx, "Failed to find all text message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	return messages, nil
}