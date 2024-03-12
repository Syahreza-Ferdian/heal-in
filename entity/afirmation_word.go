package entity

import "github.com/google/uuid"

type AfirmationWord struct {
	ID     uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	MoodID int       `json:"mood_id" gorm:"type:int;not null;"`
	Word   string    `json:"word" gorm:"type:text;"`
}
