package mysql

import (
	"log"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func questionSeeder(db *gorm.DB) error {
	questions := []entity.JournalingQuestion{
		{
			Question: "Question 1 Sample",
		},
		{
			Question: "Question 2 Sample",
		},
		{
			Question: "Question 3 Sample",
		},
		{
			Question: "Question 4 Sample",
		},
	}

	err := db.Create(&questions).Error

	if err != nil {
		return err
	}

	return nil
}

func moodSeeder(db *gorm.DB) error {
	moods := []entity.JournalingMood{
		{
			Mood: "Happy",
		},
		{
			Mood: "Sad",
		},
		{
			Mood: "Angry",
		},
		{
			Mood: "Neutral",
		},
	}

	err := db.Create(&moods).Error

	if err != nil {
		return err
	}

	return nil

}

func userSeeder(db *gorm.DB, bcrypt bcrypt.BcryptInterface) error {
	pass, err := bcrypt.HashPassword("syahreza")

	if err != nil {
		return err
	}

	user := []entity.User{
		{
			ID:               uuid.New(),
			Name:             "Syahreza",
			Email:            "superadmin@admin.com",
			Password:         pass,
			IsEmailVerified:  true,
			VerificationCode: "",
			IsSubscribed:     true,
		},
		{
			ID:               uuid.New(),
			Name:             "Also Syahreza",
			Email:            "me@syahreza.com",
			Password:         pass,
			IsEmailVerified:  true,
			VerificationCode: "",
			IsSubscribed:     false,
		},
	}

	err = db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func afirmationWordSeeder(db *gorm.DB) error {
	words := []entity.AfirmationWord{
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 4",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Sad 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 4",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 3",
		},
	}

	err := db.CreateInBatches(&words, len(words)).Error
	if err != nil {
		return err
	}

	return nil
}

func SeedData(db *gorm.DB, bcrypt *bcrypt.BcryptInterface) {
	var totalQuestion int64
	var totalMood int64
	var totalUser int64
	var totalAfirmationWord int64

	err := db.Model(&entity.JournalingQuestion{}).Count(&totalQuestion).Error
	if err != nil {
		log.Fatalf("Error counting question data: %v", err)
	}

	err = db.Model(&entity.JournalingMood{}).Count(&totalMood).Error
	if err != nil {
		log.Fatalf("Error counting mood data: %v", err)
	}

	err = db.Model(&entity.User{}).Count(&totalUser).Error
	if err != nil {
		log.Fatalf("Error counting user data: %v", err)
	}

	err = db.Model(&entity.AfirmationWord{}).Count(&totalAfirmationWord).Error
	if err != nil {
		log.Fatalf("Error counting affirmation word data: %v", err)
	}

	// seed data if there is no data in the table
	if totalQuestion == 0 {
		err := questionSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding question data: %v", err)
		}
	}

	if totalMood == 0 {
		err := moodSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding mood data: %v", err)
		}
	}

	if totalUser == 0 {
		err := userSeeder(db, *bcrypt)

		if err != nil {
			log.Fatalf("Error seeding user data: %v", err)
		}
	}

	if totalAfirmationWord == 0 {
		err := afirmationWordSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding affirmation word data: %v", err)
		}
	}

	log.Println("Question data seeded")
	log.Println("Mood data seeded")
	log.Println("User data seeded")
	log.Println("Afirmation words data seeded")
}
