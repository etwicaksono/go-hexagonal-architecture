package middleware

import "github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"

type Middleware struct {
	Jwt *rest_util.Jwt
}

func NewMiddleware(jwt *rest_util.Jwt) *Middleware {
	return &Middleware{
		Jwt: jwt,
	}
}
