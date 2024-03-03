package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
)

type Service struct {
	UserService InterfaceUserService
}

type InitService struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.BcryptInterface
}

func NewService(param InitService) *Service {
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository, param.Bcrypt),
	}
}
