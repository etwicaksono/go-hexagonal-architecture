package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/valueobject"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/example_message_mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/user_mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mysql/example_message_mysql"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mysql/user_mysql"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"github.com/sagikazarmark/slog-shim"
)

func userDbProvider(cfg config.Config, dbClient *entity.DbClient) db.UserDbInterface {
	switch cfg.Db.Protocol {
	case valueobject.SuportedDb_MONGO:
		{
			return user_mongo.NewUserMongo(cfg, dbClient.MongoClient)
		}
	case valueobject.SuportedDb_MYSQL:
		{
			return user_mysql.NewUserMysql(cfg, dbClient.GormClient)
		}
	default:
		{
			slog.Error(errorsConst.ErrUnsupportedDbProtocol.Error(), slog.String("protocol", cfg.Db.Protocol.ToString()))
			return nil
		}
	}
}

func messageDbProvider(cfg config.Config, dbClient *entity.DbClient) db.ExampleMessageDbInterface {
	switch cfg.Db.Protocol {
	case valueobject.SuportedDb_MONGO:
		{
			return example_message_mongo.NewExampleMessageMongo(cfg, dbClient.MongoClient)
		}
	case valueobject.SuportedDb_MYSQL:
		{
			return example_message_mysql.NewExampleMessageMysql(cfg, dbClient.GormClient)
		}
	default:
		{
			slog.Error(errorsConst.ErrUnsupportedDbProtocol.Error(), slog.String("protocol", cfg.Db.Protocol.ToString()))
			return nil
		}
	}
}
