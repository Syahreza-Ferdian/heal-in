package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) InterfaceUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := ur.db.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
