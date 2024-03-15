package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfacePodcastRepository interface {
	NewPodcast(newPodcast *entity.Podcast) (*entity.Podcast, error)
	GetPodcastByID(id string) (*entity.Podcast, error)
	GetAllPodcasts(limit bool) ([]*entity.Podcast, error)
}

type PodcastRepository struct {
	db *gorm.DB
}

func NewPodcastRepository(db *gorm.DB) InterfacePodcastRepository {
	return &PodcastRepository{
		db: db,
	}
}

func (pr *PodcastRepository) NewPodcast(newPodcast *entity.Podcast) (*entity.Podcast, error) {
	err := pr.db.Debug().Create(&newPodcast).Error

	if err != nil {
		return nil, err
	}

	return newPodcast, nil
}

func (pr *PodcastRepository) GetPodcastByID(id string) (*entity.Podcast, error) {
	podcast := &entity.Podcast{}
	err := pr.db.Debug().Where("id = ?", id).First(&podcast).Error

	if err != nil {
		return nil, err
	}

	return podcast, nil
}

func (pr *PodcastRepository) GetAllPodcasts(limit bool) ([]*entity.Podcast, error) {
	var podcasts []*entity.Podcast

	if limit {
		err := pr.db.Debug().Limit(1).Find(&podcasts).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := pr.db.Debug().Find(&podcasts).Error
		if err != nil {
			return nil, err
		}
	}

	return podcasts, nil
}
