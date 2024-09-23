package authentication_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
}

func NewAuthenticationRestHandler() rest.AuthenticationHandlerInterface {
	return &adapter{}
}
