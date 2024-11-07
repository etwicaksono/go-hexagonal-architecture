package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (a authenticationApp) Login(ctx context.Context, request entity.LoginRequest) (entity.UserAccess, error) {
	//TODO implement me
	panic("implement me")
}
