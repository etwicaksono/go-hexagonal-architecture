package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
)

type ExampleMessageHandler struct {
	app app.ExampleMessageAppInterface
}

func NewExampleRestHandler(
	app app.ExampleMessageAppInterface,
) ExampleMessageHandler {
	return ExampleMessageHandler{
		app: app,
	}
}
