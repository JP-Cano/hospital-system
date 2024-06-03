package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hospital-system/src/controllers"
	"hospital-system/src/routes"
	"hospital-system/src/services"
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

	register(router, api.DB)

	err := router.Run(api.Port)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server running on port: %s", api.Port)
}

func register(router *gin.Engine, db *gorm.DB) {
	patientRepository, medicalAppointmentRepository, medicalRecordRepository := initializeRepositories(db)
	patientController, medicalAppointmentController, medicalRecordController := initializeControllers(patientRepository, medicalAppointmentRepository, medicalRecordRepository)
	routes.Patient(router, patientController)
	routes.MedicalRecord(router, medicalRecordController)
	routes.MedicalAppointment(router, medicalAppointmentController)
}

func initializeRepositories(db *gorm.DB) (patientRepository *services.PatientStorage, medicalAppointmentRepository *services.MedicalAppointmentStorage, medicalRecordRepository *services.MedicalRecordStorage) {
	patientRepository = services.NewPatientStorage(db)
	medicalAppointmentRepository = services.NewMedicalAppointmentStorage(db)
	medicalRecordRepository = services.NewMedicalRecordStorage(db)
	return
}

func initializeControllers(patientRepository services.PatientRepository, medicalAppointmentRepository services.MedicalAppointmentRepository, medicalRecordRepository services.MedicalRecordRepository) (patientController *controllers.PatientController, medicalAppointmentController *controllers.MedicalAppointmentController, medicalRecordController *controllers.MedicalRecordController) {
	patientController = controllers.NewPatientController(patientRepository)
	medicalAppointmentController = controllers.NewMedicalAppointmentController(medicalAppointmentRepository)
	medicalRecordController = controllers.NewMedicalRecordController(medicalRecordRepository)
	return
}
