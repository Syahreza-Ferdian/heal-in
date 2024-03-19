package model

import (
	"mime/multipart"
)

type EventImageUploadParam struct {
	Image *multipart.FileHeader `form:"image"`
}

type NewEventRequest struct {
	Title            string `form:"title" binding:"required"`
	Body             string `form:"body" binding:"required"`
	StartDate        string `form:"start_date" binding:"required"`
	EndDate          string `form:"end_date" binding:"required"`
	Location         string `form:"location" binding:"required"`
	IsRequirePayment bool   `form:"is_require_payment" binding:"required"`
	PaymentAmount    int    `form:"payment_amount"`
}
