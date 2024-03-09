package model

import (
	"fmt"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/google/uuid"
)

type UserRegister struct {
	ID        uuid.UUID `json:"-"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func UserRegisterToEntity(ur UserRegister) entity.User {
	return entity.User{
		ID:       ur.ID,
		Name:     fmt.Sprintf("%s %s", ur.FirstName, ur.LastName),
		Email:    ur.Email,
		Password: ur.Password,
	}
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type GetUserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}
