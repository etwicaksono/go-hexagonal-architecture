package example_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type exampleCore struct {
	db    db.ExampleDbInterface
	minio infrastructure.MinioInterface
}

func NewExampleCore(
	db db.ExampleDbInterface,
	minio infrastructure.MinioInterface,
) core.ExampleCoreInterface {
	return &exampleCore{
		db:    db,
		minio: minio,
	}
}
