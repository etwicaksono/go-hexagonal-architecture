package example_message_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/minio"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
)

type exampleMessageCore struct {
	db    db.ExampleMessageDbInterface
	minio minio.MinioInterface
}

func NewExampleMessageCore(
	db db.ExampleMessageDbInterface,
	minio minio.MinioInterface,
) core.ExampleMessageCoreInterface {
	return &exampleMessageCore{
		db:    db,
		minio: minio,
	}
}
