package minio

import (
	"bytes"
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	minio2 "github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/minio"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log/slog"
)

type adapterMinio struct {
	ctx        context.Context
	client     *minio.Client
	bucketName string
}

var minioInstance *adapterMinio

func MinioProvider(ctx context.Context, cfg config.Config) minio2.MinioInterface {
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

func (a adapterMinio) Upload(ctx context.Context, data []byte, contentType string, filePath string) (result minio.UploadInfo, err error) {
	fileSize := int64(len(data))
	reader := bytes.NewReader(data)
	return a.client.PutObject(ctx, a.bucketName, filePath, reader, fileSize, minio.PutObjectOptions{ContentType: contentType})
}

func (a adapterMinio) Remove(ctx context.Context, filePath string) (err error) {
	return a.client.RemoveObject(ctx, a.bucketName, filePath, minio.RemoveObjectOptions{})
}

func (a adapterMinio) GetClient() *minio.Client {
	return a.client
}