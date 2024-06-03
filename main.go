package main

import (
	"fmt"
	"hospital-system/src/config"
	"hospital-system/src/database"
	"hospital-system/src/server"
	"log"
	"os"
)

func main() {
	dsn := config.DatabaseUrlBuilder()
	db, err := database.New(dsn)
	defer database.Close(db)
	database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := server.New(fmt.Sprintf(":%s", port), db)
	s.Serve()
}
