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
func (r UserPostgresRepositoryImpl) CreateUser(ctx context.Context, entity *model.User) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("email = ?", entity.Email).Omit("avatar", "phone_number").FirstOrCreate(entity).Error

		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Register.CreateUser")
		}
		return nil
	})
}

func (r UserPostgresRepositoryImpl) UpdateUser(ctx context.Context, entity *model.User) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.User{}).Where("user_id = ?", entity.UserID).Updates(entity).Error
		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Update.UpdateUser")
		}
		return nil
	})
}

func (r UserPostgresRepositoryImpl) DeleteUser(ctx context.Context, entity *model.User) error {
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

func (r UserPostgresRepositoryImpl) FindByEmail(ctx context.Context, entity *model.User, email string) error {
	tx := r.DB.WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("email = ?", email).First(entity).Error; err != nil {
		return err
	}

	return nil
}

func (r UserPostgresRepositoryImpl) FindByUsername(ctx context.Context, entity *model.User, username string) error {
	tx := r.DB.WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("email = ?", username).First(entity).Error; err != nil {
		return err
	}

	return nil
}

/* --------------------------------- Address -------------------------------- */
func (r UserPostgresRepositoryImpl) CreateAddress(ctx context.Context, entity *model.Address) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ?", entity.UserID).Create(entity).Error

		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Create.CreateAddress")
		}
		return nil
	})
}

func (r UserPostgresRepositoryImpl) UpdateAddress(ctx context.Context, entity *model.Address) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Address{}).Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Error; err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Update.UpdateAddress")
		}
		return nil
	})
}

func (r UserPostgresRepositoryImpl) DeleteAddress(ctx context.Context, entity *model.Address) error {
	tx := r.DB.WithContext(ctx)

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ? AND user_id = ?", entity.ID, entity.UserID).Delete(&model.Address{}).Error

		if err != nil {
			return errors.Wrap(err, "UserPostgresRepository.Delete.DeleteAddress")
		}

		return nil
	})
}

func (r UserPostgresRepositoryImpl) FindAddress(ctx context.Context, entity *model.Address, uuid string, id uint) error {
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("id = ? AND user_id = ?", id, uuid).Take(entity).Error; err != nil {
		return errors.Wrap(err, "UserPostgresRepository.Find.FindAddress")
	}

	return nil
}

func (r UserPostgresRepositoryImpl) ListAddress(ctx context.Context, uuid string) ([]model.Address, error) {
	var Addresses []model.Address
	tx := r.DB.WithContext(ctx)

	if err := tx.Where("user_id = ?", uuid).Find(&Addresses).Error; err != nil {
		return nil, errors.Wrap(err, "UserPostgresRepository.Find.FindAddress")
	}

	return Addresses, nil
}
