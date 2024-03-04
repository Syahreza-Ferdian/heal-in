package jwt

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTInterface interface {
	GenerateJWTToken(userId uuid.UUID) (string, error)
	VerifyJWTToken(tokenString string) (uuid.UUID, error)
	GetCurrentLoginUser(ctx *gin.Context) (entity.User, error)
}

type JsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func Init() JWTInterface {
	secretKey := os.Getenv("JWT_SECRET")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	if err != nil {
		log.Fatalf("Error when parsing JWT_EXPIRED_TIME: %v", err)
	}

	return &JsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expiredTime) * time.Minute,
	}
}

func (j *JsonWebToken) GenerateJWTToken(userId uuid.UUID) (string, error) {
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
			Issuer:    "heal-in",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func (j *JsonWebToken) VerifyJWTToken(tokenString string) (uuid.UUID, error) {
	var (
		claims Claims
		userId uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return userId, err
	}

	if !token.Valid {
		return userId, err
	}

	userId = claims.UserId

	return userId, nil
}

func (j *JsonWebToken) GetCurrentLoginUser(ctx *gin.Context) (entity.User, error) {
	user, ok := ctx.Get("user")

	if !ok {
		return entity.User{}, errors.New("user not found")
	}

	return user.(entity.User), nil
}