package example_message_app

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/go-playground/validator/v10"
)

type exampleMessageApp struct {
	core      core.ExampleMessageCoreInterface
	validator *validator.Validate
}

func NewExampleMessageApp(
	core core.ExampleMessageCoreInterface,
	validator *validator.Validate,
) app.ExampleMessageAppInterface {
	return &exampleMessageApp{
		core:      core,
		validator: validator,
	}
}
