package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceEventImageRepository interface {
	CreateEventImage(image *entity.EventImage) (*entity.EventImage, error)
	GetEventImage(id string) (*entity.EventImage, error)
}

type EventImageRepository struct {
	db *gorm.DB
}

func NewEventImageRepository(db *gorm.DB) InterfaceEventImageRepository {
	return &EventImageRepository{
		db: db,
	}
}

func (eir *EventImageRepository) CreateEventImage(image *entity.EventImage) (*entity.EventImage, error) {
	err := eir.db.Debug().Create(&image).Error

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (eir *EventImageRepository) GetEventImage(id string) (*entity.EventImage, error) {
	image := &entity.EventImage{}
	err := eir.db.Debug().Where("id = ?", id).First(&image).Error

	if err != nil {
		return nil, err
	}

	return image, nil
}
