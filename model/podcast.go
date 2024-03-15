package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type NewPodcastRequest struct {
	ID          uuid.UUID             `form:"-"`
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Podcast     *multipart.FileHeader `form:"podcast"`
}
