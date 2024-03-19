package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/email"
	"github.com/Syahreza-Ferdian/heal-in/pkg/jwt"
	"github.com/Syahreza-Ferdian/heal-in/pkg/scheduler"
	"github.com/Syahreza-Ferdian/heal-in/pkg/supabase"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type Service struct {
	UserService           InterfaceUserService
	PaymentService        InterfacePaymentService
	ArtikelService        InterfaceArtikelService
	VideoService          InterfaceVideoService
	PodcastService        InterfacePodcastService
	JournalingService     InterfaceJournalingService
	AfirmationWordService InterfaceAfirmationWordService
	EventService          InterfaceEventService
}

type InitService struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.BcryptInterface
	JwtAuth    jwt.JWTInterface
	SnapClient snap.Client
	CoreApi    coreapi.Client
	Supabase   supabase.SupabaseInterface
	Scheduler  scheduler.SchedulerInterface
	Mail       email.EmailService
}

func NewService(param InitService) *Service {
	return &Service{
		UserService:           NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth),
		PaymentService:        NewPaymentService(param.Repository.PaymentRepository, param.SnapClient, param.CoreApi, param.Repository.UserRepository, param.Scheduler, param.Mail),
		ArtikelService:        NewArtikelService(param.Repository.ArtikelRepository, param.Supabase, param.Repository.ArtikelImageRepository, param.Repository.UserRepository),
		VideoService:          NewVideoService(param.Repository.VideoRepository, param.Supabase, param.Repository.UserRepository),
		PodcastService:        NewPodcastService(param.Repository.PodcastRepository, param.Supabase, param.Repository.UserRepository),
		JournalingService:     NewJournalingService(param.Repository.JournalingAnsRepository, param.Repository.JournalingEntryRepository, param.Repository.JournalingQuestionRepository, param.Repository.UserRepository),
		AfirmationWordService: NewAfirmationWordService(param.Repository.AfirmationWordRepository),
		EventService:          NewEventService(param.Repository.EventRepository, param.Repository.EventImageRepository, param.Supabase),
	}
}
