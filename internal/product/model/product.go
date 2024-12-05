package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Summary     string         `gorm:"column:summary"`
	Images      pq.StringArray `gorm:"column:images;type:text[]"`
	UserID      string         `gorm:"column:user_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ProductSKUs []ProductSKU   `gorm:"foreignKey:product_id;references:id"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductSKU struct {
	ID        uint           `gorm:"primarykey"`
	ProductID uint           `gorm:"column:product_id"`
	Size      string         `gorm:"column:size"`
	Color     string         `gorm:"column:color"`
	SKU       string         `gorm:"column:sku"`
	Price     float32        `gorm:"column:price"`
	Quantity  int            `gorm:"column:quantity"`
	UserID    string         `gorm:"column:user_id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Product   Product        `gorm:"foreignKey:product_id;references:id"`
}

func (p *ProductSKU) TableName() string {
	return "product_skus"
}
