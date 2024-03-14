package mysql

import (
	"log"

	"github.com/Syahreza-Ferdian/heal-in/entity"
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

func userSeeder(db *gorm.DB) error {
	user := []entity.User{
		{
			ID: uuid.New(),
			Name: "Syahreza",
			Email: "superadmin@admin.com",
			Password: "syahreza",
			IsEmailVerified: true,
			VerificationCode: "",
			IsSubscribed: true,
		},
		{
			ID: uuid.New(),
			Name: "Also Syahreza",
			Email: "me@syahreza.com",
			Password: "syahreza",
			IsEmailVerified: true,
			VerificationCode: "",
			IsSubscribed: false,
		},
	}

	err := db.Create(&user).Error
	
	if err != nil {
		return err
	}

	return nil
}

func SeedData(db *gorm.DB) {
	var totalQuestion int64
	var totalMood int64
	var totalUser int64

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
		err := userSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding user data: %v", err)
		}
	}

	log.Println("Question data seeded")
	log.Println("Mood data seeded")
	log.Println("User data seeded")
}
