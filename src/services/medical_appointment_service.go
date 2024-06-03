package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/src/entities"
	"time"
)

type MedicalAppointmentRepository interface {
	CreateAppointment(appointment entities.MedicalAppointment) error
	GetDoctorAppointments(doctorId uuid.UUID) ([]entities.DoctorAppointment, error)
}

type MedicalAppointmentStorage struct {
	db *gorm.DB
}

func NewMedicalAppointmentStorage(db *gorm.DB) *MedicalAppointmentStorage {
	return &MedicalAppointmentStorage{db: db}
}

func (m *MedicalAppointmentStorage) CreateAppointment(appointment entities.MedicalAppointment) error {
	return m.db.Create(&appointment).Error
}

func (m *MedicalAppointmentStorage) GetDoctorAppointments(doctorId uuid.UUID) ([]entities.DoctorAppointment, error) {
	var doctorAppointments []entities.DoctorAppointment
	if err := m.db.
		Table("medical_appointments").
		Select("patients.name, patients.last_name, patients.email, patients.phone, patients.health_severity, medical_appointments.date_time, medical_appointments.reason").
		Joins("join patients on medical_appointments.patient_id = patients.id").
		Where("medical_appointments.doctor_id = ? AND medical_appointments.date_time > ?", doctorId, time.Now()).
		Scan(&doctorAppointments).Error; err != nil {
		return nil, err
	}
	return doctorAppointments, nil
}
