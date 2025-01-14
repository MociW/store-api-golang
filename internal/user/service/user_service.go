package service

import (
	"context"
	"fmt"

	"github.com/MociW/store-api-golang/config"
	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
)

type UserServiceImpl struct {
	cfg     *config.Config
	pgRepo  user.UserPostgresRepository
	awsRepo user.UserAWSRepository
}

func NewUserService(cfg *config.Config, pgRepo user.UserPostgresRepository, awsRepo user.UserAWSRepository) user.UserService {
	return &UserServiceImpl{cfg: cfg, pgRepo: pgRepo, awsRepo: awsRepo}
}

/* ---------------------------------- User ---------------------------------- */

func (user *UserServiceImpl) UpdateUser(ctx context.Context, entity *dto.UserUpdateRequest) (*dto.UserResponse, error) {

	request := &model.User{
		UserID:      entity.UserID,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		Email:       entity.Email,
		PhoneNumber: entity.PhoneNumber,
	}

	result, err := user.pgRepo.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserResponse(result), nil
}

func (user *UserServiceImpl) GetCurrentUser(ctx context.Context, id string) (*dto.UserResponse, error) {
	request := &model.User{
		Email: id,
	}

	result, err := user.pgRepo.GetCurrentUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserResponse(result), nil
}

func (user *UserServiceImpl) UploadAvatar(ctx context.Context, id string, file *model.UserUploadInput) (*dto.UserResponse, error) {
	uploadInfo, err := user.awsRepo.PutObject(ctx, file)
	if err != nil {
		return nil, err
	}

	avatarURL := user.generateAWSMinioURL(file.BucketName, uploadInfo.Key)
	updatedUser, err := user.pgRepo.UpdateUser(ctx, &model.User{UserID: id, Avatar: avatarURL})
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserResponse(updatedUser), nil
}

func (user *UserServiceImpl) generateAWSMinioURL(bucket string, key string) string {
	return fmt.Sprintf("%s/%s/%s", user.cfg.AWS.Endpoint, bucket, key)
}

/* --------------------------------- Address -------------------------------- */

func (user *UserServiceImpl) CreateAddress(ctx context.Context, entity *dto.CreateAddressRequest) (*dto.AddressResponse, error) {
	request := &model.Address{
		UserID:     entity.UserID,
		Title:      entity.Title,
		Street:     entity.Street,
		Country:    entity.Country,
		City:       entity.City,
		PostalCode: entity.PostalCode,
	}

	result, err := user.pgRepo.CreateAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertAddressResponse(result), nil
}

func (user *UserServiceImpl) UpdateAddress(ctx context.Context, entity *dto.UpdateAddressRequest) (*dto.AddressResponse, error) {
	request := &model.Address{
		ID:         entity.ID,
		UserID:     entity.UserID,
		Title:      entity.Title,
		Street:     entity.Street,
		Country:    entity.Country,
		City:       entity.City,
		PostalCode: entity.PostalCode,
	}

	result, err := user.pgRepo.UpdateAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertAddressResponse(result), nil
}

func (user *UserServiceImpl) DeleteAddress(ctx context.Context, entity *dto.DeleteAddressRequest) error {
	request := &model.Address{
		ID:     entity.ID,
		UserID: entity.UserID,
	}

	err := user.pgRepo.DeleteAddress(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (user *UserServiceImpl) FindAddress(ctx context.Context, entity *dto.FindAddressRequest) (*dto.AddressResponse, error) {
	request := &model.Address{
		ID:     entity.ID,
		UserID: entity.UserID,
	}

	result, err := user.pgRepo.FindAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertAddressResponse(result), nil
}

func (user *UserServiceImpl) ListAddress(ctx context.Context, id string) ([]dto.AddressResponse, error) {

	result, err := user.pgRepo.ListAddress(ctx, id)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.AddressResponse, len(result))
	for i, address := range result {
		responses[i] = *dto.ConvertAddressResponse(&address)
	}

	return responses, nil
}
