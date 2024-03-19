package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceEventRepository interface {
	CreateNewEvent(newEvent *entity.Event) (*entity.Event, error)
	GetEventByID(id string) (*entity.Event, error)
	GetAllEvents() ([]*entity.Event, error)
}

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) InterfaceEventRepository {
	return &EventRepository{
		db: db,
	}
}

func (er *EventRepository) CreateNewEvent(newEvent *entity.Event) (*entity.Event, error) {
	err := er.db.Debug().Create(&newEvent).Error

	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

func (er *EventRepository) GetEventByID(id string) (*entity.Event, error) {
	event := &entity.Event{}
	err := er.db.Debug().Where("id = ?", id).Preload("EventImage").First(&event).Error

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (er *EventRepository) GetAllEvents() ([]*entity.Event, error) {
	events := []*entity.Event{}
	err := er.db.Debug().Preload("EventImage").Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}
