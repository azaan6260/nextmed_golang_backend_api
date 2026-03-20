package routes

import (
	"nextmed-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupFeedRoutes configures the social feed routes
func SetupFeedRoutes(r *gin.Engine) {
	feedController := &controllers.FeedController{}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/posts", feedController.CreatePost)
		v1.GET("/posts/get-upload-url", feedController.GetUploadURL)
		v1.GET("/feed", feedController.GetFeed)
	}
}
