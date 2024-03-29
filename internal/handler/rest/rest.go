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

	r.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	mainRouterGroup := r.router.Group("/api")

	user := mainRouterGroup.Group("/user")
	user.POST("/register", r.CreateUser)
	user.POST("/login", r.Login)
	user.GET("/login-user", r.middleware.AuthenticateUser, Testing)
	user.GET("/email/verify/:verificationCode", r.VerifyEmail)
	user.GET("/journaling/get", r.middleware.AuthenticateUser, r.GetCurrentUserJournalingEntries)
	user.GET("/events/get", r.middleware.AuthenticateUser, r.GetUserEvents)

	payment := mainRouterGroup.Group("/payment")
	payment.POST("/new", r.middleware.AuthenticateUser, r.NewPayment)
	payment.POST("/notification", r.PaymentNotification)

	artikel := mainRouterGroup.Group("/artikel")
	artikel.POST("/new", r.NewArtikel)
	artikel.GET("/:id", r.GetArtikel)
	artikel.GET("/all", r.middleware.AuthenticateUser, r.GetAllArtikel)
	artikel.GET("/sample", r.GetFewSampleArtikel)

	video := mainRouterGroup.Group("/video")
	video.POST("/new/upload", r.NewVideo)
	video.POST("/new/link", r.NewVideoWithLink)
	video.GET("/all", r.middleware.AuthenticateUser, r.GetAllVideos)
	video.GET("/:id", r.GetSpecificVideo)

	podcast := mainRouterGroup.Group("/podcast")
	podcast.POST("/new", r.NewPodcast)
	podcast.GET("/all", r.middleware.AuthenticateUser, r.GetAllPodcastsBasedOnUserStatus)
	podcast.GET("/:id", r.GetPodcastByID)

	journaling := mainRouterGroup.Group("/journaling")
	journaling.POST("/new", r.middleware.AuthenticateUser, r.NewJournalingEntry)
	journaling.GET("/:id", r.GetJournalingEntryByID)

	event := mainRouterGroup.Group("/event")
	event.POST("/new", r.NewEvent)
	event.GET("/:id", r.GetEventByID)
	event.GET("/all", r.GetAllEvents)
	event.POST("/join", r.middleware.AuthenticateUser, r.EventPayment)
}

func (r *Rest) Start() {
	err := r.router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

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
