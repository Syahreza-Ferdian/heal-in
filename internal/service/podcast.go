package service

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/google/uuid"
)

type InterfacePodcastService interface {
	NewPodcast(podcastReq model.NewPodcastRequest) (*entity.Podcast, error)
	GetAllPodcastsBasedOnUserStatus(userID string) ([]*entity.Podcast, error)
	GetPodcastByID(podcastID string) (*entity.Podcast, error)
}

type PodcastService struct {
	pdr      repository.InterfacePodcastRepository
	ur       repository.InterfaceUserRepository
	supabase supabase.SupabaseInterface
}

func NewPodcastService(pdr repository.InterfacePodcastRepository, spb supabase.SupabaseInterface, ur repository.InterfaceUserRepository) InterfacePodcastService {
	return &PodcastService{
		pdr:      pdr,
		ur:       ur,
		supabase: spb,
	}
}

func (ps *PodcastService) NewPodcast(podcastReq model.NewPodcastRequest) (*entity.Podcast, error) {
	podcastReq.ID = uuid.New()

	podcastLink, err := ps.supabase.UploadFile(podcastReq.Podcast)
	if err != nil {
		return nil, err
	}

	newPodcastEntity := &entity.Podcast{
		ID:          podcastReq.ID,
		Title:       podcastReq.Title,
		Link:        podcastLink,
		Description: podcastReq.Description,
	}

	newPodcast, err := ps.pdr.NewPodcast(newPodcastEntity)
	if err != nil {
		return nil, err
	}

	return newPodcast, nil
}

func (ps *PodcastService) GetAllPodcastsBasedOnUserStatus(userID string) ([]*entity.Podcast, error) {
	var status int
	var err error
	var isLimited bool

	status, err = ps.ur.GetUserSubscriptionStatus(userID)
	if err != nil {
		return nil, err
	}

	if status == 1 {
		isLimited = false
	} else {
		isLimited = true
	}

	podcasts, err := ps.pdr.GetAllPodcasts(isLimited)
	if err != nil {
		return nil, err
	}

	return podcasts, nil
}

func (ps *PodcastService) GetPodcastByID(podcastID string) (*entity.Podcast, error) {
	podcast, err := ps.pdr.GetPodcastByID(podcastID)
	if err != nil {
		return nil, err
	}

	return podcast, nil
}
