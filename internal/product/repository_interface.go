package product

import (
	"context"
	"net/url"
	"time"

	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/minio/minio-go/v7"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, entity *model.Product) (*model.Product, error)

	UpdateProduct(ctx context.Context, entity *model.Product) (*model.Product, error)

	DeleteProduct(ctx context.Context, entity *model.Product) error

	FindProduct(ctx context.Context, entity *model.Product) (*model.Product, error)

	ListProduct(ctx context.Context, entity string) ([]model.Product, error)
}

type ProductSKURepository interface {
	CreateSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error)

	UpdateSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error)

	DeleteSKU(ctx context.Context, entity *model.ProductSKU) error

	FindSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error)

	ListSKU(ctx context.Context, entity *model.ProductSKU) ([]model.ProductSKU, error)
}

type ProductAWSRepository interface {
	PutObject(ctx context.Context, entity *model.ProductUploadInput) (*minio.UploadInfo, error)

	GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error)

	RemoveObject(ctx context.Context, bucketName, objectName string) error

	PresignedGetObject(ctx context.Context, bucketName, objectName string, expiry time.Duration) (*url.URL, error)
}
