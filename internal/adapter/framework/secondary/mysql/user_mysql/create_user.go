package user_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mysql/model"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"log/slog"
)

func (u userMysql) CreateUser(ctx context.Context, objs []entity.User) (result entity.BulkWriteResult, err error) {
	if len(objs) == 0 {
		return entity.BulkWriteResult{}, errorsConst.ErrNoObjectToInsert
	}

	userMysqlModels := make([]model.User, len(objs))
	insertedIds := make(map[int64]interface{})
	for i, obj := range objs {
		user := model.FromUserEntity(obj)
		userMysqlModels[i] = user
		insertedIds[int64(i)] = user.ID
	}

	tx := u.client.Table(u.table).Create(&userMysqlModels)
	if tx.Error != nil {
		slog.ErrorContext(ctx, "Failed to BulkWrite user", slog.String(entity.Error, tx.Error.Error()))
		return entity.BulkWriteResult{}, tx.Error
	}
	return entity.BulkWriteResult{
		InsertedCount: tx.RowsAffected,
		UpsertedIDs:   insertedIds,
	}, nil
}
