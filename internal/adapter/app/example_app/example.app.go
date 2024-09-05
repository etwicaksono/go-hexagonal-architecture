package example_app

import (
	"context"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
)

type exampleApp struct {
	ctx  context.Context
	core core.ExampleCoreInterface
}

func NewExampleApp(
	ctx context.Context,
	core core.ExampleCoreInterface,
) app.ExampleAppInterface {
	return &exampleApp{
		core: core,
	}
}
