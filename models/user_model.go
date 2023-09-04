package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username      string `gorm:"type:varchar(55);unique;not null"`
	first_name    string `gorm:"type:varchar(55);not null"`
	last_name     string `gorm:"type:varchar(55);not null"`
	email         string `gorm:"type:varchar(100);unique;not null"`
	password_hash string `gorm:"type:varchar(255);not null"`
	created_at    time.Time
	updated_at    time.Time
}
