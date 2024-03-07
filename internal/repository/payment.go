package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InterfacePaymentRepository interface {
	CreatePayment(payment *entity.Payment) (*entity.Payment, error)
	UpdatePaymentOnSuccess(orderID string) error
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) InterfacePaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (pr *PaymentRepository) CreatePayment(payment *entity.Payment) (*entity.Payment, error) {
	err := pr.db.Debug().Create(&payment).Error
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (pr *PaymentRepository) UpdatePaymentOnSuccess(orderID string) error {
	err := pr.db.Debug().Model(&entity.Payment{}).Where("id = ?", orderID).Update("is_completed", 1).Error

	if err != nil {
		return err
	}

	var userID uuid.UUID
	err = pr.db.Debug().Model(&entity.Payment{}).Where("id = ?", orderID).Select("user_id").Row().Scan(&userID)

	if err != nil {
		return err
	}

	err = pr.db.Debug().Model(&entity.User{}).Where("id = ?", userID).Update("is_subscribed", 1).Error

	if err != nil {
		return err
	}

	return nil
}
