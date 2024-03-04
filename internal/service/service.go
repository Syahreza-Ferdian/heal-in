package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
)

type Service struct {
	UserService InterfaceUserService
}

type InitService struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.BcryptInterface
	JwtAuth    jwt.JWTInterface
}

func NewService(param InitService) *Service {
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth),
	}
}
