package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	Amount      int       `json:"amount" gorm:"type:int;not null;"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	ExpiredAt   time.Time `json:"expired_at" gorm:"type:timestamp"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
}
