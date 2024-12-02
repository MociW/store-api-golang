package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID     uint   `gorm:"column:user_id"`
	Title      string `gorm:"column:title"`
	Street     string `gorm:"column:street"`
	Country    string `gorm:"column:country"`
	City       string `gorm:"column:city"`
	PostalCode string `gorm:"column:postal_code"`
	User       *User  `gorm:"foreignKey:user_id;references:user_id"`
}
