package repository

import (
	"context"
	"net/url"
	"time"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type ProductAWSRepositoryImpl struct {
	s3Client *minio.Client
}

func NewProductAWSRepository(s3Client *minio.Client) product.ProductAWSRepository {
	return &ProductAWSRepositoryImpl{s3Client: s3Client}
}

func (aws ProductAWSRepositoryImpl) PutObject(ctx context.Context, entity *model.ProductUploadInput) (*minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
		ContentType:  entity.ContentType,
	}

	uploadInfo, err := aws.s3Client.PutObject(ctx, entity.BucketName, entity.ObjectName, entity.Object, entity.ObjectSize, opts)
	if err != nil {
		return nil, errors.Wrap(err, "ProductAWSRepository.PutObject.s3Client.PutObject")
	}
	return &uploadInfo, nil
}

func (aws ProductAWSRepositoryImpl) GetObject(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	object, err := aws.s3Client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "ProductAWSRepository.PutObject.s3Client.GetObject")
	}
	defer object.Close()

	return object, err
}

func (aws ProductAWSRepositoryImpl) RemoveObject(ctx context.Context, bucketName string, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := aws.s3Client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return errors.Wrap(err, "ProductAWSRepository.PutObject.s3Client.RemoveObject")
	}

	return nil
}

func (aws ProductAWSRepositoryImpl) PresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry time.Duration) (*url.URL, error) {
	var reqParam = make(url.Values)

	presignedUrl, err := aws.s3Client.PresignedGetObject(ctx, bucketName, objectName, expiry, reqParam)
	if err != nil {
		return nil, errors.Wrap(err, "ProductAWSRepository.PresignedGetObject.s3Client.PresignedGetObject")
	}

	return presignedUrl, nil
}
