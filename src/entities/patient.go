package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Patient struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigratePatients(db *gorm.DB) error {
	err := db.AutoMigrate(&Patient{})
	return err
}
