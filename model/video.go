package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type NewVideoRequest struct {
	ID          uuid.UUID             `form:"-"`
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Video       *multipart.FileHeader `form:"video"`
}

type NewVideoRequestAlt struct {
	ID          uuid.UUID `form:"-"`
	Title       string    `form:"title"`
	Description string    `form:"description"`
	Link        string    `form:"link"`
}
