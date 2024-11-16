package authentication_core

import (
	"context"
)

func (a authenticationCore) Logout(ctx context.Context, accessKey string) (err error) {
	err = a.cache.DeleteAuthenticatedToken(ctx, accessKey)
	if err != nil {
		return
	}
	return
}
