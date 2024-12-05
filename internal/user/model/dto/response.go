package dto

import (
	"github.com/MociW/store-api-golang/internal/user/model"
)

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
	Username    string            `json:"username"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	Avatar      string            `json:"avatar,omitempty"`
	PhoneNumber string            `json:"phone_number,omitempty"`
	Addresses   []AddressResponse `json:"addresses,omitempty"`
}

type AddressResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Street     string `json:"street"`
	Country    string `json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func ConvertUserResponse(entity *model.User) *UserResponse {

	responses := make([]AddressResponse, len(entity.Addresses))
	for i, address := range entity.Addresses {
		responses[i] = *ConvertAddressResponse(&address)
	}

	if len(responses) == 0 {
		responses = nil
	}

	return &UserResponse{
		UserID:      entity.UserID,
		Username:    entity.Username,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		Email:       entity.Email,
		Avatar:      entity.Avatar,
		PhoneNumber: entity.PhoneNumber,
		Addresses:   responses,
	}
}

func ConvertAddressResponse(entity *model.Address) *AddressResponse {
	return &AddressResponse{
		ID:         entity.ID,
		Title:      entity.Title,
		Street:     entity.Street,
		Country:    entity.Country,
		City:       entity.City,
		PostalCode: entity.PostalCode,
	}
}
