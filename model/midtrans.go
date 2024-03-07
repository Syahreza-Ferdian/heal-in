package model

import "github.com/google/uuid"

type MidtransRequest struct {
	OrderId     uuid.UUID `json:"order_id"`
	UserID      uuid.UUID `json:"user_id"`
	Amount      int       `json:"amount" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

type MidtransResponse struct {
	Token   string `json:"token"`
	SnapURL string `json:"snap_url"`
}

type NotificationPayload map[string]interface{}
