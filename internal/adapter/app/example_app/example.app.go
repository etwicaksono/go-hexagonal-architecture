package example_app

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/go-playground/validator/v10"
)

type exampleApp struct {
	core      core.ExampleCoreInterface
	validator *validator.Validate
}

func NewExampleApp(
	core core.ExampleCoreInterface,
	validator *validator.Validate,
) app.ExampleAppInterface {
	return &exampleApp{
		core:      core,
		validator: validator,
	}
}
