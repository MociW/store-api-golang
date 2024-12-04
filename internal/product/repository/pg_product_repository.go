package repository

import (
	"context"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (r ProductRepositoryImpl) CreateProduct(ctx context.Context, entity *model.Product) (*model.Product, error) {
	// Ensure entity is not nil
	if entity == nil {
		return nil, errors.New("ProductRepository.CreateProdduct: entity cannot be nil")
	}

	tx := r.DB.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		return tx.Save(entity).Error
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r ProductRepositoryImpl) UpdateProduct(ctx context.Context, entity *model.Product) (*model.Product, error) {
	// Ensure entity is not nil
	if entity == nil {
		return nil, errors.New("ProductRepository.UpdateProduct: entity cannot be nil")
	}

	tx := r.DB.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.Product{}).Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Updates(entity).Error
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r ProductRepositoryImpl) DeleteProduct(ctx context.Context, entity *model.Product) error {
	tx := r.DB.WithContext(ctx)
	return tx.Transaction(func(tx *gorm.DB) error {
		product := new(model.Product)
		if err := tx.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).First(product).Error; err != nil {
			return errors.Wrap(err, "ProductRepository.Delete.DeleteProduct")
		}

		if err := tx.Delete(product).Error; err != nil {
			return errors.Wrap(err, "ProductRepository.Delete.DeleteProduct")
		}

		return nil
	})
}

func (r ProductRepositoryImpl) FindProduct(ctx context.Context, entity *model.Product) (*model.Product, error) {
	product := new(model.Product)
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).First(product).Error; err != nil {
		return nil, errors.Wrap(err, "ProductRepository.Find.FindProduct")
	}

	return product, nil
}

func (r ProductRepositoryImpl) ListProduct(ctx context.Context, entity string) ([]model.Product, error) {
	var products []model.Product
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("user_id = ?", entity).Find(&products).Error; err != nil {
		return nil, errors.Wrap(err, "ProductRepository.List.ListProduct")
	}

	return products, nil
}
