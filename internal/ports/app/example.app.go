package app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleAppInterface interface {
	GetTextMessage(ctx context.Context) ([]entity.MessageTextItem, error)
	SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error
	SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error
	GetMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error)
}
