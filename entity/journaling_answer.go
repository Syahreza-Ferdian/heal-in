package entity

import "github.com/google/uuid"

type JournalingAnswer struct {
	ID         uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	EntryID    uuid.UUID `json:"entry_id" gorm:"type:varchar(36);"`
	QuestionID int       `json:"question_id" gorm:"type:int;"`
	Answer     string    `json:"answer" gorm:"type:text;not null;"`
}
