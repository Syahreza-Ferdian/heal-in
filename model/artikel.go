package model

import (
	"mime/multipart"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/google/uuid"
)

type ArtikelUploadImageParam struct {
	Image *multipart.FileHeader `form:"image"`
}

type NewArtikelRequest struct {
	ID    uuid.UUID `form:"-"`
	Title string    `form:"title"`
	Body  string    `form:"body"`
}

func ArtikelRequestToEntity(ar NewArtikelRequest) entity.Artikel {
	return entity.Artikel{
		ID:    ar.ID,
		Title: ar.Title,
		Body:  ar.Body,
	}
}
