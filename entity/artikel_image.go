package entity

import "github.com/google/uuid"

type ArtikelImage struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	ArtikelID uuid.UUID `json:"artikel_id" gorm:"type:varchar(36);not null;foreignKey:ID;references:artikels;onUpdate:CASCADE;onDelete:CASCADE;"`
	Image     string    `json:"image" gorm:"type:varchar(255);not null;"`
}
