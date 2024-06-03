package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hospital-system/src/entities"
	"hospital-system/src/services"
	"hospital-system/src/utils"
	"net/http"
	"time"
)

type MedicalAppointmentController struct {
	repository services.MedicalAppointmentRepository
}

func NewMedicalAppointmentController(repository services.MedicalAppointmentRepository) *MedicalAppointmentController {
	return &MedicalAppointmentController{repository: repository}
}

func (m *MedicalAppointmentController) CreateAppointment(c *gin.Context) {
	var appointmentDto entities.MedicalAppointmentDto
	if err := c.ShouldBindJSON(&appointmentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment, err := toMedicalAppointment(appointmentDto)
	utils.HandleError(c, err, http.StatusBadRequest)

	if err := m.repository.CreateAppointment(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointmentDto)
	return
}

func (m *MedicalAppointmentController) GetDoctorAppointments(c *gin.Context) {
	doctorId := uuid.MustParse(c.Param("id"))
	appointments, err := m.repository.GetDoctorAppointments(doctorId)
	utils.HandleError(c, err, http.StatusInternalServerError)
	c.JSON(http.StatusOK, appointments)
	return
}

func toMedicalAppointment(appointmentDto entities.MedicalAppointmentDto) (entities.MedicalAppointment, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", appointmentDto.DateTime)
	if err != nil {
		return entities.MedicalAppointment{}, err
	}

	return entities.MedicalAppointment{
		Id:        uuid.New(),
		PatientId: appointmentDto.PatientId,
		DoctorId:  appointmentDto.DoctorId,
		DateTime:  parsedTime,
		Reason:    appointmentDto.Reason,
	}, nil
}
