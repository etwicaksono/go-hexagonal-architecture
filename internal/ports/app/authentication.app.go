package app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type AuthenticationAppInterface interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
	Login(ctx context.Context, request entity.LoginRequest) (entity.UserAccess, error)
	Logout(ctx context.Context) error
	Refresh(ctx context.Context) (entity.UserAccess, error)
}
