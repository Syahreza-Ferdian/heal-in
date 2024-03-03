package entity

import (
	"time"

	"github.com/google/uuid"
)

type Artikel struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null;"`
	Body      string    `json:"body" gorm:"type:text;not null;"`
	PostedAt  time.Time `json:"posted_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
