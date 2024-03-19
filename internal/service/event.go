package service

import (
	"fmt"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type InterfaceEventService interface {
	NewEvent(newEvent model.NewEventRequest) (*entity.Event, error)
	GetEvent(id string) (*entity.Event, error)
	GetAllEvents() ([]*entity.Event, error)
	EventImageUploader(eventID string, eventImageReq model.EventImageUploadParam) error
	EventPayment(eventPaymentReq model.EventPaymentRequest) (model.MidtransResponse, error)
	GetUserJoinedEvents(userID string) ([]entity.Event, error)
}

type EventService struct {
	er       repository.InterfaceEventRepository
	eir      repository.InterfaceEventImageRepository
	ur       repository.InterfaceUserRepository
	per      repository.InterfacePaymentEventRepository
	supabase supabase.SupabaseInterface
	snap     snap.Client
}

func NewEventService(er repository.InterfaceEventRepository, eir repository.InterfaceEventImageRepository, supabase supabase.SupabaseInterface, ur repository.InterfaceUserRepository, snap snap.Client, per repository.InterfacePaymentEventRepository) InterfaceEventService {
	return &EventService{
		er:       er,
		eir:      eir,
		ur:       ur,
		supabase: supabase,
		snap:     snap,
		per:      per,
	}
}

func (es *EventService) NewEvent(newEvent model.NewEventRequest) (*entity.Event, error) {
	parseStartDate, err := time.Parse("2006-01-02", newEvent.StartDate)
	if err != nil {
		return nil, err
	}
	parseEndDate, err := time.Parse("2006-01-02", newEvent.EndDate)
	if err != nil {
		return nil, err
	}

	eventEntity := &entity.Event{
		ID:               uuid.New(),
		Title:            newEvent.Title,
		Body:             newEvent.Body,
		StartDate:        parseStartDate,
		EndDate:          parseEndDate,
		Location:         newEvent.Location,
		IsRequirePayment: newEvent.IsRequirePayment,
		PaymentAmount:    newEvent.PaymentAmount,
	}

	event, err := es.er.CreateNewEvent(eventEntity)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (es *EventService) GetEvent(id string) (*entity.Event, error) {
	event, err := es.er.GetEventByID(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (es *EventService) GetAllEvents() ([]*entity.Event, error) {
	events, err := es.er.GetAllEvents()

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (es *EventService) EventImageUploader(eventID string, eventImageReq model.EventImageUploadParam) error {
	event, err := es.er.GetEventByID(eventID)
	if err != nil {
		return err
	}

	link, err := es.supabase.UploadFile(eventImageReq.Image)
	if err != nil {
		return err
	}

	eventImageEntity := &entity.EventImage{
		ID:        uuid.New(),
		EventID:   event.ID,
		ImageLink: link,
	}

	_, err = es.eir.CreateEventImage(eventImageEntity)

	if err != nil {
		err1 := es.supabase.DeleteFile(link)

		if err1 != nil {
			return err
		}
		return err
	}

	return nil
}

func (es *EventService) EventPayment(eventPaymentReq model.EventPaymentRequest) (model.MidtransResponse, error) {
	event, err := es.er.GetEventByID(eventPaymentReq.EventID.String())
	if err != nil {
		return model.MidtransResponse{}, err
	}

	if !event.IsRequirePayment {
		return model.MidtransResponse{}, fmt.Errorf("this event doesn't require payment")
	}

	user, err := es.ur.GetUser(&model.GetUserParam{
		ID: eventPaymentReq.UserID,
	})
	if err != nil {
		return model.MidtransResponse{}, err
	}

	orderID := uuid.New()

	snapReq := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID.String(),
			GrossAmt: int64(event.PaymentAmount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				Qty:   1,
				Price: int64(event.PaymentAmount),
				Name:  fmt.Sprintf("%s payment for user %s", event.Title, user.Name),
			},
		},
		CustomField1: "event payment",
	}

	response, err2 := es.snap.CreateTransaction(&snapReq)
	if err2 != nil {
		return model.MidtransResponse{}, err
	}

	_, err = es.per.CreateNewPayment(&entity.PaymentEvent{
		ID:          orderID,
		EventID:     event.ID,
		UserID:      user.ID,
		Amount:      event.PaymentAmount,
		Description: fmt.Sprintf("%s payment for user %s", event.Title, user.Name),
		IsCompleted: false,
	})
	if err != nil {
		return model.MidtransResponse{}, err
	}

	midtransResponse := model.MidtransResponse{
		Token:   response.Token,
		SnapURL: response.RedirectURL,
	}

	return midtransResponse, nil
}

func (es *EventService) GetUserJoinedEvents(userID string) ([]entity.Event, error) {
	events, err := es.ur.GetUserJoinedEvents(userID)

	if err != nil {
		return nil, err
	}

	return events, nil
}
