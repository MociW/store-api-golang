package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/MociW/store-api-golang/config"
	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/MociW/store-api-golang/pkg/email"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	cfg    *config.Config
	pgRepo user.UserPostgresRepository
	rdb    *redis.Client
	mail   email.EmailService
	rdb    *redis.Client
	mail   email.EmailService
}

func NewAuthService(cfg *config.Config, pgRepo user.UserPostgresRepository, rdb *redis.Client, mail email.EmailService) user.AuthService {
	return &AuthServiceImpl{cfg: cfg, pgRepo: pgRepo, rdb: rdb, mail: mail}
func NewAuthService(cfg *config.Config, pgRepo user.UserPostgresRepository, rdb *redis.Client, mail email.EmailService) user.AuthService {
	return &AuthServiceImpl{cfg: cfg, pgRepo: pgRepo, rdb: rdb, mail: mail}
}

func (auth *AuthServiceImpl) Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
func (auth *AuthServiceImpl) Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
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

	_, err = auth.pgRepo.CreateUser(ctx, user)
	_, err = auth.pgRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "AuthService.Register")
	}

	otp := util.GenerateRandomNumber(4)
	reference := util.GenerateRandomString(16)

	if err := auth.rdb.Set(ctx, "otp:"+reference, otp, 10*time.Minute).Err(); err != nil {
		return nil, errors.Wrap(err, "AuthService.Redis.Set")
	}

	if err := auth.rdb.Set(ctx, "user-ref:"+reference, user.Email, 10*time.Minute).Err(); err != nil {
		return nil, errors.Wrap(err, "AuthService.Redis.Set")
	}

	data := email.EmailData{
		Name: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		OTP:  otp,
	}

	err = auth.mail.Send(user.Email, "OTP CODE", data)
	if err != nil {
		return nil, errors.Wrap(err, "AuthService.Mail.Send")
	}

	return &dto.UserRegisterResponse{
		ReferenceID: reference,
	}, nil
}

// ValidateUser implements user.AuthService.
func (auth *AuthServiceImpl) ValidateUser(ctx context.Context, entity *dto.UserValidateRequest) error {
	token := strings.TrimSpace(entity.ReferenceID)

	val, err := auth.rdb.Get(ctx, "otp:"+token).Result()
	if err != nil {
		return errors.Wrap(err, "AuthService.Redis.Get")
	}

	if val != entity.OTP {
		return errors.New("invalid OTP")
	}

	val, err = auth.rdb.Get(ctx, "user-ref:"+token).Result()
	if err != nil {
		return errors.Wrap(err, "AuthService.Redis.Get")
	}

	request := &model.User{
		Email: val,
	}

	result, err := auth.pgRepo.FindByEmail(ctx, request)
	if err != nil {
		return errors.Wrap(err, "UserRepository.FindByEmail")
	}

	_, err = auth.pgRepo.UpdateUser(ctx, &model.User{UserID: result.UserID, VerifiedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	if err != nil {
		return errors.Wrap(err, "UserRepository.UpdateUser")
	}

	return nil
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
