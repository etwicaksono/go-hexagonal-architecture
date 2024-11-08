package authentication_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

func (a authenticationCore) Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error) {
	//TODO implement me
	panic("implement me")
}
