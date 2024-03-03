package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/google/uuid"
)

type InterfaceUserService interface {
	Register(userReq model.UserRegister) error
	// Login(userReq model.UserLogin) error
}

type UserService struct {
	ur     repository.InterfaceUserRepository
	bcrypt bcrypt.BcryptInterface
}

func NewUserService(ur repository.InterfaceUserRepository, bcrypt bcrypt.BcryptInterface) InterfaceUserService {
	return &UserService{
		ur: ur,
		bcrypt: bcrypt,
	}
}

func (us *UserService) Register(userReq model.UserRegister) error {
	hashPassword, err := us.bcrypt.HashPassword(userReq.Password)

	if err != nil {
		return err
	}

	userReq.ID = uuid.New()
	userReq.Password = hashPassword

	user := model.UserRegisterToEntity(userReq)

	_, err = us.ur.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}
