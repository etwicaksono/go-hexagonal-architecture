package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type UserDbInterface interface {
	CreateUser(ctx context.Context, objs []entity.User) (result entity.BulkWriteResult, err error)
	FindByFilter(ctx context.Context, filter entity.UserFindFilter) (result entity.User, err error)
	GetByFilter(ctx context.Context, filter entity.UserGetFilter) (result []entity.User, err error)
}
