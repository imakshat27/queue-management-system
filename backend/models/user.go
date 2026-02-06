package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name        string    `json:"name"`
	Username    string    `json:"username" gorm:"uniqueIndex;not null"`
	Password    string    `json:"password"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null"`
	Phone       string    `json:"phone" gorm:"uniqueIndex"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
}
