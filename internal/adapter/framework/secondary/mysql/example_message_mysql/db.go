package example_message_mysql

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"gorm.io/gorm"
)

type userMongo struct {
	client *gorm.DB
	table  string
}

func NewExampleMessageMysql(config config.Config, gormDb *gorm.DB) db.ExampleMessageDbInterface {
	return &userMongo{
		client: gormDb,
		table:  config.Db.ExampleMessageCollection,
	}
}
