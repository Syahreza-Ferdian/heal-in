package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceJournalingQuestionRepository interface {
	GetJournalingQuestionText(id int) (string, error)
}

type JournalingQuestionRepository struct {
	db *gorm.DB
}

func NewJournalingQuestionRepository(db *gorm.DB) InterfaceJournalingQuestionRepository {
	return &JournalingQuestionRepository{
		db: db,
	}
}

func (jqr *JournalingQuestionRepository) GetJournalingQuestionText(id int) (string, error) {
	question := &entity.JournalingQuestion{}
	err := jqr.db.Debug().Where("id = ?", id).First(&question).Error

	if err != nil {
		return "", err
	}

	return question.Question, nil
}
