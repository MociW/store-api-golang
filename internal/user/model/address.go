package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         uint   `gorm:"primarykey"`
	UserID     string `gorm:"column:user_id"`
	Title      string `gorm:"column:title"`
	Street     string `gorm:"column:street"`
	Country    string `gorm:"column:country"`
	City       string `gorm:"column:city"`
	PostalCode string `gorm:"column:postal_code"`
	User       *User  `gorm:"foreignKey:user_id;references:user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (a *Address) TableName() string {
	return "addresses"
}
