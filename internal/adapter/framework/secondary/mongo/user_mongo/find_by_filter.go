package user_mongo

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
)

func (e userMongo) FindByFilter(ctx context.Context, filter entity.UserFindFilter) (entity.User, error) {
	users, err := e.GetByFilter(ctx, entity.UserGetFilter{
		IDs:       []string{filter.ID},
		Emails:    []string{filter.Email},
		Names:     []string{filter.Name},
		Usernames: []string{filter.Username},
		Active:    filter.Active,
	})
	if err != nil {
		return entity.User{}, err
	}
	if len(users) == 0 {
		return entity.User{}, errors2.ErrNoData
	}
	return users[0], nil
}
