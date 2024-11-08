package core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleMessageCoreInterface interface {
	GetTextMessage(ctx context.Context) (result []entity.MessageTextItem, err error)
	SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) (err error)
	SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) (err error)
	GetMultimediaMessage(ctx context.Context) (result []entity.MessageMultimediaItem, err error)
}
