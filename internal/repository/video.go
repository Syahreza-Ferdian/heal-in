package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceVideoRepository interface {
	NewVideo(video *entity.Video) (*entity.Video, error)
	GetVideo(id string) (*entity.Video, error)
	GetAllVideos(limit bool) ([]*entity.Video, error)
}

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) InterfaceVideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (vr *VideoRepository) NewVideo(video *entity.Video) (*entity.Video, error) {
	err := vr.db.Debug().Create(&video).Error

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (vr *VideoRepository) GetVideo(id string) (*entity.Video, error) {
	video := &entity.Video{}
	err := vr.db.Debug().Where("id = ?", id).First(&video).Error

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (vr *VideoRepository) GetAllVideos(limit bool) ([]*entity.Video, error) {
	var videos []*entity.Video

	if limit {
		err := vr.db.Debug().Limit(1).Find(&videos).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := vr.db.Debug().Find(&videos).Error
		if err != nil {
			return nil, err
		}
	}

	return videos, nil
}
