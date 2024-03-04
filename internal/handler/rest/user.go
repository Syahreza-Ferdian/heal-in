package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/validate"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (r *Rest) CreateUser(c *gin.Context) {
	var userRequest model.UserRegister

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	// user service call
	err = r.service.UserService.Register(userRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			response.OnFailed(c, http.StatusConflict, "Email sudah terdaftar", fmt.Errorf("email '%s' sudah terdaftar di sistem", userRequest.Email))
			return
		}
		response.OnFailed(c, http.StatusInternalServerError, "Gagal registrasi user", err)
		return
	}

	response.OnSuccess(c, http.StatusCreated, "Berhasil registrasi user", nil)
}
