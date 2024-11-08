package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
)

type MinioInterface interface {
	Upload(ctx context.Context, data []byte, contentType string, filePath string) (result minio.UploadInfo, err error)
	Remove(ctx context.Context, filePath string) (err error)
	GetClient() (minioClient *minio.Client)
	GetBucketName() (bucketName string)
	GetEndpoint() (minioEndpoint string)
	IsUseSSL() (isUseSSL bool)
}
