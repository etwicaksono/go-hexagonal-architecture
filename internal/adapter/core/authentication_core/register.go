package authentication_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a authenticationCore) Register(ctx context.Context, request entity.RegisterRequest) (err error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return errors.ErrFailedToHashPassword
	}

	user := entity.User{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Name:      request.Name,
		Username:  request.Username,
		Password:  string(hashedPassword),
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
