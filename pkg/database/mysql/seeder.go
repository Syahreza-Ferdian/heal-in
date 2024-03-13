package mysql

import (
	"log"

	"github.com/Syahreza-Ferdian/heal-in/entity"
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

func SeedData(db *gorm.DB) {
	var totalQuestion int64

	err := db.Model(&entity.JournalingQuestion{}).Count(&totalQuestion).Error

	if err != nil {
		log.Fatalf("Error counting question data: %v", err)
	}

	if totalQuestion == 0 {
		err := questionSeeder(db)
		
		if err != nil {
			log.Fatalf("Error seeding question data: %v", err)
		}
	}


	if err != nil {
		log.Fatalf("Error seeding question data: %v", err)
	}

	log.Println("Question data seeded")
}
