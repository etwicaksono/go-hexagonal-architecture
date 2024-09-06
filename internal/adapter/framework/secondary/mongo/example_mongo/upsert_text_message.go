package example_mongo

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (e exampleMongo) UpsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (entity.BulkWriteResult, error) {
	//TODO implement me
	panic("implement me")
}
