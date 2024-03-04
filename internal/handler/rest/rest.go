package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/pkg/middleware"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.MiddlewareInterface
}

func NewRest(router *gin.Engine, service *service.Service, middleware middleware.MiddlewareInterface) *Rest {
	return &Rest{
		router:     router,
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) EndPoint() {
	mainRouterGroup := r.router.Group("/api")

	user := mainRouterGroup.Group("/user")
	user.POST("/register", r.CreateUser)
	user.POST("/login", r.Login)
	user.GET("/login-user", r.middleware.AuthenticateUser, Testing)
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
