package infrastructure

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log/slog"
	"mime/multipart"
)

type adapterMinio struct {
	ctx        context.Context
	client     *minio.Client
	bucketName string
}

var minioInstance *adapterMinio

func MinioProvider(ctx context.Context, cfg config.Config) infrastructure.MinioInterface {
	if minioInstance != nil {
		return minioInstance
	}

	endpoint := cfg.Minio.Endpoint
	accessKeyID := cfg.Minio.AccessKeyID
	secretAccessKey := cfg.Minio.SecretAccessKey
	useSSL := cfg.Minio.UseSSL
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		slog.Info("Minio client initialization failed", slog.String(entity.Error, err.Error()))
	}
	slog.Info("Minio client initialized successfully")

	minioInstance = &adapterMinio{
		ctx:        ctx,
		client:     client,
		bucketName: cfg.Minio.BucketName,
	}
	return minioInstance
}

func (a adapterMinio) Upload(ctx context.Context, file *multipart.FileHeader, filePath string) (result minio.UploadInfo, err error) {
	// get buffer
	buffer, err := file.Open()
	if err != nil {
		return minio.UploadInfo{}, err
	}
	defer func() {
		err := buffer.Close()
		if err != nil {
			slog.ErrorContext(ctx, "Failed to close buffer", slog.String(entity.Error, err.Error()))
		}
	}()

	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	return a.client.PutObject(ctx, a.bucketName, filePath, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
}

func (a adapterMinio) Remove(ctx context.Context, filePath string) (err error) {
	return a.client.RemoveObject(ctx, a.bucketName, filePath, minio.RemoveObjectOptions{})
}

func (a adapterMinio) GetClient() *minio.Client {
	return a.client
}
