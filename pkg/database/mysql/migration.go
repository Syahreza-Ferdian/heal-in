package mysql

import (
	"log"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().DropTable(
	// &entity.User{},
	// &entity.Artikel{},
	// &entity.Payment{},
	// &entity.ArtikelImage{},
	// &entity.Video{},
	// &entity.Podcast{},
	// &entity.JournalingEntry{},
	// &entity.JournalingQuestion{},
	// &entity.JournalingAnswer{},
	// &entity.JournalingMood{},
	// &entity.AfirmationWord{},
	// &entity.Event{},
	// &entity.EventImage{},
	)

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Payment{},
		&entity.Artikel{},
		&entity.ArtikelImage{},
		&entity.Video{},
		&entity.Podcast{},
		&entity.JournalingEntry{},
		&entity.JournalingQuestion{},
		&entity.JournalingAnswer{},
		&entity.JournalingMood{},
		&entity.AfirmationWord{},
		&entity.Event{},
		&entity.EventImage{},
	)

	if err != nil {
		log.Fatalf("There is an error while migrating database. Error: %v", err)
	} else {
		log.Println("Successfully migrating database")
	}
}
