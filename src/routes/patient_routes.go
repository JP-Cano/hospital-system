package routes

import "github.com/gin-gonic/gin"

func Patient(router *gin.Engine) {
	patientRoutes := router.Group("/patient")
	{
		patientRoutes.GET("/")
		patientRoutes.POST("/")
		patientRoutes.GET("/:id")
	}
}
