package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceArtikelRepository interface {
	CreateArtikel(artikel *entity.Artikel) (*entity.Artikel, error)
	GetArtikel(id string) (*entity.Artikel, error)
	GetArtikelByUserStatus(moreContent bool) ([]*entity.Artikel, error)
}

type ArtikelRepository struct {
	db *gorm.DB
}

func NewArtikelRepository(db *gorm.DB) InterfaceArtikelRepository {
	return &ArtikelRepository{
		db: db,
	}
}

func (ar *ArtikelRepository) CreateArtikel(artikel *entity.Artikel) (*entity.Artikel, error) {
	err := ar.db.Debug().Create(&artikel).Error

	if err != nil {
		return nil, err
	}

	return artikel, nil
}

func (ar *ArtikelRepository) GetArtikel(id string) (*entity.Artikel, error) {
	artikel := &entity.Artikel{}
	err := ar.db.Debug().Where("id = ?", id).Preload("ArtikelImage").First(&artikel).Error

	if err != nil {
		return nil, err
	}

	return artikel, nil
}

func (ar *ArtikelRepository) GetArtikelByUserStatus(moreContent bool) ([]*entity.Artikel, error) {
	var artikels []*entity.Artikel

	if !moreContent {
		err := ar.db.Debug().Limit(5).Preload("ArtikelImage").Find(&artikels).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := ar.db.Debug().Preload("ArtikelImage").Find(&artikels).Error
		if err != nil {
			return nil, err
		}
	}

	return artikels, nil
}
