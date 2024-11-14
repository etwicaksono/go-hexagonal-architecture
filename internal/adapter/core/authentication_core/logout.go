package authentication_core

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
)

func (a authenticationCore) Logout(ctx context.Context, authToken entity.AuthToken) (err error) {
	err = a.cache.SetAuthToken(
		ctx,
		fmt.Sprintf("%s:%s", constants.BlackListedTokenRedisPrefix, authToken.AccessKey),
		model.TokenData{
			AccessKey:   authToken.AccessKey,
			ExpiredDate: authToken.ExpiredAt,
		},
	)
	return
}
