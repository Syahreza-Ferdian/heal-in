package entity

import "github.com/google/uuid"

type PaymentEvent struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	EventID     uuid.UUID `json:"event_id" gorm:"type:varchar(36);not null;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null;"`
	Amount      int       `json:"amount" gorm:"type:int;not null;"`
	Description string    `json:"description" gorm:"type:text;not null;"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
}
