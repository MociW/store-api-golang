package service

import (
	"context"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	pgRepo user.UserPostgresRepository
}

func NewAuthService(pgRepo user.UserPostgresRepository) user.AuthService {
	return &AuthServiceImpl{pgRepo: pgRepo}
}

func (auth AuthServiceImpl) Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()

	user := &model.User{
		UserID:    id,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Username:  entity.Username,
		Email:     entity.Email,
		Password:  string(password),
	}

	result, err := auth.pgRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserResponse(result), nil
}

func (auth AuthServiceImpl) Login(ctx context.Context, entity *dto.UserLoginRequest) (*dto.JwtToken, error) {
	user := &model.User{
		Email:    entity.Email,
		Password: entity.Password,
	}

	result, err := auth.pgRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(entity.Password)); err != nil {
		return nil, err
	}

	accToken, refToken, err := util.GenerateTokenPair(user)
	if err != nil {
		return nil, errors.Wrap(err, "AuthService.Login")
	}

	return &dto.JwtToken{
		AccessToken:  accToken,
		RefreshToken: refToken,
	}, nil
}
