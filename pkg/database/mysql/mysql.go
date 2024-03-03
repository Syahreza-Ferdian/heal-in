package mysql

import (
	"log"

	"github.com/Syahreza-Ferdian/heal-in/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDb() *gorm.DB {
	dsn := config.LoadConfigDatabase()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	} else {
		log.Println("Successfully connected to database")
		return db
	}
}
