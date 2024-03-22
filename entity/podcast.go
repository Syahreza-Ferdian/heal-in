package entity

import (
	"time"

	"github.com/google/uuid"
)

type Podcast struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null;"`
	Link        string    `json:"link" gorm:"type:varchar(255);not null;"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Thumbnail   string    `json:"thumbnail" gorm:"type:varchar(255)"`
}
