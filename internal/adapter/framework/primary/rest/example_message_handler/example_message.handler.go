package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
)

type ExampleMessageHandler struct {
	app app.ExampleMessageAppInterface
	jwt *rest_util.Jwt
}

func NewExampleRestHandler(
	app app.ExampleMessageAppInterface,
	jwt *rest_util.Jwt,
) ExampleMessageHandler {
	return ExampleMessageHandler{
		app: app,
		jwt: jwt,
	}
}
