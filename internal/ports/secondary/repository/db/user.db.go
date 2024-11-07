package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type UserDbInterface interface {
	CreateUser(ctx context.Context, objs []entity.User) (entity.BulkWriteResult, error)
	FindByFilter(ctx context.Context, filter entity.UserFindFilter) (entity.User, error)
	GetByFilter(ctx context.Context, filter entity.UserGetFilter) ([]entity.User, error)
}
