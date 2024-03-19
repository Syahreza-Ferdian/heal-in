package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title            string         `json:"title" gorm:"type:varchar(255);not null;"`
	Body             string         `json:"body" gorm:"type:text;not null;"`
	PostedAt         time.Time      `json:"posted_at" gorm:"autoCreateTime"`
	StartDate        time.Time      `json:"start_date" gorm:"type:datetime;not null;"`
	EndDate          time.Time      `json:"end_date" gorm:"type:datetime;not null;"`
	Location         string         `json:"location" gorm:"type:varchar(255);not null;"`
	IsRequirePayment bool           `json:"is_require_payment" gorm:"type:bool;not null;"`
	PaymentAmount    int            `json:"payment_amount" gorm:"type:int;not null;"`
	EventImage       []EventImage   `json:"event_images" gorm:"foreignKey:EventID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentEvent     []PaymentEvent `json:"-" gorm:"foreignKey:EventID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
