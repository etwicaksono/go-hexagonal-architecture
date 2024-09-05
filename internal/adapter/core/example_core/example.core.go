package example_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
)

type exampleCore struct {
}

func NewExampleCore() core.ExampleCoreInterface {
	return &exampleCore{}
}
