package service

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/google/uuid"
)

type InterfaceVideoService interface {
	NewVideo(video model.NewVideoRequest) (*entity.Video, error)
	GetAllVideos(userID string) ([]*entity.Video, error)
	GetVideo(id string) (*entity.Video, error)
	NewVideoWithLink(video model.NewVideoRequestAlt) (*entity.Video, error)
}

type VideoService struct {
	vr       repository.InterfaceVideoRepository
	ur       repository.InterfaceUserRepository
	supabase supabase.SupabaseInterface
}

func NewVideoService(vr repository.InterfaceVideoRepository, supabase supabase.SupabaseInterface, ur repository.InterfaceUserRepository) InterfaceVideoService {
	return &VideoService{
		vr:       vr,
		ur:       ur,
		supabase: supabase,
	}
}

func (vs *VideoService) NewVideo(video model.NewVideoRequest) (*entity.Video, error) {
	video.ID = uuid.New()

	link, err := vs.supabase.UploadFile(video.Video)

	if err != nil {
		return nil, err
	}

	videoEntity := &entity.Video{
		ID:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		Link:        link,
	}

	newVideo, err := vs.vr.NewVideo(videoEntity)

	if err != nil {
		return nil, err
	}

	return newVideo, nil
}

func (vs *VideoService) NewVideoWithLink(video model.NewVideoRequestAlt) (*entity.Video, error) {
	video.ID = uuid.New()

	videoEntity := &entity.Video{
		ID:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		Link:        video.Link,
	}

	newVideo, err := vs.vr.NewVideo(videoEntity)

	if err != nil {
		return nil, err
	}

	return newVideo, nil
}

func (vs *VideoService) GetAllVideos(userID string) ([]*entity.Video, error) {
	var status int
	var err error
	var limit bool

	status, err = vs.ur.GetUserSubscriptionStatus(userID)

	if err != nil {
		return nil, err
	}

	if status == 1 {
		limit = false
	} else {
		limit = true
	}

	videos, err := vs.vr.GetAllVideos(limit)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (vs *VideoService) GetVideo(id string) (*entity.Video, error) {
	video, err := vs.vr.GetVideo(id)

	if err != nil {
		return nil, err
	}

	return video, nil
}
