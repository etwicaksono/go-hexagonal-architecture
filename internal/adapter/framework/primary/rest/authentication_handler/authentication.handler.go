package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
)

type AuthenticationHandler struct {
	app app.AuthenticationAppInterface
	jwt *rest_util.Jwt
}

func NewAuthenticationRestHandler(
	app app.AuthenticationAppInterface,
	jwt *rest_util.Jwt,
) AuthenticationHandler {
	return AuthenticationHandler{
		app: app,
		jwt: jwt,
	}
}
