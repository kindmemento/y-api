package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountType string

const (
	PersonalAccount AccountType = "personal"
	TeamAccount     AccountType = "team"
)

type RegistrationInput struct {
	Username    string      `json:"username"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	AccountType AccountType `json:"account_type"`
}

type LoginInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Username     string         `gorm:"type:varchar(55);unique;not null" validate:"required"`
	FirstName    string         `gorm:"type:varchar(55);not null"`
	LastName     string         `gorm:"type:varchar(55);not null"`
	Email        string         `gorm:"type:varchar(100);unique;not null" validate:"required,email"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	AccountType  string         `gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
