package authentication_app

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/go-playground/validator/v10"
)

type authenticationApp struct {
	core      core.AuthenticationCoreInterface
	validator *validator.Validate
	jwt       *rest_util.Jwt
}

func NewAuthenticationApp(
	core core.AuthenticationCoreInterface,
	validator *validator.Validate,
	jwt *rest_util.Jwt,
) app.AuthenticationAppInterface {
	return &authenticationApp{
		core:      core,
		validator: validator,
		jwt:       jwt,
	}
}
