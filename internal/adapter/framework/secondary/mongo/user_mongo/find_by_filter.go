package user_mongo

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
)

func (e userMongo) FindByFilter(ctx context.Context, filter entity.UserFindFilter) (entity.User, error) {
	getByFilter := entity.UserGetFilter{
		Active: filter.Active,
	}

	if filter.ID.Valid {
		getByFilter.IDs = []string{filter.ID.String}
	}
	if filter.Email.Valid {
		getByFilter.Emails = []string{filter.Email.String}
	}
	if filter.Name.Valid {
		getByFilter.Names = []string{filter.Name.String}
	}
	if filter.Username.Valid {
		getByFilter.Usernames = []string{filter.Username.String}
	}

	users, err := e.GetByFilter(ctx, getByFilter)
	if err != nil {
		return entity.User{}, err
	}
	if len(users) == 0 {
		return entity.User{}, errors2.ErrNoData
	}
	return users[0], nil
}
