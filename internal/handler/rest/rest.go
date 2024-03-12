package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/middleware"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.MiddlewareInterface
	mail       email.EmailService
}

func NewRest(router *gin.Engine, service *service.Service, middleware middleware.MiddlewareInterface, mail email.EmailService) *Rest {
	return &Rest{
		router:     router,
		service:    service,
		middleware: middleware,
		mail:       mail,
	}
}

func (r *Rest) EndPoint() {
	r.router.Use(r.middleware.Cors)

	mainRouterGroup := r.router.Group("/api")

	user := mainRouterGroup.Group("/user")
	user.POST("/register", r.CreateUser)
	user.POST("/login", r.Login)
	user.GET("/login-user", r.middleware.AuthenticateUser, Testing)
	user.GET("/email/verify/:verificationCode", r.VerifyEmail)

	payment := mainRouterGroup.Group("/payment")
	payment.POST("/new", r.middleware.AuthenticateUser, r.NewPayment)
	payment.POST("/notification", r.PaymentNotification)

	artikel := mainRouterGroup.Group("/artikel")
	artikel.POST("/new", r.NewArtikel)
	artikel.GET("/:id", r.GetArtikel)
	artikel.GET("/all", r.middleware.AuthenticateUser, r.GetAllArtikel)
	artikel.GET("/sample", r.GetFewSampleArtikel)

	video := mainRouterGroup.Group("/video")
	video.POST("/new", r.NewVideo)
	video.GET("/all", r.middleware.AuthenticateUser, r.GetAllVideos)
	video.GET("/:id", r.GetSpecificVideo)
}

func (r *Rest) Start() {
	err := r.router.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_ADDRESS"), os.Getenv("APP_PORT")))

	if err != nil {
		log.Fatalf("Cannot running the server. Error: %v", err)
	}
}

func Testing(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		response.OnFailed(c, http.StatusOK, "Failed to get user", fmt.Errorf("failed to get user"))
		return
	}

	response.OnSuccess(c, http.StatusOK, "Success get user", user.(*entity.User))
}
