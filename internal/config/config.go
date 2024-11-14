package config

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"log/slog"
	"os"
	"time"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/spf13/viper"
)

type Config struct {
	App     AppConfig
	Db      DbConfig
	Swagger SwaggerConfig
	Minio   MinioConfig
	Redis   RedisConfig
}

type AppConfig struct {
	Env                  string
	RestHost             string
	RestPort             int
	RestPrefork          bool
	RestRecovery         bool
	RestEnableStackTrace bool
	RestCorsAllowOrigins string
	RestCorsAllowHeaders string
	RestCorsAllowMethods string
	GrpcHost             string
	GrpcPort             int
	IdleTimeout          time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	GracefulTimeout      time.Duration
	Name                 string
	Version              string
	Host                 string
	LogLevel             string
	JwtTokenKey          string
	JwtTokenExpiration   string
	JwtTokenRefresh      string
}

type DbConfig struct {
	Protocol                 valueobject.SupportedDb
	Address                  string
	Name                     string
	Username                 string
	Password                 string
	Option                   string
	MaxOpenConnections       int
	MaxIdleConnections       int
	MaxConnectionLifetime    time.Duration
	MaxConnectionIdletime    time.Duration
	ExampleMessageCollection string
	UserCollection           string
}

type SwaggerConfig struct {
	DeepLinking  bool
	DocExpansion string
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

type RedisConfig struct {
	Db       int
	Host     string
	Port     int
	Username string
	Password string
}

var configInstance *Config

func LoadConfig() Config {
	if configInstance != nil {
		return *configInstance
	}
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to get current working directory", slog.String(entity.Error, err.Error()))
		panic(err.Error())
	}

	// Root folder of this project
	projectRoot := wd
	vpr := viper.New()

	vpr.AddConfigPath(projectRoot)
	vpr.SetConfigFile(fmt.Sprintf("%s/.env", projectRoot))
	vpr.AutomaticEnv()

	err = vpr.ReadInConfig()
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to read config file", slog.String(entity.Error, err.Error()))
		panic(err.Error())
	}

	configInstance = &Config{
		App: AppConfig{
			Env:                  vpr.GetString("APP_ENV"),
			RestHost:             vpr.GetString("APP_REST_HOST"),
			RestPort:             vpr.GetInt("APP_REST_PORT"),
			RestPrefork:          vpr.GetBool("FIBER_PREFORK"),
			RestRecovery:         vpr.GetBool("FIBER_RECOVERY"),
			RestEnableStackTrace: vpr.GetBool("FIBER_ENABLE_STACK_TRACE"),
			RestCorsAllowOrigins: vpr.GetString("FIBER_CORS_ALLOW_ORIGINS"),
			RestCorsAllowHeaders: vpr.GetString("FIBER_CORS_ALLOW_HEADERS"),
			RestCorsAllowMethods: vpr.GetString("FIBER_CORS_ALLOW_METHODS"),
			GrpcHost:             vpr.GetString("APP_GRPC_HOST"),
			GrpcPort:             vpr.GetInt("APP_GRPC_PORT"),
			IdleTimeout:          vpr.GetDuration("APP_IDLE_TIMEOUT"),
			ReadTimeout:          vpr.GetDuration("APP_READ_TIMEOUT"),
			WriteTimeout:         vpr.GetDuration("APP_WRITE_TIMEOUT"),
			GracefulTimeout:      vpr.GetDuration("APP_GRACEFUL_TIMEOUT"),
			Name:                 vpr.GetString("APP_NAME"),
			Version:              vpr.GetString("APP_VERSION"),
			Host:                 vpr.GetString("APP_HOST"),
			LogLevel:             vpr.GetString("APP_LOG_LEVEL"),
			JwtTokenKey:          vpr.GetString("APP_JWT_TOKEN_KEY"),
			JwtTokenExpiration:   vpr.GetString("APP_JWT_TOKEN_EXPIRATION"),
			JwtTokenRefresh:      vpr.GetString("APP_JWT_TOKEN_REFRESH"),
		},
		Db: DbConfig{
			Protocol:                 valueobject.SupportedDbFromString(vpr.GetString("DB_PROTOCOL")),
			Address:                  vpr.GetString("DB_ADDRESS"),
			Name:                     vpr.GetString("DB_NAME"),
			Username:                 vpr.GetString("DB_USERNAME"),
			Password:                 vpr.GetString("DB_PASSWORD"),
			Option:                   vpr.GetString("DB_OPTION"),
			MaxOpenConnections:       vpr.GetInt("DB_MAX_OPEN_CONNECTIONS"),
			MaxIdleConnections:       vpr.GetInt("DB_MAX_IDLE_CONNECTIONS"),
			MaxConnectionLifetime:    vpr.GetDuration("DB_MAX_CONNECTION_LIFETIME"),
			MaxConnectionIdletime:    vpr.GetDuration("DB_MAX_CONNECTION_IDLE_TIME"),
			ExampleMessageCollection: vpr.GetString("DB_EXAMPLE_MESSAGE_COLLECTION"),
			UserCollection:           vpr.GetString("DB_USER_COLLECTION"),
		},
		Swagger: SwaggerConfig{
			DeepLinking:  vpr.GetBool("SWAGGER_DEEP_LINKING"),
			DocExpansion: vpr.GetString("SWAGGER_DOC_EXPANSION"),
		},
		Minio: MinioConfig{
			Endpoint:        vpr.GetString("MINIO_ENDPOINT"),
			AccessKeyID:     vpr.GetString("MINIO_ACCESS_KEY_ID"),
			SecretAccessKey: vpr.GetString("MINIO_SECRET_ACCESS_KEY"),
			UseSSL:          vpr.GetBool("MINIO_USE_SSL"),
			BucketName:      vpr.GetString("MINIO_BUCKET_NAME"),
		},
		Redis: RedisConfig{
			Db:       vpr.GetInt("REDIS_DB"),
			Host:     vpr.GetString("REDIS_HOST"),
			Port:     vpr.GetInt("REDIS_PORT"),
			Username: vpr.GetString("REDIS_USERNAME"),
			Password: vpr.GetString("REDIS_PASSWORD"),
		},
	}

	return *configInstance
}
