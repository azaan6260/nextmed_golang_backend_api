package controllers

import (
	"net/http"
	"nextmed-backend/config"
	"nextmed-backend/models"

	"github.com/gin-gonic/gin"
)

type FeedController struct{}

// CreatePostRequest payload
type CreatePostRequest struct {
	DoctorID    uint   `json:"doctor_id" binding:"required"`
	ContentText string `json:"content_text" binding:"required"`
	CategoryTag string `json:"category_tag"`
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
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
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
