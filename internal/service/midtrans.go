package service

import (
	"fmt"

	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/midtrans/midtrans-go/coreapi"
)

type InterfaceMidtransService interface {
	HandleAfterPayment(payload model.NotificationPayload) error
}

type MidtransService struct {
	pr      repository.InterfacePaymentRepository
	per     repository.InterfacePaymentEventRepository
	coreapi coreapi.Client
}

func NewMidtransService(coreapi coreapi.Client, pr repository.InterfacePaymentRepository, per repository.InterfacePaymentEventRepository) InterfaceMidtransService {
	return &MidtransService{
		coreapi: coreapi,
		pr:      pr,
		per:     per,
	}
}

func (ms *MidtransService) HandleAfterPayment(payload model.NotificationPayload) error {
	orderID, ada := payload["order_id"].(string)

	if !ada {
		return fmt.Errorf("order_id not found")
	}

	transactionResponse, err := ms.coreapi.CheckTransaction(orderID)

	if err != nil {
		return err
	}

	if transactionResponse != nil {
		if transactionResponse.TransactionStatus == "settlement" {
			if transactionResponse.CustomField1 == "subscription payment" {
				err := ms.pr.UpdatePaymentOnSuccess(orderID)
				if err != nil {
					return err
				}
			} else if transactionResponse.CustomField1 == "event payment" {
				err := ms.per.UpdatePaymentOnSuccess(orderID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
