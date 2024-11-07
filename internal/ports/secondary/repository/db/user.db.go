package db

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type UserDbInterface interface {
	CreateUser(ctx context.Context, objs []entity.User) (entity.BulkWriteResult, error)
}
