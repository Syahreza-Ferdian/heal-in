package rest

import (
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateUser(c *gin.Context) {
	var userRequest model.UserRegister

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		response.OnFailed(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	// user service call
	err = r.service.UserService.Register(userRequest)
	if err != nil {
		response.OnFailed(c, http.StatusInternalServerError, "Gagal registrasi user", err)
		return
	}

	response.OnSuccess(c, http.StatusCreated, "Berhasil registrasi user", "Data is secret, you can not see it!")
}
