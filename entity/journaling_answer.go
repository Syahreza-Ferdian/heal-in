package entity

import "github.com/google/uuid"

type JournalingAnswer struct {
	ID         int       `json:"id" gorm:"type:int;primary_key;autoIncrement;"`
	EntryID    uuid.UUID `json:"entry_id" gorm:"type:varchar(36);"`
	QuestionID int       `json:"question_id" gorm:"type:int;"`
	Answer     string    `json:"answer" gorm:"type:text;not null;"`
}
