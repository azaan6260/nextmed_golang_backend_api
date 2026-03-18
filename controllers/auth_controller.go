package controllers

import (
	"fmt"
	"net/http"
	"nextmed-backend/config"
	"nextmed-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthController handles authentication logic
type AuthController struct{}

// SendOTPRequest represents the payload for sending OTP
type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// VerifyOTPRequest represents the payload for verifying OTP
type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	OTP         string `json:"otp" binding:"required"`
}

// SendOTP mocks sending an OTP to the user's phone
func (a *AuthController) SendOTP(c *gin.Context) {
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock OTP sending
	mockOTP := "123456"
	fmt.Printf("Mock sending OTP %s to %s...\n", mockOTP, req.PhoneNumber)

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
	})
}

// VerifyOTP verifies the OTP and handles user/profile creation
func (a *AuthController) VerifyOTP(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hardcoded OTP for development
	if req.OTP != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	var user models.User
	var profile models.PatientProfile

	// Check if user exists
	result := config.DB.Where("phone = ?", req.PhoneNumber).First(&user)
	if result.Error != nil {
		// User doesn't exist, create new User and SELF Profile
		user = models.User{
			Phone: req.PhoneNumber,
			Role:  "patient",
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		profile = models.PatientProfile{
			LinkedUserID: user.ID,
			FullName:     "New Patient", // Placeholder
			Relationship: "SELF",
			DOB:          time.Now(), // Placeholder
		}
		if err := config.DB.Create(&profile).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
			return
		}
	} else {
		// User exists, find their SELF profile
		config.DB.Where("linked_user_id = ? AND relationship = ?", user.ID, "SELF").First(&profile)
	}

	// Generate Mock JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret")) // Use env variable in production
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":            tokenString,
		"user":             user,
		"active_profile":   profile,
		"managed_profiles": []models.PatientProfile{}, // Empty for now
	})
}
