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
	"gorm.io/gorm"

	bcrypt_import "golang.org/x/crypto/bcrypt"
)

func (r *Rest) CreateUser(c *gin.Context) {
	var userRequest model.UserRegister

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		// form validation
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	// user service call
	newUser, emailData, err := r.service.UserService.Register(userRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			// error handling: email sudah terdaftar tapi dipake buat register lagi
			response.OnFailed(c, http.StatusConflict, "Email sudah terdaftar", fmt.Errorf("email '%s' sudah terdaftar di sistem", userRequest.Email))
			return
		}
		response.OnFailed(c, http.StatusInternalServerError, "Gagal registrasi user", err)
		return
	}

	err = r.mail.SendEmail(&userRequest, &emailData)

	if err != nil {
		response.OnFailed(c, http.StatusInternalServerError, "Gagal mengirim email verifikasi, silakan coba register kembali", err)
		r.service.UserService.DeleteUser(&newUser)
		return
	}

	response.OnSuccess(c, http.StatusCreated, "Berhasil registrasi user, silakan cek email untuk verifikasi", nil)
}

func (r *Rest) Login(c *gin.Context) {
	var userRequest model.UserLogin

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			// error handling: validasi input, e.g: email tidak valid
			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	token, err := r.service.UserService.Login(userRequest)

	if err != nil {
		// error handling: password salah
		if errors.Is(err, bcrypt_import.ErrMismatchedHashAndPassword) {
			response.OnFailed(c, http.StatusBadRequest, "Gagal login", fmt.Errorf("password yang anda masukkan salah"))
			return
		}

		// error handling: email tidak terdaftar
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OnFailed(c, http.StatusBadRequest, "Gagal login", fmt.Errorf("email yang anda masukkan tidak terdaftar"))
			return

		}

		// another error
		response.OnFailed(c, http.StatusBadRequest, "Gagal login", err)
		return
	}

	response.OnSuccess(c, http.StatusOK, "Berhasil login ke sistem", token)
}

func (r *Rest) VerifyEmail(c *gin.Context) {
	verificationCode := c.Param("verificationCode")

	err := r.service.UserService.Verify(verificationCode)

	if err != nil {
		response.OnFailed(c, http.StatusBadRequest, "Gagal verifikasi email", err)
		return
	}

	response.OnSuccess(c, http.StatusOK, "Berhasil verifikasi email", nil)
}
