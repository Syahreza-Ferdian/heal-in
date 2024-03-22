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
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	newUser, emailData, err := r.service.UserService.Register(userRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			response.OnFailed(c, http.StatusConflict, "Email sudah terdaftar", fmt.Errorf("email '%s' sudah terdaftar di sistem", userRequest.Email))
			return
		}
		response.OnFailed(c, http.StatusInternalServerError, "Gagal registrasi user", err)
		return
	}

	err = r.mail.SendVerificationEmail(&userRequest, &emailData)

	if err != nil {
		response.OnFailed(c, http.StatusInternalServerError, "Gagal mengirim email verifikasi, silakan coba register kembali", err)
		r.service.UserService.DeleteUser(&newUser)
		return
	}

	response.OnSuccess(c, http.StatusCreated, "Berhasil registrasi user, silakan cek email untuk verifikasi. NOTE: Apabila Anda tidak menerima email, silakan pastikan kembali bahwa email yang Anda masukkan sudah benar.", nil)
}

func (r *Rest) Login(c *gin.Context) {
	var userRequest model.UserLogin

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	token, err := r.service.UserService.Login(userRequest)

	if err != nil {
		if errors.Is(err, bcrypt_import.ErrMismatchedHashAndPassword) {
			response.OnFailed(c, http.StatusBadRequest, "Gagal login", fmt.Errorf("password yang anda masukkan salah"))
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OnFailed(c, http.StatusBadRequest, "Gagal login", fmt.Errorf("email yang anda masukkan tidak terdaftar"))
			return

		}

		response.OnFailed(c, http.StatusBadRequest, "Gagal login", err)
		return
	}

	response.OnSuccess(c, http.StatusOK, "Berhasil login ke sistem", token)
}

func (r *Rest) VerifyEmail(c *gin.Context) {
	verificationCode := c.Param("verificationCode")

	err := r.service.UserService.Verify(verificationCode)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OnFailed(c, http.StatusBadRequest, "Gagal verifikasi email", fmt.Errorf("kode verifikasi tidak valid"))
			return
		}

		response.OnFailed(c, http.StatusBadRequest, "Gagal verifikasi email", err)
		return
	}

	response.OnSuccess(c, http.StatusOK, "Berhasil verifikasi email", nil)
}
