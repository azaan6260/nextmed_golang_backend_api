package routes

import (
	"nextmed-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures the authentication routes
func SetupAuthRoutes(r *gin.Engine) {
	authController := &controllers.AuthController{}

	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/send-otp", authController.SendOTP)
		v1.POST("/verify-otp", authController.VerifyOTP)
	}
}
