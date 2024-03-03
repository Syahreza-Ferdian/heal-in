package main

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/handler/rest"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/config"
	"github.com/Syahreza-Ferdian/heal-in/pkg/database/mysql"
	"github.com/gin-gonic/gin"

	bcrypt_import "golang.org/x/crypto/bcrypt"
)

func main() {
	config.LoadEnv()

	db := mysql.ConnectToDb()

	repository := repository.NewRepository(db)

	bcrypt := bcrypt.NewBcrypt(bcrypt_import.DefaultCost)

	service := service.NewService(service.InitService{Repository: repository, Bcrypt: bcrypt})

	rest := rest.NewRest(gin.Default(), *service)

	mysql.Migrate(db)

	rest.EndPoint()

	rest.Start()
}
