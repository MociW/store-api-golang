package user

import (
	"context"

	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
)

type AuthService interface {
	Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)

	ValidateUser(ctx context.Context, entity *dto.UserValidate) error

	Login(ctx context.Context, entity *dto.UserLoginRequest) (*dto.JwtToken, error)
}

type UserService interface {
	/* ---------------------------------- User ---------------------------------- */

	UpdateUser(ctx context.Context, entity *dto.UserUpdateRequest) (*dto.UserResponse, error)

	UploadAvatar(ctx context.Context, id string, file *model.UserUploadInput) (*dto.UserResponse, error)

	GetCurrentUser(ctx context.Context, id string) (*dto.UserResponse, error)

	// ChangePassword(ctx context.Context, entity *dto.UserUpdateRequest) (*dto.UserResponse, error)

	// ChangeEmail(ctx context.Context, entity *dto.UserUpdateRequest) (*dto.UserResponse, error)

	// ChangeUsername(ctx context.Context, entity *dto.UserUpdateRequest) (*dto.UserResponse, error)

	/* --------------------------------- Address -------------------------------- */

	CreateAddress(ctx context.Context, entity *dto.CreateAddressRequest) (*dto.AddressResponse, error)

	UpdateAddress(ctx context.Context, entity *dto.UpdateAddressRequest) (*dto.AddressResponse, error)

	DeleteAddress(ctx context.Context, entity *dto.DeleteAddressRequest) error

	FindAddress(ctx context.Context, entity *dto.FindAddressRequest) (*dto.AddressResponse, error)

	ListAddress(ctx context.Context, id string) ([]dto.AddressResponse, error)
}
