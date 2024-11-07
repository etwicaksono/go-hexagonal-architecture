package core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type AuthenticationCoreInterface interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
}
