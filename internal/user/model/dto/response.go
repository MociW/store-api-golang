package dto

import "github.com/MociW/store-api-golang/internal/user/model"

type ApiUserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type UserTokenResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Jwt     JwtToken `json:"jwt"`
}

type JwtToken struct {
	AccessToken  string `json:"access_Token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	UserID      string            `json:"user_id"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Avatar      string            `json:"avatar,omitempty"`
	PhoneNumber string            `json:"phone_number,omitempty"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
	Addresses   []AddressResponse `json:"addresses,omitempty"`
}

type AddressResponse struct {
	ID         uint   `json:"id"`
	UserID     string `json:"user_id"`
	Title      string `json:"title"`
	Street     string `json:"street"`
	Country    string `json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func ConvertUserResponse(entity *model.User) *UserResponse {
	return &UserResponse{
		UserID:      entity.UserID,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		Email:       entity.Email,
		Password:    entity.Password,
		Avatar:      entity.Avatar,
		PhoneNumber: entity.PhoneNumber,
		CreatedAt:   entity.CreatedAt.Unix(),
		UpdatedAt:   entity.UpdatedAt.Unix(),
	}
}

func ConvertAddressResponse(entity *model.Address) *AddressResponse {
	return &AddressResponse{
		ID:         entity.ID,
		UserID:     entity.UserID,
		Title:      entity.Title,
		Street:     entity.Street,
		Country:    entity.Country,
		City:       entity.City,
		PostalCode: entity.PostalCode,
	}
}
