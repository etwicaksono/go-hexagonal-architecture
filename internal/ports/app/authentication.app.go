package app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type AuthenticationAppInterface interface {
	Register(ctx context.Context, request entity.RegisterRequest) (err error)
	Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error)
	Logout(ctx context.Context, accessKey string) (err error)
	Refresh(ctx context.Context, request entity.RefreshTokenRequest) (result entity.TokenGenerated, err error)
}
