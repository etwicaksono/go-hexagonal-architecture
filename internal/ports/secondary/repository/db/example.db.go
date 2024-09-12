package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleDbInterface interface {
	FindAllTextMessage(ctx context.Context) ([]entity.MessageTextItem, error)
	FindAllMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error)
	InsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (entity.BulkWriteResult, error)
	InsertMultimediaMessage(ctx context.Context, objs []entity.MessageMultimediaItem) (entity.BulkWriteResult, error)
}
