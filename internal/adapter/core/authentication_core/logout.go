package authentication_core

import (
	"context"
	"time"
)

func (a authenticationCore) Logout(ctx context.Context, accessKey string, expiredAt time.Time) (err error) {
	err = a.cache.DeleteAuthenticatedToken(ctx, accessKey)
	if err != nil {
		return
	}
	err = a.cache.SetBlacklistedToken(ctx, accessKey, expiredAt) // TODO: remove this step
	return
}
