package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceAfirmationWordRepository interface {
	GetRandomWordByMoodID(moodID int) (*entity.AfirmationWord, error)
}

type AfirmationWordRepository struct {
	db *gorm.DB
}

func NewAfirmationWordRepository(db *gorm.DB) InterfaceAfirmationWordRepository {
	return &AfirmationWordRepository{
		db: db,
	}
}

func (awr *AfirmationWordRepository) GetRandomWordByMoodID(moodID int) (*entity.AfirmationWord, error) {
	word := &entity.AfirmationWord{}

	err := awr.db.Debug().Where("mood_id = ?", moodID).Order("RAND()").First(&word).Error

	if err != nil {
		return nil, err
	}

	return word, nil
}
