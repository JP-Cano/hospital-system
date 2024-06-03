package entities

import (
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	Id             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name           string     `json:"name"`
	LastName       string     `json:"last_name"`
	Phone          string     `json:"phone"`
	Email          string     `json:"email"`
	HealthSeverity int        `json:"health_severity"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
