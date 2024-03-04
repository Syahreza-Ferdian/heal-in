package service

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/google/uuid"
)

type InterfaceUserService interface {
	Register(userReq model.UserRegister) error
	Login(userReq model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param *model.GetUserParam) (*entity.User, error)
}

type UserService struct {
	ur      repository.InterfaceUserRepository
	bcrypt  bcrypt.BcryptInterface
	jwtAuth jwt.JWTInterface
}

func NewUserService(ur repository.InterfaceUserRepository, bcrypt bcrypt.BcryptInterface, jwtAuth jwt.JWTInterface) InterfaceUserService {
	return &UserService{
		ur:      ur,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
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

func (us *UserService) Login(userReq model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := us.ur.GetUser(&model.GetUserParam{
		Email: userReq.Email,
	})

	if err != nil {
		return result, err
	}

	err = us.bcrypt.ComparePassword(user.Password, userReq.Password)

	if err != nil {
		return result, err
	}

	token, err := us.jwtAuth.GenerateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (us *UserService) GetUser(param *model.GetUserParam) (*entity.User, error) {
	return us.ur.GetUser(param)
}
