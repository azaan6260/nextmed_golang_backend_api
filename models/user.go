package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the login credentials and core account info
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Phone     string         `gorm:"uniqueIndex;not null" json:"phone"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Role      string         `gorm:"default:patient" json:"role"` // e.g., patient, doctor, admin
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
