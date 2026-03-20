package routes

import (
	"nextmed-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupTriageRoutes registers triage-related endpoints
func SetupTriageRoutes(r *gin.Engine) {
	triageController := new(controllers.TriageController)

	v1 := r.Group("/api/v1/triage")
	{
		v1.POST("/chat", triageController.HandleTriageChat)
	}
}
