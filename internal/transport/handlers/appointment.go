package handlers

import (
	"appointmentScheduler/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateAppointment(c *gin.Context) {
	var allAppDate models.AllAppointmentDate
	c.BindJSON(&allAppDate)

	var err error
	allAppDate.Client.UserId, err = getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.Appointment.Create(allAppDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done",
		"id": id})
}

func (h *Handler) GetAppointments(c *gin.Context) {

}

func (h *Handler) DeleteAppointment(c *gin.Context) {

}

func (h *Handler) GetAvailableAppointment(c *gin.Context) {

}
