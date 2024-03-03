package rest

import (
	"fmt"
	"log"
	"os"

	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service service.Service
}

func NewRest(router *gin.Engine, service service.Service) *Rest {
	return &Rest{
		router:  router,
		service: service,
	}
}

func (r *Rest) EndPoint() {
	mainRouterGroup := r.router.Group("/api")

	user := mainRouterGroup.Group("/user")
	user.POST("/register", r.CreateUser)
	// user.POST("/login", r.Login)
}

func (r *Rest) Start() {
	err := r.router.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_ADDRESS"), os.Getenv("APP_PORT")))

	if err != nil {
		log.Fatalf("Cannot running the server. Error: %v", err)
	}
}
