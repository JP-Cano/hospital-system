package routes

import (
	"github.com/gin-gonic/gin"
	"hospital-system/src/controllers"
)

func MedicalAppointment(router *gin.Engine, medicalAppointmentController *controllers.MedicalAppointmentController) {
	medicalAppointmentRoutes := router.Group("/medical-appointments")
	{
		medicalAppointmentRoutes.GET("/doctor/:id", medicalAppointmentController.GetDoctorAppointments)
		medicalAppointmentRoutes.POST("/", medicalAppointmentController.CreateAppointment)
	}
}
