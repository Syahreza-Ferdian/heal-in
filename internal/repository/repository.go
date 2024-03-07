package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    InterfaceUserRepository
	PaymentRepository InterfacePaymentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserRepository(db),
		PaymentRepository: NewPaymentRepository(db),
	}
}
