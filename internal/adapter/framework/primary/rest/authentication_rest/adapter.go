package authentication_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
	app app.AuthenticationAppInterface
}

func NewAuthenticationRestHandler(
	app app.AuthenticationAppInterface,
) rest.AuthenticationHandlerInterface {
	return &adapter{
		app: app,
	}
}
