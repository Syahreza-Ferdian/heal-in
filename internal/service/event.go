package service

import (
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/google/uuid"
)

type InterfaceEventService interface {
	NewEvent(newEvent model.NewEventRequest) (*entity.Event, error)
	GetEvent(id string) (*entity.Event, error)
	GetAllEvents() ([]*entity.Event, error)
	EventImageUploader(eventID string, eventImageReq model.EventImageUploadParam) error
}

type EventService struct {
	er       repository.InterfaceEventRepository
	eir      repository.InterfaceEventImageRepository
	supabase supabase.SupabaseInterface
}

func NewEventService(er repository.InterfaceEventRepository, eir repository.InterfaceEventImageRepository, supabase supabase.SupabaseInterface) InterfaceEventService {
	return &EventService{
		er:       er,
		eir:      eir,
		supabase: supabase,
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
