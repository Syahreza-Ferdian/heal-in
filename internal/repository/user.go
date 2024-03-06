package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(param *model.GetUserParam) (*entity.User, error)
	GetUserColoumn(colName string, value string) (*entity.User, error)
	UpdateUserData(updatedUser *entity.User) error
	DeleteUser(user entity.User) error
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

func (ur *UserRepository) GetUser(param *model.GetUserParam) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.Debug().Where(&param).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserColoumn(colName string, value string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Debug().Where(colName+" = ?", value).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUserData(updatedUser *entity.User) error {
	err := ur.db.Debug().Save(updatedUser).Error

	return err
}

func (ur *UserRepository) DeleteUser(user entity.User) error {
	err := ur.db.Debug().Delete(user).Error

	return err
}
