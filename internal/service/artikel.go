package service

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/google/uuid"
)

type InterfaceArtikelService interface {
	NewArtikel(artikelReq model.NewArtikelRequest) (*entity.Artikel, error)
	GetArtikel(id string) (*entity.Artikel, error)
	UploadArtikelImage(artikelID string, artikelImageReq model.ArtikelUploadImageParam) error
	GetAllArtikel(userID string) ([]*entity.Artikel, error)
	GetFewSampleArtikel() ([]*entity.Artikel, error)
}

type ArtikelService struct {
	ar       repository.InterfaceArtikelRepository
	air      repository.InterfaceArtikelImageRepository
	ur       repository.InterfaceUserRepository
	supabase supabase.SupabaseInterface
}

func NewArtikelService(ar repository.InterfaceArtikelRepository, supabase supabase.SupabaseInterface, air repository.InterfaceArtikelImageRepository, ur repository.InterfaceUserRepository) InterfaceArtikelService {
	return &ArtikelService{
		ar:       ar,
		supabase: supabase,
		air:      air,
		ur:       ur,
	}
}

func (as *ArtikelService) NewArtikel(artikelReq model.NewArtikelRequest) (*entity.Artikel, error) {
	artikelReq.ID = uuid.New()

	artikelEntity := model.ArtikelRequestToEntity(artikelReq)

	artikel, err := as.ar.CreateArtikel(&artikelEntity)

	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	return artikel, nil
}

func (as *ArtikelService) GetArtikel(id string) (*entity.Artikel, error) {
	artikel, err := as.ar.GetArtikel(id)

	if err != nil {
		return nil, err
	}

	return artikel, nil
}

func (as *ArtikelService) UploadArtikelImage(artikelID string, artikelImageReq model.ArtikelUploadImageParam) error {
	artikel, err := as.ar.GetArtikel(artikelID)

	if err != nil {
		return err
	}

	link, err := as.supabase.UploadFile(artikelImageReq.Image)

	if err != nil {
		return err
	}

	artikelEntity := &entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel.ID,
		Image:     link,
	}

	_, err = as.air.CreateArtikelImage(artikelEntity)

	if err != nil {
		err1 := as.supabase.DeleteFile(link)
		
		if err1 != nil {
			return err
		}
		return err
	}

	return nil
}

func (as *ArtikelService) GetAllArtikel(userID string) ([]*entity.Artikel, error) {
	var status int
	var err error

	status, err = as.ur.GetUserSubscriptionStatus(userID)

	if err != nil {
		return nil, err
	}

	var moreContent bool

	if status == 0 {
		moreContent = false
	} else if status == 1 {
		moreContent = true
	}

	artikels, err := as.ar.GetArtikelByUserStatus(moreContent)
	if err != nil {
		return nil, err
	}

	return artikels, nil
}

func (as *ArtikelService) GetFewSampleArtikel() ([]*entity.Artikel, error) {
	artikels, err := as.ar.GetArtikelByUserStatus(false)

	if err != nil {
		return nil, err
	}

	return artikels, nil
}
