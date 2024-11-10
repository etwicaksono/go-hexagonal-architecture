package authentication_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils"
	"github.com/google/uuid"
	"time"
)

func (a authenticationCore) Register(ctx context.Context, request entity.RegisterRequest) (err error) {
	// Hash the password
	hashedPassword, err := utils.PasswordGenerate(request.Password)
	if err != nil {
		return errors2.ErrFailedToHashPassword
	}

	// Check email is used
	_, err = a.db.FindByFilter(ctx, entity.UserFindFilter{Email: request.Email})
	if err == nil || !errors.Is(err, errors2.ErrNoData) {
		return errors2.ErrEmailAlreadyUsed
	}

	// Check username is used
	_, err = a.db.FindByFilter(ctx, entity.UserFindFilter{Username: request.Username})
	if err == nil || !errors.Is(err, errors2.ErrNoData) {
		return errors2.ErrUsernameAlreadyUsed
	}

	user := entity.User{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Name:      request.Name,
		Username:  request.Username,
		Password:  hashedPassword,
		Active:    true,
		CreatedAt: time.Now(),
		CreatedBy: request.Username,
	}
	_, err = a.db.CreateUser(ctx, []entity.User{user})
	if err != nil {
		return err
	}

	return
}
