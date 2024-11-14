package example_message_core

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/minio"
)

type exampleMessageCore struct {
	db        db.ExampleMessageDbInterface
	minio     minio.MinioInterface
	appConfig config.AppConfig
}

func NewExampleMessageCore(
	db db.ExampleMessageDbInterface,
	minio minio.MinioInterface,
	config config.Config,
) core.ExampleMessageCoreInterface {
	return &exampleMessageCore{
		db:        db,
		minio:     minio,
		appConfig: config.App,
	}
}
