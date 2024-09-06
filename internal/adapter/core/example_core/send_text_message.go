package example_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (e exampleCore) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error {
	objs := []entity.MessageTextItem{
		{
			Sender:   request.Sender,
			Receiver: request.Receiver,
			Message:  request.Message,
		},
	}
	_, err := e.db.UpsertTextMessage(ctx, objs)
	if err != nil {
		return err
	}

	return nil
}
