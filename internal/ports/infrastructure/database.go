package infrastructure

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type DbInterface interface {
	Connect() (err error)
	Disconnect()
	GetClient() (client *entity.DbClient)
}
