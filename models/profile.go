package models

import (
	"time"

	"gorm.io/gorm"
)

// PatientProfile represents the medical record/identity of a patient
type PatientProfile struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	LinkedUserID uint           `gorm:"index" json:"linked_user_id"` // The User who manages this profile
	FullName     string         `gorm:"not null" json:"full_name"`
	DOB          time.Time      `gorm:"type:date" json:"dob"`
	Relationship string         `gorm:"default:self" json:"relationship"` // self, spouse, child, etc.
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
