package aws

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewAWSClient(config *viper.Viper) (*minio.Client, error) {
	client, err := minio.New(config.GetString("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(config.GetString("MINIO_ACCESS_KEY"), config.GetString("MINIO_SECRET_KEY"), ""),
		Secure: config.GetBool("MINIO_SECURE"),
	})
	if err != nil {
		return nil, errors.Wrap(err, "NewAWSClient.minio.New")
	}

	return client, err
}
