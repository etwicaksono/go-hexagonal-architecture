package example_message_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (e exampleMessageCore) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error {
	objs := []entity.MessageTextItem{
		{
			Sender:   request.Sender,
			Receiver: request.Receiver,
			Message:  request.Message,
		},
	}
	_, err := e.db.InsertTextMessage(ctx, objs)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert text message", slog.String(entity.Error, err.Error()))
		return err
	}

	return nil
}
