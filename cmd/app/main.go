package main

import (
	"os"

	"github.com/Syahreza-Ferdian/heal-in/internal/handler/rest"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/config"
	"github.com/Syahreza-Ferdian/heal-in/pkg/database/mysql"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	bcrypt_import "golang.org/x/crypto/bcrypt"
)

func main() {
	config.LoadEnv()

	db := mysql.ConnectToDb()

	repository := repository.NewRepository(db)

	jwt := jwt.Init()

	bcrypt := bcrypt.NewBcrypt(bcrypt_import.DefaultCost)

	midtransSnapApi := snap.Client{}

	midtransSnapApi.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	midtransCoreApi := coreapi.Client{}

	midtransCoreApi.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Production)

	service := service.NewService(
		service.InitService{
			Repository: repository,
			Bcrypt:     bcrypt,
			JwtAuth:    jwt,
			SnapClient: midtransSnapApi,
			CoreApi:    midtransCoreApi,
		},
	)

	middleware := middleware.Init(jwt, service)

	mail := email.NewEmailSender(os.Getenv("SMTP_USER"), os.Getenv("EMAIL_FROM"), os.Getenv("SMTP_PASS"))

	rest := rest.NewRest(gin.Default(), service, middleware, mail)

	mysql.Migrate(db)

	rest.EndPoint()

	rest.Start()
}
