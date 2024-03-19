package rest

import (
	"errors"
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/validate"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (r *Rest) NewPayment(c *gin.Context) {
	var paymentReq model.MidtransRequest

	err := c.ShouldBindJSON(&paymentReq)

	if err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			err := validate.GetErrors(ve)

			response.OnErrorValidate(c, http.StatusBadRequest, "Validation error", err)
			return
		}
	}

	newPaymentResponse, err := r.service.PaymentService.NewPayment(paymentReq, c)

	if err != nil {
		response.OnFailed(c, http.StatusInternalServerError, "Gagal membuat payment", err)
		return
	}

	response.OnSuccess(c, http.StatusCreated, "Berhasil membuat payment", newPaymentResponse)
}

func (r *Rest) PaymentNotification(c *gin.Context) {
	var payload model.NotificationPayload

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		response.OnFailed(c, http.StatusBadRequest, "Gagal bind payload", err)
		return
	}

	err = r.service.MidtransService.HandleAfterPayment(payload)

	if err != nil {
		response.OnFailed(c, http.StatusInternalServerError, "Payment Failed", err)
		return
	}

	response.OnSuccess(c, http.StatusOK, "Payment Success", nil)
}

func (r *Rest) InitPaymentExpScheduler() {
	r.service.PaymentService.SetupScheduler()
}
