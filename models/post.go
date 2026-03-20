package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post represents a social feed entry created by a doctor
type Post struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	DoctorID    uint      `gorm:"index;not null" json:"doctor_id"` // Links to models.User (ID)
	ContentText string    `gorm:"type:text;not null" json:"content_text"`
	CategoryTag string    `json:"category_tag"`
	MediaURL    string    `json:"media_url"`
	MediaType   string    `json:"media_type"` // 'VIDEO_HLS', 'IMAGE', or 'TEXT'
	CreatedAt   time.Time `json:"created_at"`
}

// BeforeCreate GORM hook to generate UUID for the Post
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
