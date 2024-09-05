package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
	app app.ExampleAppInterface
}

func NewExampleRestHandler(app app.ExampleAppInterface) rest.ExampleHandlerInterface {
	return &adapter{
		app: app,
	}
}
