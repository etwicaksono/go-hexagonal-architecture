package app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleMessageAppInterface interface {
	GetTextMessage(ctx context.Context) (result []entity.MessageTextItem, err error)
	SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) (err error)
	SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) (err error)
	GetMultimediaMessage(ctx context.Context) (result []entity.MessageMultimediaItem, err error)
}
