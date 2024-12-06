package model

import (
	productModel "github.com/MociW/store-api-golang/internal/product/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar      string                 `gorm:"column:avatar"`
	UserID      string                 `gorm:"column:user_id"`
	FirstName   string                 `gorm:"column:first_name"`
	LastName    string                 `gorm:"column:last_name"`
	Username    string                 `gorm:"column:username"`
	Email       string                 `gorm:"column:email"`
	Password    string                 `gorm:"column:password"`
	PhoneNumber string                 `gorm:"column:phone_number"`
	Addresses   []Address              `gorm:"foreignKey:user_id;references:user_id"`
	Products    []productModel.Product `gorm:"foreignKey:user_id;references:user_id"`
}

func (u *User) TableName() string {
	return "users"
}
