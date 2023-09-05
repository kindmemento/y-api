package models

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationInput struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	gorm.Model
	Username     string         `gorm:"type:varchar(55);unique;not null"`
	FirstName    string         `gorm:"type:varchar(55);not null"`
	LastName     string         `gorm:"type:varchar(55);not null"`
	Email        string         `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
