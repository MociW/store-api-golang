package dto

import "github.com/lib/pq"

type ApiProductResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ProductResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Summary     string         `json:"summary"`
	Images      pq.StringArray `json:"images"`
	UserID      string         `json:"user_id"`
}

type ProductSKUResponese struct {
	ID uint `json:"id"`
}
