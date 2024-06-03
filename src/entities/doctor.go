package entities

import (
	"github.com/google/uuid"
	"time"
)

type Doctor struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Specialty string    `json:"specialty"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
}

type DoctorAppointment struct {
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	HealthSeverity int       `json:"health_severity"`
	DateTime       time.Time `json:"date_time"`
	Reason         string    `json:"reason"`
}
