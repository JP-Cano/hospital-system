package services

import (
	"gorm.io/gorm"
	"hospital-system/src/entities"
)

type PatientRepository interface {
	GetPatients() ([]entities.Patient, error)
	SearchPatients(name, phone, email string) ([]entities.Patient, error)
}

type PatientStorage struct {
	db *gorm.DB
}

func NewPatientStorage(db *gorm.DB) *PatientStorage {
	return &PatientStorage{db: db}
}

func (p *PatientStorage) GetPatients() ([]entities.Patient, error) {
	var patients []entities.Patient
	if err := p.db.Where("health_severity IN (?)", []int{1, 2, 3}).Order("health_severity").Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func (p *PatientStorage) SearchPatients(name, phone, email string) ([]entities.Patient, error) {
	var patients []entities.Patient
	query := p.db

	if name != "" {
		query = query.Where("name ILIKE ? OR last_name ILIKE ?", "%"+name+"%", "%"+name+"%")
	}
	if phone != "" {
		query = query.Where("phone ILIKE ?", "%"+phone+"%")
	}
	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
