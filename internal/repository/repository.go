package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository InterfaceUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
