package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

type MiddlewareInterface interface {
	AuthenticateUser(ctx *gin.Context)
	Cors(c *gin.Context)
}

type Middleware struct {
	jwtAuth jwt.JWTInterface
	service *service.Service
}

func Init(jwtAuth jwt.JWTInterface, service *service.Service) MiddlewareInterface {
	return &Middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}

func (m *Middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		response.OnFailed(ctx, 401, "Unauthorized", errors.New("tidak ada token yang diberikan"))
		return
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtAuth.VerifyJWTToken(token)

	if err != nil {
		response.OnFailed(ctx, 401, "Unauthorized", errors.New("token invalid"))
		return
	}

	user, err := m.service.UserService.GetUser(&model.GetUserParam{
		ID: userId,
	})

	if err != nil {
		response.OnFailed(ctx, 401, "Unauthorized", errors.New("failed to get user"))
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}

func (m *Middleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Content-Disposition")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Max-Age", "12h")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
