package middleware

import (
	"errors"
	"strings"

	"github.com/Syahreza-Ferdian/heal-in/internal/service"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

type MiddlewareInterface interface {
	AuthenticateUser(ctx *gin.Context)
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
