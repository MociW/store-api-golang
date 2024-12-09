package dto

import (
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/lib/pq"
)

type ApiProductResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ProductResponse struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Summary     string                `json:"summary"`
	Images      pq.StringArray        `json:"images"`
	UserID      string                `json:"user_id"`
	ProductSKU  []ProductSKUResponese `json:"product_sku,omitempty"`
}

type ProductSKUResponese struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Size      string  `json:"size"`
	Color     string  `json:"color"`
	SKU       string  `json:"sku"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	UserID    string  `json:"user_id"`
}

func ConvertProductResponse(entity *model.Product) *ProductResponse {

	responses := make([]ProductSKUResponese, len(entity.ProductSKUs))
	for i, sku := range entity.ProductSKUs {
		responses[i] = *ConvertSKUResponse(&sku)
	}

	if len(responses) == 0 {
		responses = nil
	}

	return &ProductResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Summary:     entity.Summary,
		Images:      entity.Images,
		UserID:      entity.UserID,
		ProductSKU:  responses,
	}
}

func ConvertSKUResponse(entity *model.ProductSKU) *ProductSKUResponese {
	return &ProductSKUResponese{
		ID:        entity.ID,
		ProductID: entity.ID,
		Size:      entity.Size,
		Color:     entity.Color,
		SKU:       entity.SKU,
		Price:     entity.Price,
		Quantity:  entity.Quantity,
		UserID:    entity.UserID,
	}
}
