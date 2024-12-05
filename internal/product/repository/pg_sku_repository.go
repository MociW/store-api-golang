package repository

import (
	"context"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductSKURepositoryImpl struct {
	DB *gorm.DB
}

func ProductSKURepository(db *gorm.DB) product.ProductSKURepository {
	return &ProductSKURepositoryImpl{DB: db}
}

func (r *ProductSKURepositoryImpl) CreateSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error) {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		return tx.Where("product_id = ? AND user_id = ?", entity.ProductID, entity.UserID).Create(entity).Error
	})

	if err != nil {
		return nil, errors.Wrap(err, "ProductRepository.Create.CreateSKU")
	}

	return entity, nil
}

func (r *ProductSKURepositoryImpl) UpdateSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error) {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.ProductSKU{}).Where("id = ? AND product_id = ? AND user_id = ?", entity.ID, entity.ProductID, entity.UserID).Updates(entity).Error
	})

	if err != nil {
		return nil, errors.Wrap(err, "ProductRepository.Update.UpdateSKU")
	}

	return entity, nil
}

func (r *ProductSKURepositoryImpl) DeleteSKU(ctx context.Context, entity *model.ProductSKU) error {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		sku := new(model.ProductSKU)
		if err := tx.Where("id = ? AND product_id = ? AND user_id = ?", entity.ID, entity.ProductID, entity.UserID).First(sku).Error; err != nil {
			return errors.Wrap(err, "ProductRepository.Delete.DeleteProduct")
		}

		if err := tx.Delete(sku).Error; err != nil {
			return errors.Wrap(err, "ProductRepository.Delete.DeleteProduct")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "ProductRepository.Delete.DeleteSKU")
	}

	return nil
}

func (r *ProductSKURepositoryImpl) FindSKU(ctx context.Context, entity *model.ProductSKU) (*model.ProductSKU, error) {
	sku := new(model.ProductSKU)
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("id = ? AND product_id = ? AND user_id = ?", entity.ID, entity.ProductID, entity.UserID).First(sku).Error; err != nil {
		return nil, errors.Wrap(err, "ProductRepository.Find.FindProduct")
	}

	return sku, nil
}

func (r *ProductSKURepositoryImpl) ListSKU(ctx context.Context, entity *model.ProductSKU) ([]model.ProductSKU, error) {
	var skus []model.ProductSKU
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("user_id = ?", entity).Find(&skus).Error; err != nil {
		return nil, errors.Wrap(err, "ProductRepository.List.ListProduct")
	}

	return skus, nil
}
