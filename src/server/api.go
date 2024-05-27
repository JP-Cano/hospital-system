package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hospital-system/src/routes"
	"log"
)

type Api struct {
	Port string
	DB   *gorm.DB
}

func New(port string, db *gorm.DB) *Api {
	return &Api{
		Port: port,
		DB:   db,
	}
}

func (api *Api) Serve() {
	router := gin.Default()

	registerRoutes(router)

	err := router.Run(api.Port)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server running on port: %s", api.Port)
}

func registerRoutes(router *gin.Engine) {
	routes.Patient(router)
}
