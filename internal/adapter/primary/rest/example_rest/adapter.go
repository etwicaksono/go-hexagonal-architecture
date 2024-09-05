package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct{}

type Config struct{}

func NewExampleRestHandler(config Config) rest.ExampleHandlerInterface {
	return &adapter{}
}
