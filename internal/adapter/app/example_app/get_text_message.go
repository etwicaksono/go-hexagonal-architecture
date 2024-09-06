package example_app

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (e exampleApp) GetTextMessage() ([]entity.MessageTextItem, error) {
	messages, err := e.core.GetTextMessage()
	if err != nil {
		slog.ErrorContext(e.ctx, "Error on getting text message", slog.String("error", err.Error()))
		return nil, err
	}

	return messages, nil
}
