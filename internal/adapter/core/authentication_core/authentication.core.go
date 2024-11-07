package authentication_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type authenticationCore struct {
	db db.UserDbInterface
}

func NewAuthenticationCore(
	db db.UserDbInterface,
) core.AuthenticationCoreInterface {
	return &authenticationCore{
		db: db,
	}
}
