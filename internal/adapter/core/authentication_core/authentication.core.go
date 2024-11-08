package authentication_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type authenticationCore struct {
	db     db.UserDbInterface
	config config.Config
}

func NewAuthenticationCore(
	db db.UserDbInterface,
	config config.Config,
) core.AuthenticationCoreInterface {
	return &authenticationCore{
		db:     db,
		config: config,
	}
}
