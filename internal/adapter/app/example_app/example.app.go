package example_app

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
)

type exampleApp struct {
	core core.ExampleCoreInterface
}

func NewExampleApp(
	core core.ExampleCoreInterface,
) app.ExampleAppInterface {
	return &exampleApp{
		core: core,
	}
}
