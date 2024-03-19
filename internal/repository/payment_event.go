package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

type InterfacePaymentEventRepository interface {
	CreateNewPayment(payment *entity.PaymentEvent) (*entity.PaymentEvent, error)
	UpdatePaymentOnSuccess(orderID string) error
}

type PaymentEventRepository struct {
	db *gorm.DB
}

func NewPaymentEventRepository(db *gorm.DB) InterfacePaymentEventRepository {
	return &PaymentEventRepository{
		db: db,
	}
}

func (per *PaymentEventRepository) CreateNewPayment(payment *entity.PaymentEvent) (*entity.PaymentEvent, error) {
	err := per.db.Debug().Create(&payment).Error
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (per *PaymentEventRepository) UpdatePaymentOnSuccess(orderID string) error {
	err := per.db.Debug().Model(&entity.PaymentEvent{}).Where("id = ?", orderID).Update("is_completed", 1).Error

	if err != nil {
		return err
	}

	return nil
}