package repository

import (
	"context"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserPostgresRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) user.UserPostgresRepository {
	return &UserPostgresRepositoryImpl{DB: db}
}

/* ---------------------------------- User ---------------------------------- */

func (r *UserPostgresRepositoryImpl) CreateUser(ctx context.Context, entity *model.User) (*model.User, error) {

	tx := r.DB.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("email = ?", entity.Email).Omit("avatar", "phone_number").FirstOrCreate(entity)

		if result.RowsAffected == 0 {
			return gorm.ErrRegistered
		}

		if result.Error != nil {
			return errors.Wrap(result.Error, "UserPostgresRepository.Register.CreateUser")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *UserPostgresRepositoryImpl) UpdateUser(ctx context.Context, entity *model.User) (*model.User, error) {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.User{}).Where("user_id = ?", entity.UserID).Updates(entity)

		if result.RowsAffected == 0 {
			return gorm.ErrInvalidData
		}

		if result.Error != nil {
			return errors.Wrap(result.Error, "UserPostgresRepository.Update.UpdateUser")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *UserPostgresRepositoryImpl) DeleteUser(ctx context.Context, entity *model.User) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		user := new(model.User)
		if err := tx.Where("email = ?", entity.Email).First(user).Error; err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Delete.DeleteUser")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(entity.Password)); err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Delete.DeleteUser")
		}

		if err := tx.Delete(user).Error; err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Delete.DeleteUser")
		}

		return nil
	})
}

func (r *UserPostgresRepositoryImpl) FindByEmail(ctx context.Context, entity *model.User) (*model.User, error) {
	user := new(model.User)
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("email = ?", entity.Email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserPostgresRepositoryImpl) FindByUsername(ctx context.Context, entity *model.User) (*model.User, error) {
	user := new(model.User)
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("email = ?", entity.Email).First(user).Error; err != nil {
		return nil, errors.Wrap(err, "UserPostgresRepository.Find.FindByUsername")
	}

	return user, nil
}

// GetCurrentUser implements user.UserPostgresRepository.
func (r *UserPostgresRepositoryImpl) GetCurrentUser(ctx context.Context, entity *model.User) (*model.User, error) {
	user := new(model.User)
	tx := r.DB.WithContext(ctx)

	if err := tx.Model(&model.User{}).Preload("Addresses").Preload("Products").Take(user, "email = ?", entity.Email).Error; err != nil {
		return nil, errors.Wrap(err, "UserPostgresRepository.Find.FindByUsername")
	}

	return user, nil
}

/* --------------------------------- Address -------------------------------- */

func (r *UserPostgresRepositoryImpl) CreateAddress(ctx context.Context, entity *model.Address) (*model.Address, error) {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ?", entity.UserID).Create(entity).Error

		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Create.CreateAddress")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *UserPostgresRepositoryImpl) UpdateAddress(ctx context.Context, entity *model.Address) (*model.Address, error) {
	tx := r.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Address{}).Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Error; err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Update.UpdateAddress")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *UserPostgresRepositoryImpl) DeleteAddress(ctx context.Context, entity *model.Address) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Delete(&model.Address{}).Error

		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Delete.DeleteAddress")
		}

		return nil
	})
}

func (r *UserPostgresRepositoryImpl) FindAddress(ctx context.Context, entity *model.Address) (*model.Address, error) {
	address := new(model.Address)

	tx := r.DB.WithContext(ctx)

	if err := tx.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Take(address).Error; err != nil {
		return nil, errors.Wrap(err, "UserPostgresRepository.Find.FindAddress")
	}

	return entity, nil
}

func (r *UserPostgresRepositoryImpl) ListAddress(ctx context.Context, uuid string) ([]model.Address, error) {
	var Addresses []model.Address
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("user_id = ?", uuid).Find(&Addresses).Error; err != nil {
		return nil, errors.Wrap(err, "UserPostgresRepository.Find.FindAddress")
	}

	return Addresses, nil
}
