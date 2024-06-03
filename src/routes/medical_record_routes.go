package routes

import (
	"github.com/gin-gonic/gin"
	"hospital-system/src/controllers"
)

func MedicalRecord(router *gin.Engine, medicalRecordController *controllers.MedicalRecordController) {
	medicalRecordRoutes := router.Group("/medical-record")
	{
		medicalRecordRoutes.GET("/doctor/:id/report", medicalRecordController.GetDoctorReport)
		medicalRecordRoutes.PATCH("/patient/:id", medicalRecordController.UpdateMedicalRecord)
	}
}
