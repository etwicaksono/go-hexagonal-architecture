package example_message_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (u userMongo) FindAllTextMessage(ctx context.Context) (result []entity.MessageTextItem, err error) {
	//TODO implement me
	panic("implement me")
}
