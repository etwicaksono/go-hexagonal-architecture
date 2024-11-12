package example_message_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (u userMongo) InsertMultimediaMessage(ctx context.Context, objs []entity.MessageMultimediaItem) (result entity.BulkWriteResult, err error) {
	//TODO implement me
	panic("implement me")
}
