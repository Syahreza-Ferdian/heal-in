package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository         InterfaceUserRepository
	PaymentRepository      InterfacePaymentRepository
	ArtikelRepository      InterfaceArtikelRepository
	ArtikelImageRepository InterfaceArtikelImageRepository
	VideoRepository        InterfaceVideoRepository
	PodcastRepository      InterfacePodcastRepository

	// journaling related repositories
	JournalingAnsRepository      InterfaceJournalingAnsRepository
	JournalingEntryRepository    InterfaceJournalingEntryRepository
	JournalingQuestionRepository InterfaceJournalingQuestionRepository
	AfirmationWordRepository     InterfaceAfirmationWordRepository

	// event related repositories
	EventRepository      InterfaceEventRepository
	EventImageRepository InterfaceEventImageRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:               NewUserRepository(db),
		PaymentRepository:            NewPaymentRepository(db),
		ArtikelRepository:            NewArtikelRepository(db),
		ArtikelImageRepository:       NewArtikelImageRepository(db),
		VideoRepository:              NewVideoRepository(db),
		JournalingAnsRepository:      NewJournalingAnsRepository(db),
		JournalingEntryRepository:    NewJournalingEntryRepository(db),
		JournalingQuestionRepository: NewJournalingQuestionRepository(db),
		AfirmationWordRepository:     NewAfirmationWordRepository(db),
		PodcastRepository:            NewPodcastRepository(db),
		EventRepository:              NewEventRepository(db),
		EventImageRepository:         NewEventImageRepository(db),
	}
}
