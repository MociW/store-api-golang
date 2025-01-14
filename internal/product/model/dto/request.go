package dto

import "github.com/MociW/store-api-golang/internal/product/model"

/* --------------------------------- Product -------------------------------- */

type ProductCreateRequest struct {
	Name        string                     `json:"name" validate:"required,max=100,alpha"`
	Description string                     `json:"description" validate:"required,max=100,alpha"`
	Summary     string                     `json:"summary" validate:"required,max=100,alpha"`
	Images      []model.ProductUploadInput `json:"-"`
	UserID      string                     `json:"user_id"`
}

type ProductUpdateRequest struct {
	ID          uint                       `json:"id"`
	Name        string                     `json:"name" validate:"required,max=100,alpha"`
	Description string                     `json:"description" validate:"required,max=100,alpha"`
	Summary     string                     `json:"summary" validate:"required,max=100,alpha"`
	Images      []model.ProductUploadInput `json:"-"`
	UserID      string                     `json:"user_id"`
}

type ProductDeleteRequest struct {
	ID     uint   `json:"id"`
	UserID string `json:"user_id"`
}

type ProductFindRequest struct {
	ID     uint   `json:"id"`
	UserID string `json:"user_id"`
}

/* ------------------------------- Product SKU ------------------------------ */

type ProductSKUCreateRequest struct {
	ProductID uint    `json:"-"`
	Name      string  `json:"name"`
	Size      string  `json:"size"`
	Color     string  `json:"color"`
	SKU       string  `json:"sku"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	UserID    string  `json:"-"`
}

type ProductSKUUpdateRequest struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Size      string  `json:"size"`
	Color     string  `json:"color"`
	SKU       string  `json:"sku"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	UserID    string  `json:"-"`
}

type ProductSKUDeleteRequest struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	UserID    string `json:"-"`
}

type ProductSKUFindRequest struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	UserID    string `json:"-"`
}

type ProductSKUListRequest struct {
	ProductID uint   `json:"product_id"`
	UserID    string `json:"-"`
}
