package controllers

import (
	"net/http"
	"nextmed-backend/config"
	"nextmed-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TriageController struct{}

// TriageChatRequest payload
type TriageChatRequest struct {
	PatientProfileID string `json:"patient_profile_id" binding:"required"`
	Message          string `json:"message" binding:"required"`
}

// HandleTriageChat processes user messages and generates AI responses
func (tc *TriageController) HandleTriageChat(c *gin.Context) {
	var req TriageChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profileID, err := uuid.Parse(req.PatientProfileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient_profile_id format"})
		return
	}

	// 1. Insert user message
	userMsg := models.TriageChat{
		PatientProfileID: profileID,
		Message:          req.Message,
		SenderType:       "USER",
	}
	if err := config.DB.Create(&userMsg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user message"})
		return
	}

	// 2. Mock AI delay
	time.Sleep(1500 * time.Millisecond)

	// 3. Generate dummy AI response
	aiResponseText := "Based on your symptoms, I recommend seeing a General Physician. Shall I find one nearby?"

	// 4. Insert AI response
	aiMsg := models.TriageChat{
		PatientProfileID: profileID,
		Message:          aiResponseText,
		SenderType:       "AI",
	}
	if err := config.DB.Create(&aiMsg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save AI response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": aiResponseText})
}
