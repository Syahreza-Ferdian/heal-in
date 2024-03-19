package service

import (
	"fmt"
	"log"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/scheduler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type InterfacePaymentService interface {
	NewPayment(paymentReq model.MidtransRequest, c *gin.Context) (model.MidtransResponse, error)
	// PaymentNotification(payload model.NotificationPayload) error
	SendEmailToExpiredSubs() ([]email.EmailDataExpSubs, error)
	SetupScheduler()
}

type PaymentService struct {
	pr         repository.InterfacePaymentRepository
	ur         repository.InterfaceUserRepository
	snapClient snap.Client
	coreApi    coreapi.Client
	scheduler  scheduler.SchedulerInterface
	mail       email.EmailService
}

func NewPaymentService(pr repository.InterfacePaymentRepository, snapClient snap.Client, coreApi coreapi.Client, ur repository.InterfaceUserRepository, sch scheduler.SchedulerInterface, mail email.EmailService) InterfacePaymentService {
	return &PaymentService{
		pr:         pr,
		ur:         ur,
		snapClient: snapClient,
		coreApi:    coreApi,
		scheduler:  sch,
		mail:       mail,
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
		CustomField1: "subscription payment",
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

// func (ps *PaymentService) PaymentNotification(payload model.NotificationPayload) error {
// 	orderID, oke := payload["order_id"].(string)

// 	if !oke {
// 		return fmt.Errorf("failed to get order_id")
// 	}

// 	transactionResponse, err := ps.coreApi.CheckTransaction(orderID)

// 	if err != nil {
// 		return err
// 	}

// 	if transactionResponse != nil {
// 		if transactionResponse.TransactionStatus == "settlement" {
// 			err := ps.pr.UpdatePaymentOnSuccess(orderID)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

func (ps *PaymentService) SendEmailToExpiredSubs() ([]email.EmailDataExpSubs, error) {
	expPayments, err := ps.pr.GetExpiredSubscriptions()
	if err != nil {
		return nil, err
	}

	var emailData []email.EmailDataExpSubs

	for _, payment := range expPayments {
		user, err := ps.ur.GetUser(&model.GetUserParam{
			ID: payment.UserID,
		})
		if err != nil {
			return nil, err
		}

		err = ps.pr.UpdatePaymentOnExpired(payment.ID.String())
		if err != nil {
			return nil, err
		}

		newData := &email.EmailDataExpSubs{
			FirstName: user.Name,
			Subject:   "Pemberitahuan Langganan Anda Telah Berakhir",
			ToEmail:   user.Email,
		}

		emailData = append(emailData, *newData)
	}

	return emailData, nil
}

func (ps *PaymentService) SetupScheduler() {
	ps.scheduler.Stop()

	// corn job pattern : second, minute, hour, day, month, weekday

	// Run every day at 00:00:00
	ps.scheduler.AddFunction("0 0 0 * * *", ps.ExpEmailSender)

	// Testing buat memastikan scheduler jalan dengan baik
	// ps.scheduler.AddFunction("@every 5m", func() {
	// 	log.Println("Scheduler: If you seen this message, it means the scheduler is running properly.")
	// })

	ps.scheduler.Start()
}

func (ps *PaymentService) ExpEmailSender() {
	emailData, err := ps.SendEmailToExpiredSubs()
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	for _, data := range emailData {
		err := ps.mail.SendExpirationSubsEmail(&data)
		if err != nil {
			log.Fatalf("Error while sending email : %v", err)
		}
	}
}
