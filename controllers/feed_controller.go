package controllers

import (
	"fmt"
	"net/http"
	"nextmed-backend/config"
	"nextmed-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FeedController struct{}

// CreatePostRequest payload
type CreatePostRequest struct {
	DoctorID    uint   `json:"doctor_id" binding:"required"`
	ContentText string `json:"content_text" binding:"required"`
	CategoryTag string `json:"category_tag"`
	MediaURL    string `json:"media_url"`
	MediaType   string `json:"media_type"` // 'VIDEO_HLS', 'IMAGE', or 'TEXT'
}

// CreatePost handles the creation of a new social feed post
func (f *FeedController) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		DoctorID:    req.DoctorID,
		ContentText: req.ContentText,
		CategoryTag: req.CategoryTag,
		MediaURL:    req.MediaURL,
		MediaType:   req.MediaType,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetUploadURL mocks generating a secure upload URL for media
func (f *FeedController) GetUploadURL(c *gin.Context) {
	doctorID := c.Query("doctor_id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor_id is required"})
		return
	}

	// Mock server-to-server call to Cloudflare Stream or AWS S3
	videoUID := uuid.New().String()
	mockUploadURL := fmt.Sprintf("https://upload.cloudflare.com/stream/%s", videoUID)

	c.JSON(http.StatusOK, gin.H{
		"upload_url": mockUploadURL,
		"video_uid":  videoUID,
	})
}

// GetFeed fetches the latest posts for the social feed
func (f *FeedController) GetFeed(c *gin.Context) {
	var posts []models.Post

	// Fetch posts ordered by latest first
	if err := config.DB.Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feed"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
