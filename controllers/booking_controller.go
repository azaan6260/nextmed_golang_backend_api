package controllers

import (
	"net/http"
	"nextmed-backend/config"
	"nextmed-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingController struct{}

// GetAvailableSlots returns a list of free time slots for a doctor on a specific date
func (b *BookingController) GetAvailableSlots(c *gin.Context) {
	doctorID := c.Query("doctor_id")
	dateStr := c.Query("date") // Expected format: YYYY-MM-DD

	if doctorID == "" || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor_id and date are required"})
		return
	}

	// 1. Fetch FacilityDoctor record
	var fd models.FacilityDoctor
	if err := config.DB.Where("doctor_id = ?", doctorID).First(&fd).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor schedule not found"})
		return
	}

	// 2. Parse shift times
	startTime, err := time.Parse("15:04", fd.ShiftStart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid shift start time"})
		return
	}
	endTime, err := time.Parse("15:04", fd.ShiftEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid shift end time"})
		return
	}

	// 3. Generate all possible slots
	allSlots := []string{}
	current := startTime
	for current.Before(endTime) {
		allSlots = append(allSlots, current.Format("15:04"))
		current = current.Add(time.Duration(fd.AvgMinsPerPatient) * time.Minute)
	}

	// 4. Fetch existing appointments for that doctor on that date
	var appointments []models.Appointment
	// We filter by date range for the specific day
	startOfDay, _ := time.Parse("2006-01-02", dateStr)
	endOfDay := startOfDay.Add(24 * time.Hour)

	config.DB.Where("doctor_id = ? AND scheduled_time >= ? AND scheduled_time < ?", doctorID, startOfDay, endOfDay).Find(&appointments)

	// 5. Filter out booked slots
	bookedMap := make(map[string]bool)
	for _, app := range appointments {
		bookedMap[app.ScheduledTime.Format("15:04")] = true
	}

	availableSlots := []string{}
	for _, slot := range allSlots {
		if !bookedMap[slot] {
			availableSlots = append(availableSlots, slot)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor_id":       doctorID,
		"date":            dateStr,
		"available_slots": availableSlots,
	})
}

// CreateAppointmentRequest payload
type CreateAppointmentRequest struct {
	PatientProfileID uint    `json:"patient_profile_id" binding:"required"`
	BookedByUserID   uint    `json:"booked_by_user_id" binding:"required"`
	DoctorID         uint    `json:"doctor_id" binding:"required"`
	FacilityID       uint    `json:"facility_id" binding:"required"`
	ScheduledTime    string  `json:"scheduled_time" binding:"required"` // Format: "2006-01-02 15:04"
	PlatformFee      float64 `json:"platform_fee"`
	DoctorFee        float64 `json:"doctor_fee"`
}

// CreateAppointment handles the booking of a new appointment
func (b *BookingController) CreateAppointment(c *gin.Context) {
	var req CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduledTime, err := time.Parse("2006-01-02 15:04", req.ScheduledTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scheduled_time format. Use YYYY-MM-DD HH:MM"})
		return
	}

	// Double check if slot is already taken (Race condition prevention)
	var count int64
	config.DB.Model(&models.Appointment{}).Where("doctor_id = ? AND scheduled_time = ?", req.DoctorID, scheduledTime).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This slot is already booked"})
		return
	}

	appointment := models.Appointment{
		PatientProfileID: req.PatientProfileID,
		BookedByUserID:   req.BookedByUserID,
		DoctorID:         req.DoctorID,
		FacilityID:       req.FacilityID,
		ScheduledTime:    scheduledTime,
		Status:           "PENDING",
		PlatformFee:      req.PlatformFee,
		DoctorFee:        req.DoctorFee,
		PaymentStatus:    "UNPAID",
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}
