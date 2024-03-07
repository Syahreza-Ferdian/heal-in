package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type Service struct {
	UserService    InterfaceUserService
	PaymentService InterfacePaymentService
}

type InitService struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.BcryptInterface
	JwtAuth    jwt.JWTInterface
	SnapClient snap.Client
	CoreApi    coreapi.Client
}

func NewService(param InitService) *Service {
	return &Service{
		UserService:    NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth),
		PaymentService: NewPaymentService(param.Repository.PaymentRepository, param.SnapClient, param.CoreApi),
	}
}
