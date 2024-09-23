package example_message_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
	app app.ExampleMessageAppInterface
}

func NewExampleRestHandler(
	app app.ExampleMessageAppInterface,
) rest.ExampleMessageHandlerInterface {
	return &adapter{
		app: app,
	}
}