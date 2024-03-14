package entity

import (
	"time"

	"github.com/google/uuid"
)

type JournalingEntry struct {
	ID        uuid.UUID          `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID    uuid.UUID          `json:"user_id" gorm:"type:varchar(36);not null;foreignKey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE;"`
	CreatedAt time.Time          `json:"created_at" gorm:"autoCreateTime"`
	Answers   []JournalingAnswer `json:"answers" gorm:"foreignKey:EntryID;references:ID;onUpdate:CASCADE;onDelete:CASCADE;"`
	Mood      int                `json:"mood" gorm:"type:int;"`
	// Mood      []JournalingMood   `json:"mood" gorm:"foreignKey:EntryID;references:ID;onUpdate:CASCADE;onDelete:CASCADE;"`
}
