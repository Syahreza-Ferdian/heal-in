package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceArtikelImageRepository interface {
	CreateArtikelImage(image *entity.ArtikelImage) (*entity.ArtikelImage, error)
	GetArtikelImage(id string) (*entity.ArtikelImage, error)
}

type ArtikelImageRepository struct {
	db *gorm.DB
}

func NewArtikelImageRepository(db *gorm.DB) InterfaceArtikelImageRepository {
	return &ArtikelImageRepository{
		db: db,
	}
}

func (ar *ArtikelImageRepository) CreateArtikelImage(image *entity.ArtikelImage) (*entity.ArtikelImage, error) {
	err := ar.db.Debug().Create(&image).Error

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (ar *ArtikelImageRepository) GetArtikelImage(id string) (*entity.ArtikelImage, error) {
	image := &entity.ArtikelImage{}
	err := ar.db.Debug().Where("id = ?", id).First(&image).Error

	if err != nil {
		return nil, err
	}

	return image, nil
}
