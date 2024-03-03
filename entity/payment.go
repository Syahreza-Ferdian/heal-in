package entity

import "github.com/google/uuid"

type Payment struct {
	ID     uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE;"`
	Amount int       `json:"amount" gorm:"type:int;not null;"`
}
