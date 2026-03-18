package routes

import (
	"nextmed-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupBookingRoutes configures the booking and slot engine routes
func SetupBookingRoutes(r *gin.Engine) {
	bookingController := &controllers.BookingController{}

	v1 := r.Group("/api/v1/bookings")
	{
		v1.GET("/slots", bookingController.GetAvailableSlots)
		v1.POST("/create", bookingController.CreateAppointment)
	}
}
