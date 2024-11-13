package user_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mysql/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/string_util"
	"log/slog"
)

func (u userMysql) GetByFilter(ctx context.Context, filter entity.UserGetFilter) (result []entity.User, err error) {
	var userMysqlModels []model.User

	pipeline := u.client.Table(u.table).Where("deleted_at is null")

	if len(filter.IDs) > 0 {
		pipeline = pipeline.Where("LOWER(id) in (?)", string_util.ToLowerSlice(filter.IDs))
	}
	if len(filter.Emails) > 0 {
		pipeline = pipeline.Where("LOWER(email) in (?)", string_util.ToLowerSlice(filter.Emails))
	}
	if len(filter.Names) > 0 {
		pipeline = pipeline.Where("LOWER(name) in (?)", string_util.ToLowerSlice(filter.Names))
	}
	if len(filter.Usernames) > 0 {
		pipeline = pipeline.Where("LOWER(username) in (?)", string_util.ToLowerSlice(filter.Usernames))
	}
	if filter.Active.Valid {
		pipeline = pipeline.Where("active = ?", filter.Active.Bool)
	}

	tx := pipeline.Find(&userMysqlModels)
	if tx.Error != nil {
		slog.ErrorContext(ctx, "Failed to GetByFilter user", slog.String(entity.Error, tx.Error.Error()))
		return nil, tx.Error
	}

	var users []entity.User
	for _, userMysqlModel := range userMysqlModels {
		users = append(users, userMysqlModel.ToEntity())
	}
	return users, nil
}
