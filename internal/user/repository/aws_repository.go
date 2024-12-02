package repository

import (
	"context"
	"net/url"
	"time"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type UserAWSRepositoryImpl struct {
	s3Client *minio.Client
}

func NewAWSUserRepository(s3Client *minio.Client) user.UserAWSRepository {
	return &UserAWSRepositoryImpl{s3Client: s3Client}
}

func (aws UserAWSRepositoryImpl) PutObject(ctx context.Context, entity *model.UserUploadInput) (*minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
		ContentType:  entity.ContentType,
	}

	uploadInfo, err := aws.s3Client.PutObject(ctx, entity.BucketName, entity.ObjectName, entity.Object, entity.ObjectSize, opts)
	if err != nil {
		return nil, errors.Wrap(err, "AWSUserRepository.PutObject.s3Client.PutObject")
	}
	return &uploadInfo, nil
}

func (aws UserAWSRepositoryImpl) GetObject(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	object, err := aws.s3Client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "AWSUserRepository.PutObject.s3Client.GetObject")
	}
	defer object.Close()

	return object, err
}

func (aws UserAWSRepositoryImpl) RemoveObject(ctx context.Context, bucketName string, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := aws.s3Client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return errors.Wrap(err, "AWSUserRepository.PutObject.s3Client.RemoveObject")
	}

	return nil
}

func (aws UserAWSRepositoryImpl) PresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry time.Duration) (*url.URL, error) {
	var reqParam = make(url.Values)

	presignedUrl, err := aws.s3Client.PresignedGetObject(ctx, bucketName, objectName, expiry, reqParam)
	if err != nil {
		return nil, errors.Wrap(err, "AWSUserRepository.PresignedGetObject.s3Client.PresignedGetObject")
	}

	return presignedUrl, nil
}
