package service

import (
	"fmt"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterfaceArtikelService interface {
	NewArtikel(artikelReq model.NewArtikelRequest) (*entity.Artikel, error)
	GetArtikel(id string) (*entity.Artikel, error)
	UploadArtikelImage(artikelID string, artikelImageReq model.ArtikelUploadImageParam) error
	GetAllArtikel(c *gin.Context) ([]*entity.Artikel, error)
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
		return err
	}

	return nil
}

func (as *ArtikelService) GetAllArtikel(c *gin.Context) ([]*entity.Artikel, error) {
	currUser, ada := c.Get("user")

	var userID uuid.UUID
	var status int
	var err error

	if !ada {
		return nil, fmt.Errorf("failed to get current user")
	} else {
		userID = currUser.(*entity.User).ID
		status, err = as.ur.GetUserSubscriptionStatus(userID.String())

		if err != nil {
			return nil, err
		}
	}

	fmt.Println(ada)
	fmt.Println(status)

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
