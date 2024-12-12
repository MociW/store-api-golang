package service

import (
	"context"

	"github.com/MociW/store-api-golang/config"
	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	cfg    *config.Config
	pgRepo user.UserPostgresRepository
}

func NewAuthService(cfg *config.Config, pgRepo user.UserPostgresRepository) user.AuthService {
	return &AuthServiceImpl{cfg: cfg, pgRepo: pgRepo}
}

func (auth *AuthServiceImpl) Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
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
		return nil, errors.Wrap(err, "AuthService.Register")
	}

	return dto.ConvertUserResponse(result), nil
}

func (auth *AuthServiceImpl) Login(ctx context.Context, entity *dto.UserLoginRequest) (*dto.JwtToken, error) {
	user := &model.User{
		Email:    entity.Email,
		Password: entity.Password,
	}

	result, err := auth.pgRepo.FindByEmail(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound // Propagate "user not found" error
		}
		return nil, errors.Wrap(err, "AuthService.Login: error finding user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(entity.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, bcrypt.ErrMismatchedHashAndPassword // Propagate "invalid password" error
		}
		return nil, errors.Wrap(err, "AuthService.Login: password comparison failed")
	}

	accToken, refToken, err := util.GenerateTokenPair(result, auth.cfg)
	if err != nil {
		return nil, errors.Wrap(err, "AuthService.Login: token generation failed")
	}

	return &dto.JwtToken{
		AccessToken:  accToken,
		RefreshToken: refToken,
	}, nil
}
