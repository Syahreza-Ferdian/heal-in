package service

import (
	"fmt"
	"os"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/random"
	"github.com/google/uuid"
)

type InterfaceUserService interface {
	Register(userReq model.UserRegister) (entity.User, email.EmailData, error)
	Login(userReq model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param *model.GetUserParam) (*entity.User, error)
	Verify(verifCode string) error
	DeleteUser(user *entity.User) error
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

// TODO : Register nge return entity.User, email.EmailData, error
func (us *UserService) Register(userReq model.UserRegister) (entity.User, email.EmailData, error) {
	hashPassword, err := us.bcrypt.HashPassword(userReq.Password)

	if err != nil {
		return entity.User{}, email.EmailData{}, err
	}

	userReq.ID = uuid.New()
	userReq.Password = hashPassword

	user := model.UserRegisterToEntity(userReq)

	verificationCode, _ := random.GenerateRandomString(10)

	// EMAIL

	user.VerificationCode = verificationCode

	_, err = us.ur.CreateUser(&user)

	if err != nil {
		return entity.User{}, email.EmailData{}, err
	}

	return user,
		email.EmailData{
			RedirectURL: fmt.Sprintf("%s/api/user/email/verify/%s", os.Getenv("APP_URL"), verificationCode),
			FirstName:   user.Name,
			Subject:     "Verifikasi Email Anda",
			WebURL:      os.Getenv("APP_URL"),
		}, nil
}

func (us *UserService) Login(userReq model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := us.ur.GetUser(&model.GetUserParam{
		Email: userReq.Email,
	})

	if err != nil {
		return result, err
	}

	if !user.IsEmailVerified {
		return result, fmt.Errorf("email belum terverifikasi. silakan cek email anda untuk verifikasi")
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

func (us *UserService) Verify(verifCode string) error {
	user, err := us.ur.GetUserColoumn("verification_code", verifCode)

	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return fmt.Errorf("email sudah terverifikasi")
	}

	user.IsEmailVerified = true

	err = us.ur.UpdateUserData(user)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(user *entity.User) error {
	return us.ur.DeleteUser(*user)
}
