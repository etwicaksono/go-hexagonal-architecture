package user_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (u userMysql) CreateUser(ctx context.Context, objs []entity.User) (result entity.BulkWriteResult, err error) {
	//TODO implement me
	panic("implement me")
}
