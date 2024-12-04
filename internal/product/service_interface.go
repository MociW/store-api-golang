package product

import (
	"context"

	"github.com/MociW/store-api-golang/internal/product/model/dto"
)

type ProductService interface {
	CreateProduct(ctx context.Context, entity *dto.ProductCreateRequest) (*dto.ProductResponse, error)

	UpdateProduct(ctx context.Context, entity *dto.ProductUpdateRequest) (*dto.ProductResponse, error)

	DeleteProduct(ctx context.Context, entity *dto.ProductDeleteRequest) error

	FindProduct(ctx context.Context, entity *dto.ProductFindRequest) (*dto.ProductResponse, error)

	ListProduct(ctx context.Context, id string) ([]dto.ProductResponse, error)
}

type ProductSKUService interface {
}
