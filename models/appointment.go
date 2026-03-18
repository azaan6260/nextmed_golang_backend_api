package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment represents a booked medical consultation
type Appointment struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	PatientProfileID uint           `gorm:"index;not null" json:"patient_profile_id"`
	BookedByUserID   uint           `gorm:"index;not null" json:"booked_by_user_id"`
	DoctorID         uint           `gorm:"index;not null" json:"doctor_id"`
	FacilityID       uint           `gorm:"index;not null" json:"facility_id"`
	ScheduledTime    time.Time      `gorm:"not null" json:"scheduled_time"`
	Status           string         `gorm:"default:PENDING" json:"status"` // PENDING, CONFIRMED, CANCELLED, COMPLETED
	PlatformFee      float64        `json:"platform_fee"`
	DoctorFee        float64        `json:"doctor_fee"`
	PaymentStatus    string         `gorm:"default:UNPAID" json:"payment_status"` // UNPAID, PAID, REFUNDED
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
