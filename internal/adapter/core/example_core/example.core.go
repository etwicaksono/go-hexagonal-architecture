package example_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type exampleCore struct {
	ctx context.Context
	db  db.ExampleDbInterface
}

func NewExampleCore(
	ctx context.Context,
	db db.ExampleDbInterface,
) core.ExampleCoreInterface {
	return &exampleCore{
		ctx: ctx,
		db:  db,
	}
}
