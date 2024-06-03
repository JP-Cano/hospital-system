package entities

import (
	"github.com/google/uuid"
	"time"
)

type MedicalAppointment struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientId uuid.UUID `gorm:"type:uuid" json:"patient_id"`
	DoctorId  uuid.UUID `gorm:"type:uuid" json:"doctor_id"`
	DateTime  time.Time `json:"date_time"`
	Reason    string    `json:"reason"`
}

type MedicalAppointmentDto struct {
	PatientId uuid.UUID `json:"patient_id"`
	DoctorId  uuid.UUID `json:"doctor_id"`
	DateTime  string    `json:"date_time"`
	Reason    string    `json:"reason"`
}
