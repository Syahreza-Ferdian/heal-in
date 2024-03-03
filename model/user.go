package model

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/google/uuid"
)

type UserRegister struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRegisterToEntity(ur UserRegister) entity.User {
	return entity.User{
		ID:       ur.ID,
		Name:     ur.Name,
		Email:    ur.Email,
		Password: ur.Password,
	}
}
