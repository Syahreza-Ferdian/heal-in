package entity

import "github.com/google/uuid"

type EventImage struct {
	ID        uuid.UUID `json:"-" gorm:"type:varchar(36);primary_key;"`
	EventID   uuid.UUID `json:"-" gorm:"type:varchar(36);not null;"`
	ImageLink string    `json:"image_link" gorm:"type:varchar(255);not null;"`
}
