package user_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (u userMysql) FindByFilter(ctx context.Context, filter entity.UserFindFilter) (result entity.User, err error) {
	//TODO implement me
	panic("implement me")
}
