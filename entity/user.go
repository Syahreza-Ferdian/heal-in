package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name             string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email            string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password         string    `json:"password" gorm:"type:varchar(255);not null;"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IsEmailVerified  bool      `json:"is_email_verified" gorm:"default:false"`
	VerificationCode string    `json:"-" gorm:"type:varchar(255)"`
	IsSubscribed     bool      `json:"is_subscribed" gorm:"default:false"`
}
