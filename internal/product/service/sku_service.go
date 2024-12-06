package service

import (
	"context"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/MociW/store-api-golang/internal/product/model/dto"
	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/pkg/errors"
)

type ProductSKUServiceImpl struct {
	cfg         *config.Config
	productRepo product.ProductSKURepository
}

func NewProductSKUService(cfg *config.Config, productRepo product.ProductSKURepository) product.ProductSKUService {
	return &ProductSKUServiceImpl{cfg: cfg, productRepo: productRepo}
}

func (product *ProductSKUServiceImpl) CreateSKU(ctx context.Context, entity *dto.ProductSKUCreateRequest) (*dto.ProductSKUResponese, error) {
	sku := &model.ProductSKU{
		ProductID: entity.ProductID,
		Size:      entity.Size,
		Color:     entity.Color,
		SKU:       util.SKUGenerator(entity.UserID, entity.Name, entity.Size, entity.Color),
		Price:     entity.Price,
		Quantity:  entity.Quantity,
		UserID:    entity.UserID,
	}

	result, err := product.productRepo.CreateSKU(ctx, sku)
	if err != nil {
		return nil, errors.Wrap(err, "ProductSKUService.CreateSKU.CreateSKU")
	}

	return dto.ConvertSKUResponse(result), nil
}

func (product *ProductSKUServiceImpl) DeleteSKU(ctx context.Context, entity *dto.ProductSKUDeleteRequest) error {
	sku := &model.ProductSKU{
		ID:        entity.ID,
		ProductID: entity.ProductID,
		UserID:    entity.UserID,
	}

	err := product.productRepo.DeleteSKU(ctx, sku)
	if err != nil {
		return errors.Wrap(err, "ProductSKUService.DeleteSKU.DeleteSKU")
	}

	return nil
}

func (product *ProductSKUServiceImpl) FindSKU(ctx context.Context, entity *dto.ProductSKUFindRequest) (*dto.ProductSKUResponese, error) {
	sku := &model.ProductSKU{
		ID:        entity.ID,
		ProductID: entity.ProductID,
		UserID:    entity.UserID,
	}

	result, err := product.productRepo.FindSKU(ctx, sku)

	if err != nil {
		return nil, errors.Wrap(err, "ProductSKUService.DeleteSKU.DeleteSKU")
	}

	return dto.ConvertSKUResponse(result), nil
}

func (product *ProductSKUServiceImpl) ListSKU(ctx context.Context, productID uint, userID string) ([]dto.ProductSKUResponese, error) {
	sku := &model.ProductSKU{
		ProductID: productID,
		UserID:    userID,
	}

	result, err := product.productRepo.ListSKU(ctx, sku)
	if err != nil {
		return nil, errors.Wrap(err, "ProductSKUService.ListSKU.ListSKU")
	}

	responses := make([]dto.ProductSKUResponese, len(result))
	for i, data := range result {
		responses[i] = *dto.ConvertSKUResponse(&data)
	}

	return responses, nil
}
