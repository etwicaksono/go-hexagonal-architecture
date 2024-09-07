package example_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type exampleCore struct {
	db db.ExampleDbInterface
}

func NewExampleCore(
	db db.ExampleDbInterface,
) core.ExampleCoreInterface {
	return &exampleCore{
		db: db,
	}
}
