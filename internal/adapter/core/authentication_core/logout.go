package authentication_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"time"
)

func (a authenticationCore) Logout(ctx context.Context, authToken entity.AuthToken) (err error) {
	// Calculate the remaining time-to-live (TTL) for the token
	ttl := time.Until(authToken.ExpiredAt)
	a.cache.SetToken(ctx, authToken.AccessKey, ttl)
}
