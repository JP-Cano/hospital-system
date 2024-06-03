package entities

import (
	"github.com/google/uuid"
	"time"
)

type MedicalRecord struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientId uuid.UUID `gorm:"type:uuid" json:"-"`
	Date      time.Time `json:"date"`
	Diagnosis string    `json:"diagnosis"`
	Treatment string    `json:"treatment"`
}

type DoctorReport struct {
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	DateTime  time.Time `json:"date_time"`
	Diagnosis string    `json:"diagnosis"`
}

type UpdateMedicalRecord struct {
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
}
