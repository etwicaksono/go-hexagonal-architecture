package example_core

import (
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e exampleCore) GetTextMessage() ([]entity.MessageTextItem, error) {
	messages, err := e.db.FindAllTextMessage(e.ctx)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return messages, nil
}
