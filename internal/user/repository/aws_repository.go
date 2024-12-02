package repository

import (
	"context"
	"net/url"
	"time"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/minio/minio-go/v7"
)

type UserAWSRepositoryImpl struct {
	s3Client *minio.Client
}

func NewAWSUserRepository(s3Client *minio.Client) user.UserAWSRepository {
	return &UserAWSRepositoryImpl{s3Client: s3Client}
}

func (r UserAWSRepositoryImpl) PutObject(ctx context.Context, entity *model.UserUploadInput) (*minio.UploadInfo, error) {
	panic("not implemented") // TODO: Implement
}

func (r UserAWSRepositoryImpl) GetObject(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	panic("not implemented") // TODO: Implement
}

func (r UserAWSRepositoryImpl) RemoveObject(ctx context.Context, bucketName string, objectName string) error {
	panic("not implemented") // TODO: Implement
}

func (r UserAWSRepositoryImpl) PresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry time.Duration) (*url.URL, error) {
	panic("not implemented") // TODO: Implement
}
