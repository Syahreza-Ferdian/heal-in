package entity

type JournalingMood struct {
	ID              int              `json:"-" gorm:"type:int;primary_key;autoIncrement"`
	Mood            string           `json:"-" gorm:"type:varchar(255);not null;"`
	AfirmationWords []AfirmationWord `json:"afirmation_words" gorm:"foreignKey:MoodID;references:ID;onUpdate:CASCADE;onDelete:CASCADE;"`
}
