package config

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"log/slog"
	"os"
	"time"

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
		slog.ErrorContext(context.Background(), "Failed to get current working directory", slog.String(constants.Error, err.Error()))
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
		slog.ErrorContext(context.Background(), "Failed to read config file", slog.String(constants.Error, err.Error()))
		panic(err.Error())
	}

	configInstance = &Config{
		App: AppConfig{
			Env:                  coalesce(vpr.GetString("APP_ENV"), "local"), // TODO: 0 usage
			RestHost:             coalesce(vpr.GetString("APP_REST_HOST"), "localhost"),
			RestPort:             coalesce(vpr.GetInt("APP_REST_PORT"), 4001),
			RestPrefork:          coalesce(vpr.GetBool("FIBER_PREFORK"), false),
			RestRecovery:         coalesce(vpr.GetBool("FIBER_RECOVERY"), true),
			RestEnableStackTrace: coalesce(vpr.GetBool("FIBER_ENABLE_STACK_TRACE"), true),
			RestCorsAllowOrigins: coalesce(vpr.GetString("FIBER_CORS_ALLOW_ORIGINS"), "*"),
			RestCorsAllowHeaders: coalesce(vpr.GetString("FIBER_CORS_ALLOW_HEADERS"), "Origin, Content-Type, Accept, Authorization"),
			RestCorsAllowMethods: coalesce(vpr.GetString("FIBER_CORS_ALLOW_METHODS"), "GET, POST, DELETE, PUT, PATCH, OPTIONS"),
			GrpcHost:             coalesce(vpr.GetString("APP_GRPC_HOST"), "localhost"),
			GrpcPort:             coalesce(vpr.GetInt("APP_GRPC_PORT"), 50000),
			IdleTimeout:          coalesce(vpr.GetDuration("APP_IDLE_TIMEOUT"), 30*time.Second),
			ReadTimeout:          coalesce(vpr.GetDuration("APP_READ_TIMEOUT"), 30*time.Second),
			WriteTimeout:         coalesce(vpr.GetDuration("APP_WRITE_TIMEOUT"), 30*time.Second),
			Name:                 coalesce(vpr.GetString("APP_NAME"), "example"),
			Version:              coalesce(vpr.GetString("APP_VERSION"), "1.0.0"),
			Host:                 coalesce(vpr.GetString("APP_HOST"), "http://localhost:4001"),
			LogLevel:             coalesce(vpr.GetString("APP_LOG_LEVEL"), "info"),
			JwtTokenKey:          coalesce(vpr.GetString("APP_JWT_TOKEN_KEY"), "sometokenkey"),
			JwtTokenExpiration:   coalesce(vpr.GetString("APP_JWT_TOKEN_EXPIRATION"), "10h"),
			JwtTokenRefresh:      coalesce(vpr.GetString("APP_JWT_TOKEN_REFRESH"), "24h"),
		},
		Db: DbConfig{
			Protocol:                 valueobject.SupportedDbFromString(coalesce(vpr.GetString("DB_PROTOCOL"), "mongodb")),
			Address:                  coalesce(vpr.GetString("DB_ADDRESS"), "localhost:27017"),
			Name:                     coalesce(vpr.GetString("DB_NAME"), "example"),
			Username:                 coalesce(vpr.GetString("DB_USERNAME"), "admin"),
			Password:                 coalesce(vpr.GetString("DB_PASSWORD"), "admin_password"),
			Option:                   coalesce(vpr.GetString("DB_OPTION"), "?authSource=admin&readPreference=secondaryPreferred"),
			MaxOpenConnections:       coalesce(vpr.GetInt("DB_MAX_OPEN_CONNECTIONS"), 10),
			MaxIdleConnections:       coalesce(vpr.GetInt("DB_MAX_IDLE_CONNECTIONS"), 2),
			MaxConnectionLifetime:    coalesce(vpr.GetDuration("DB_MAX_CONNECTION_LIFETIME"), 60*time.Second),
			MaxConnectionIdletime:    coalesce(vpr.GetDuration("DB_MAX_CONNECTION_IDLE_TIME"), 20*time.Second),
			ExampleMessageCollection: coalesce(vpr.GetString("DB_EXAMPLE_MESSAGE_COLLECTION"), "messages"),
			UserCollection:           coalesce(vpr.GetString("DB_USER_COLLECTION"), "users"),
		},
		Swagger: SwaggerConfig{
			DeepLinking:  coalesce(vpr.GetBool("SWAGGER_DEEP_LINKING"), true),
			DocExpansion: coalesce(vpr.GetString("SWAGGER_DOC_EXPANSION"), "list"),
		},
		Minio: MinioConfig{
			Endpoint:        coalesce(vpr.GetString("MINIO_ENDPOINT"), "localhost:9000"),
			AccessKeyID:     coalesce(vpr.GetString("MINIO_ACCESS_KEY"), "miniouser"),
			SecretAccessKey: coalesce(vpr.GetString("MINIO_SECRET_KEY"), "miniopassword"),
			UseSSL:          coalesce(vpr.GetBool("MINIO_USE_SSL"), false),
			BucketName:      coalesce(vpr.GetString("MINIO_BUCKET_NAME"), "example"),
		},
		Redis: RedisConfig{
			Db:       coalesce(vpr.GetInt("REDIS_DB"), 0),
			Host:     coalesce(vpr.GetString("REDIS_HOST"), "localhost"),
			Port:     coalesce(vpr.GetInt("REDIS_PORT"), 6379),
			Username: coalesce(vpr.GetString("REDIS_USERNAME"), ""),
			Password: coalesce(vpr.GetString("REDIS_PASSWORD"), ""),
		},
	}

	return *configInstance
}

func coalesce[T comparable](first, second T) T {
	var zero T
	if first != zero {
		return first
	}
	return second
}
