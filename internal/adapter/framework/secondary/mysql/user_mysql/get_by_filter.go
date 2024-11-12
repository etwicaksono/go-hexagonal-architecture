package user_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (u userMysql) GetByFilter(ctx context.Context, filter entity.UserGetFilter) (result []entity.User, err error) {
	//TODO implement me
	panic("implement me")
}
