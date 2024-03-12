package entity

type JournalingQuestion struct {
	ID       int                `json:"id" gorm:"type:int;primary_key;autoIncrement"`
	Question string             `json:"question" gorm:"type:varchar(255);not null;"`
	Answers  []JournalingAnswer `json:"answers" gorm:"foreignKey:QuestionID;references:ID;onUpdate:CASCADE;onDelete:CASCADE;"`
}
