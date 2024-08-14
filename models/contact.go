package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT"`
	FullName string `json:"full_name" binding:"required" gorm:"unique;not null;type:varchar(255)"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers" gorm:"foreignKey:ContactID"`
}

type PhoneNumber struct {
	gorm.Model
	ID int `gorm:"AUTO_INCREMENT"`
	ContactID uint   `json:"-"`
	Number    string `json:"number" binding:"required"`
}