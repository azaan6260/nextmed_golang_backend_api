package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TriageChat represents a message in an AI triage conversation
type TriageChat struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientProfileID uuid.UUID `gorm:"type:uuid;index;not null" json:"patient_profile_id"`
	Message          string    `gorm:"type:text;not null" json:"message"`
	SenderType       string    `gorm:"type:varchar(10);not null" json:"sender_type"` // 'USER' or 'AI'
	CreatedAt        time.Time `json:"created_at"`
}

// BeforeCreate GORM hook to generate UUID for the TriageChat
func (tc *TriageChat) BeforeCreate(tx *gorm.DB) (err error) {
	tc.ID = uuid.New()
	return
}
