package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleMessageDbInterface interface {
	FindAllTextMessage(ctx context.Context) (result []entity.MessageTextItem, err error)
	FindAllMultimediaMessage(ctx context.Context) (result []entity.MessageMultimediaItem, err error)
	InsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (result entity.BulkWriteResult, err error)
	InsertMultimediaMessage(ctx context.Context, objs []entity.MessageMultimediaItem) (result entity.BulkWriteResult, err error)
}
