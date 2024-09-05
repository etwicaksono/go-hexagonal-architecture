package example_core

import (
	"context"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
)

type exampleCore struct {
	ctx context.Context
}

func NewExampleCore(ctx context.Context) core.ExampleCoreInterface {
	return &exampleCore{
		ctx: ctx,
	}
}
