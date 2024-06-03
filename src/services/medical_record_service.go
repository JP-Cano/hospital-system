package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hospital-system/src/entities"
	"time"
)

type MedicalRecordRepository interface {
	UpdateMedicalRecord(patientID uuid.UUID, input entities.UpdateMedicalRecord) error
	GetDoctorReport(doctorI uuid.UUID, startDate, endDate time.Time) ([]entities.DoctorReport, error)
}

type MedicalRecordStorage struct {
	db *gorm.DB
}

func NewMedicalRecordStorage(db *gorm.DB) *MedicalRecordStorage {
	return &MedicalRecordStorage{db: db}
}

func (m *MedicalRecordStorage) UpdateMedicalRecord(patientID uuid.UUID, updateMedicalRecord entities.UpdateMedicalRecord) error {
	var medicalRecord entities.MedicalRecord
	if err := m.db.Where("patient_id = ?", patientID).First(&medicalRecord).Error; err != nil {
		return err
	}

	if updateMedicalRecord.Diagnosis != "" {
		medicalRecord.Diagnosis = updateMedicalRecord.Diagnosis
	}

	if updateMedicalRecord.Treatment != "" {
		medicalRecord.Treatment = updateMedicalRecord.Treatment
	}

	if err := m.db.Save(&medicalRecord).Error; err != nil {
		return err
	}

	return nil
}

func (m *MedicalRecordStorage) GetDoctorReport(doctorId uuid.UUID, startDate, endDate time.Time) ([]entities.DoctorReport, error) {
	var results []entities.DoctorReport
	if err := m.db.Table("medical_appointments").
		Select("patients.name, patients.last_name, medical_appointments.date_time, medical_records.diagnosis").
		Joins("join patients on medical_appointments.patient_id = patients.id").
		Joins("join medical_records on patients.id = medical_records.patient_id").
		Where("medical_appointments.doctor_id = ? AND medical_appointments.date_time BETWEEN ? AND ?", doctorId, startDate, endDate).
		Order("medical_appointments.date_time").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
