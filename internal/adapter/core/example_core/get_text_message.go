package example_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e exampleCore) GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	messages, err := e.db.FindAllTextMessage(ctx)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return messages, nil
}
