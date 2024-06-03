package controllers

import (
	"github.com/gin-gonic/gin"
	"hospital-system/src/services"
	"hospital-system/src/utils"
	"log"
	"net/http"
)

type PatientController struct {
	repository services.PatientRepository
}

func NewPatientController(repository services.PatientRepository) *PatientController {
	return &PatientController{repository: repository}
}

func (s *PatientController) GetPatients(c *gin.Context) {
	patients, err := s.repository.GetPatients()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
	return
}

func (s *PatientController) SearchPatient(c *gin.Context) {
	name := c.Query("name")
	phone := c.Query("phone")
	email := c.Query("email")

	patient, err := s.repository.SearchPatients(name, phone, email)
	utils.HandleError(c, err, http.StatusInternalServerError)
	c.JSON(http.StatusOK, patient)
	return
}
