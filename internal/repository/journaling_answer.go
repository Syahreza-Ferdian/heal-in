package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceJournalingAnsRepository interface {
	NewJournalingAns(newAnswer *entity.JournalingAnswer) (*entity.JournalingAnswer, error)
}

type JournalingAnsRepository struct {
	db *gorm.DB
}

func NewJournalingAnsRepository(db *gorm.DB) InterfaceJournalingAnsRepository {
	return &JournalingAnsRepository{
		db: db,
	}
}

func (jar *JournalingAnsRepository) NewJournalingAns(newAnswer *entity.JournalingAnswer) (*entity.JournalingAnswer, error) {
	err := jar.db.Debug().Create(&newAnswer).Error

	if err != nil {
		return nil, err
	}

	return newAnswer, nil
}
