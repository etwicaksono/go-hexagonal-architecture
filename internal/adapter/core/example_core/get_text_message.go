package example_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (e exampleCore) GetTextMessage() ([]entity.MessageTextItem, error) {
	messages, err := e.db.FindAllTextMessage(e.ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
