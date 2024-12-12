package aws

import (
	"github.com/MociW/store-api-golang/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

func NewAWSClient(config *config.Config) (*minio.Client, error) {
	client, err := minio.New(config.AWS.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AWS.MinioAccessKey, config.AWS.MinioSecretKey, ""),
		Secure: config.AWS.UseSSL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "NewAWSClient.minio.New")
	}

	return client, err
}
