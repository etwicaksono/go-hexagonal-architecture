package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type ExampleDbInterface interface {
	FindAllTextMessage(ctx context.Context) ([]entity.MessageTextItem, error)
	UpsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (entity.BulkWriteResult, error)
}
