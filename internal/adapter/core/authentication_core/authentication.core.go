package authentication_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/cache"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
)

type authenticationCore struct {
	db     db.UserDbInterface
	config config.Config
	jwt    *rest_util.Jwt
	cache  cache.AuthCacheInterface
}

func NewAuthenticationCore(
	db db.UserDbInterface,
	config config.Config,
	jwt *rest_util.Jwt,
	cache cache.AuthCacheInterface,
) core.AuthenticationCoreInterface {
	return &authenticationCore{
		db:     db,
		config: config,
		jwt:    jwt,
		cache:  cache,
	}
}
