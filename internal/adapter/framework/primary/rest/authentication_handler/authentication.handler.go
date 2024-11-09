package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
)

type AuthenticationHandler struct {
	app app.AuthenticationAppInterface
}

func NewAuthenticationRestHandler(
	app app.AuthenticationAppInterface,
) AuthenticationHandler {
	return AuthenticationHandler{
		app: app,
	}
}
