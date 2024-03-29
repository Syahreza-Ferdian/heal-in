package main

import (
	"log"
	"os"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/internal/handler/rest"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/database/mysql"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/middleware"
	"github.com/Syahreza-Ferdian/heal-in/pkg/scheduler"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	bcrypt_import "golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load("../../.env")
	env := os.Getenv("ENV")

	if err != nil && env == "" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := mysql.ConnectToDb()

	repository := repository.NewRepository(db)

	jwt := jwt.Init()

	bcrypt := bcrypt.NewBcrypt(bcrypt_import.DefaultCost)

	midtransSnapApi := snap.Client{}

	midtransSnapApi.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	midtransCoreApi := coreapi.Client{}

	midtransCoreApi.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	supabase := supabase.Init()

	mail := email.NewEmailSender(os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("EMAIL_FROM"))

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")

	schdl := scheduler.NewScheduler(jakartaTime)

	service := service.NewService(
		service.InitService{
			Repository: repository,
			Bcrypt:     bcrypt,
			JwtAuth:    jwt,
			SnapClient: midtransSnapApi,
			CoreApi:    midtransCoreApi,
			Supabase:   supabase,
			Scheduler: schdl,
			Mail: mail,
		},
	)

	middleware := middleware.Init(jwt, service)

	rest := rest.NewRest(gin.Default(), service, middleware, mail)

	mysql.Migrate(db)

	mysql.SeedData(db, &bcrypt)

	rest.EndPoint()

	rest.InitPaymentExpScheduler()

	rest.Start()
}
