package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}
	log.Println("Database connected successfully")
	return db, nil
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Printf("Error retrieving database connection: %v", err)
		return
	}
	err = conn.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
		return
	}
	log.Println("Database connection closed successfully")
	return
}
