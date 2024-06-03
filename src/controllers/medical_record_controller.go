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

type MedicalRecordController struct {
	repository services.MedicalRecordRepository
}

func NewMedicalRecordController(repository services.MedicalRecordRepository) *MedicalRecordController {
	return &MedicalRecordController{repository: repository}
}

func (m *MedicalRecordController) UpdateMedicalRecord(c *gin.Context) {
	patientId := uuid.MustParse(c.Param("id"))

	var updateMedicalRecord entities.UpdateMedicalRecord
	if err := c.ShouldBindJSON(&updateMedicalRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.repository.UpdateMedicalRecord(patientId, updateMedicalRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medical record updated successfully"})
}

func (m *MedicalRecordController) GetDoctorReport(c *gin.Context) {
	doctorId := uuid.MustParse(c.Param("id"))

	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	utils.HandleError(c, err, http.StatusBadRequest)
	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	utils.HandleError(c, err, http.StatusBadRequest)
	report, err := m.repository.GetDoctorReport(doctorId, startDate, endDate)
	utils.HandleError(c, err, http.StatusInternalServerError)
	c.JSON(http.StatusOK, report)
}
