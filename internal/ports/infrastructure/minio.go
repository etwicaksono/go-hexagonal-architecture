package infrastructure

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type MinioInterface interface {
	Upload(ctx context.Context, file *multipart.FileHeader, filePath string) (result minio.UploadInfo, err error)
	Remove(ctx context.Context, filePath string) (err error)
	GetClient() *minio.Client
}
