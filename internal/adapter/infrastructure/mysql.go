package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type adapterMysql struct {
	ctx           context.Context
	config        config.Config
	connectionURL string
	client        *gorm.DB
	logger        *slog.Logger
}

func NewMysqlDb(
	ctx context.Context,
	config config.Config,
	logger *slog.Logger,
) infrastructure.DbInterface {
	return &adapterMysql{
		ctx:    ctx,
		config: config,
		connectionURL: fmt.Sprintf(
			"%s:%s@tcp(%s)/%s%s",
			config.Db.Username,
			config.Db.Password,
			config.Db.Address,
			config.Db.Name,
			config.Db.Option,
		),
		logger: logger,
	}
}

func (a *adapterMysql) Connect() error {
	// Set GORM logger to use SlogGormLogger with the desired log level
	var logLevel logger.LogLevel

	switch a.config.App.LogLevel {
	case "debug":
		logLevel = logger.Info
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	default:
		return errorsConst.ErrInvalidLogLevel
	}

	gormLogger := &slogGormLogger{logger: a.logger, level: logLevel}

	db, err := gorm.Open(mysql.Open(a.connectionURL), &gorm.Config{
		Logger:         gormLogger,
		TranslateError: true,
	})
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to connect to MySQL", slog.String("connection", a.connectionURL), slog.String(entity.Error, err.Error()))
		return err
	}

	slog.InfoContext(a.ctx, "MySQL connected", slog.String(
		"connected to", a.connectionURL,
	))
	a.client = db

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get *sql.DB: %v", slog.String(entity.Error, err.Error()))
		return err
	}

	// Set advanced connection settings
	sqlDB.SetMaxOpenConns(a.config.Db.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(a.config.Db.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(a.config.Db.MaxConnectionLifetime * time.Second)
	sqlDB.SetConnMaxIdleTime(a.config.Db.MaxConnectionIdletime * time.Second)

	return nil
}

func (a *adapterMysql) Disconnect() {
	sqlDB, err := a.client.DB()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get database object: %v", err))
	}

	if err = sqlDB.Close(); err != nil {
		slog.Error(fmt.Sprintf("Failed to close database connection: %v", err))
	}
	slog.Info("Database connection closed")
}

func (a *adapterMysql) GetClient() *entity.DbClient {
	return &entity.DbClient{
		GormClient: a.client,
	}
}

type slogGormLogger struct {
	logger *slog.Logger
	level  logger.LogLevel
}

func (l *slogGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &slogGormLogger{logger: l.logger, level: level}
}

func (l *slogGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		l.logger.InfoContext(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *slogGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.logger.WarnContext(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *slogGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		l.logger.ErrorContext(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *slogGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	logData := []interface{}{"sql", sql, "rows", rows, "elapsed", elapsed}
	if err != nil && l.level >= logger.Error {
		l.logger.ErrorContext(ctx, "Error executing SQL", slog.Any(entity.Error, logData))
	} else if l.level >= logger.Info {
		l.logger.InfoContext(ctx, "SQL executed", logData...)
	}
}
