package example_message_mysql

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
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

func (u userMongo) FindAllTextMessage(ctx context.Context) (result []entity.MessageTextItem, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userMongo) FindAllMultimediaMessage(ctx context.Context) (result []entity.MessageMultimediaItem, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userMongo) InsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (result entity.BulkWriteResult, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userMongo) InsertMultimediaMessage(ctx context.Context, objs []entity.MessageMultimediaItem) (result entity.BulkWriteResult, err error) {
	//TODO implement me
	panic("implement me")
}
