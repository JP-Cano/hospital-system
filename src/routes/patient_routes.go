package routes

import (
	"github.com/gin-gonic/gin"
	"hospital-system/src/controllers"
)

func Patient(router *gin.Engine, patientController *controllers.PatientController) {
	patientRoutes := router.Group("/patients")
	{
		patientRoutes.GET("/", patientController.GetPatients)
		patientRoutes.GET("/search", patientController.SearchPatient)
	}
}
