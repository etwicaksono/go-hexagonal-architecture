package user_mysql

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"gorm.io/gorm"
)

type userMysql struct {
	client *gorm.DB
	table  string
}

func NewUserMysql(config config.Config, gormDb *gorm.DB) db.UserDbInterface {
	return &userMysql{
		client: gormDb,
		table:  config.Db.UserCollection,
	}
}
