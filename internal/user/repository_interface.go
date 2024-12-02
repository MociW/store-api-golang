package user

import (
	"context"
	"net/url"
	"time"

	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/minio/minio-go/v7"
)

type UserPostgresRepository interface {
	/* ---------------------------------- User ---------------------------------- */
	CreateUser(ctx context.Context, entity *model.User) error

	UpdateUser(ctx context.Context, entity *model.User) error

	DeleteUser(ctx context.Context, entity *model.User) error

	FindByEmail(ctx context.Context, entity *model.User, email string) error

	FindByUsername(ctx context.Context, entity *model.User, username string) error

	/* --------------------------------- Address -------------------------------- */
	CreateAddress(ctx context.Context, entity *model.Address) error

	UpdateAddress(ctx context.Context, entity *model.Address) error

	DeleteAddress(ctx context.Context, entity *model.Address) error

	FindAddress(ctx context.Context, entity *model.Address, uuid string, id uint) error

	ListAddress(ctx context.Context, uuid string) ([]model.Address, error)
}

type UserAWSRepository interface {
	PutObject(ctx context.Context, entity *model.UserUploadInput) (*minio.UploadInfo, error)

	GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error)

	RemoveObject(ctx context.Context, bucketName, objectName string) error

	PresignedGetObject(ctx context.Context, bucketName, objectName string, expiry time.Duration) (*url.URL, error)
}
