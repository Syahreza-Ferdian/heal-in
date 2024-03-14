package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceJournalingEntryRepository interface {
	NewJournalingEntry(newEntry *entity.JournalingEntry) (*entity.JournalingEntry, error)
	GetJournalingEntryByID(id string) (*entity.JournalingEntry, error)
}

type JournalingEntryRepository struct {
	db *gorm.DB
}

func NewJournalingEntryRepository(db *gorm.DB) InterfaceJournalingEntryRepository {
	return &JournalingEntryRepository{
		db: db,
	}
}

func (jer *JournalingEntryRepository) NewJournalingEntry(newEntry *entity.JournalingEntry) (*entity.JournalingEntry, error) {
	err := jer.db.Debug().Create(&newEntry).Error

	if err != nil {
		return nil, err
	}

	return newEntry, nil
}

func (jer *JournalingEntryRepository) GetJournalingEntryByID(id string) (*entity.JournalingEntry, error) {
	entry := &entity.JournalingEntry{}
	err := jer.db.Debug().Where("id = ?", id).Preload("Answers", func(db *gorm.DB) *gorm.DB { return db.Order("journaling_answers.id ASC") }).First(&entry).Error

	if err != nil {
		return nil, err
	}

	return entry, nil
}
