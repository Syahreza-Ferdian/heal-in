package service

import (
	"fmt"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type InterfacePaymentService interface {
	NewPayment(paymentReq model.MidtransRequest, c *gin.Context) (model.MidtransResponse, error)
	PaymentNotification(payload model.NotificationPayload) error
}

type PaymentService struct {
	pr         repository.InterfacePaymentRepository
	snapClient snap.Client
	coreApi    coreapi.Client
}

func NewPaymentService(pr repository.InterfacePaymentRepository, snapClient snap.Client, coreApi coreapi.Client) InterfacePaymentService {
	return &PaymentService{
		pr:         pr,
		snapClient: snapClient,
		coreApi:    coreApi,
	}
}

func (ps *PaymentService) NewPayment(paymentReq model.MidtransRequest, c *gin.Context) (model.MidtransResponse, error) {
	currentUser, oke := c.Get("user")

	if !oke {
		return model.MidtransResponse{}, fmt.Errorf("failed to get current user")
	}

	paymentReq.OrderId = uuid.New()
	paymentReq.UserID = currentUser.(*entity.User).ID

	// var snapClient = snap.Client{}
	// snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	snapReq := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentReq.OrderId.String(),
			GrossAmt: int64(paymentReq.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: currentUser.(*entity.User).Name,
			Email: currentUser.(*entity.User).Email,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				Qty:   1,
				Price: int64(paymentReq.Amount),
				Name:  paymentReq.Description,
			},
		},
	}

	response, err := ps.snapClient.CreateTransaction(&snapReq)

	if err != nil {
		return model.MidtransResponse{}, err
	}

	ps.pr.CreatePayment(&entity.Payment{
		ID:          paymentReq.OrderId,
		UserID:      currentUser.(*entity.User).ID,
		Amount:      paymentReq.Amount,
		Description: paymentReq.Description,
		ExpiredAt:   time.Now().Add(3 * 30 * 24 * time.Hour),
		IsCompleted: false,
	})

	midtransResponse := model.MidtransResponse{
		Token:   response.Token,
		SnapURL: response.RedirectURL,
	}

	return midtransResponse, nil
}

func (ps *PaymentService) PaymentNotification(payload model.NotificationPayload) error {
	orderID, oke := payload["order_id"].(string)

	if !oke {
		return fmt.Errorf("failed to get order_id")
	}

	transactionResponse, err := ps.coreApi.CheckTransaction(orderID)

	if err != nil {
		return err
	}

	if transactionResponse != nil {
		if transactionResponse.TransactionStatus == "settlement" {
			err := ps.pr.UpdatePaymentOnSuccess(orderID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
