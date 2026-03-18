package models

import (
	"time"

	"gorm.io/gorm"
)

// Facility represents a healthcare center (clinic, hospital, etc.)
type Facility struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Type      string         `json:"type"` // e.g., Clinic, Hospital, Lab
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// FacilityDoctor represents the mapping of a doctor to a facility with their schedule
type FacilityDoctor struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	DoctorID          uint           `gorm:"index;not null" json:"doctor_id"`
	FacilityID        uint           `gorm:"index;not null" json:"facility_id"`
	ConsultationFee   float64        `json:"consultation_fee"`
	ShiftStart        string         `json:"shift_start"`          // e.g., "10:00"
	ShiftEnd          string         `json:"shift_end"`            // e.g., "13:00"
	AvgMinsPerPatient int            `gorm:"default:15" json:"avg_mins_per_patient"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
