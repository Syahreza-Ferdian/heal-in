package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository         InterfaceUserRepository
	PaymentRepository      InterfacePaymentRepository
	ArtikelRepository      InterfaceArtikelRepository
	ArtikelImageRepository InterfaceArtikelImageRepository
	VideoRepository        InterfaceVideoRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(db),
		PaymentRepository:      NewPaymentRepository(db),
		ArtikelRepository:      NewArtikelRepository(db),
		ArtikelImageRepository: NewArtikelImageRepository(db),
		VideoRepository:        NewVideoRepository(db),
	}
}
