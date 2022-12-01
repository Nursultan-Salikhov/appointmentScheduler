package handlers

import (
	"appointmentScheduler/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateAppointment(c *gin.Context) {

	var allAppDate models.AllAppointmentDate
	err := c.BindJSON(&allAppDate)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	day := c.Param("day")

	appointments, err := h.service.Appointment.Get(userId, day)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (h *Handler) DeleteAppointment(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	clientIds := c.Param("clientId")
	clientId, err := strconv.Atoi(clientIds)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Appointment.Delete(userId, clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (h *Handler) GetAvailableAppointment(c *gin.Context) {

}

func (h *Handler) GetClientInfo(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	day := c.Param("day")
	time := c.Param("time")

	clientInfo, err := h.service.Appointment.GetClientInfo(userId, day, time)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, clientInfo)
}

func (h *Handler) UpdateAppointment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	clientIdS := c.Param("clientId")
	clientId, err := strconv.Atoi(clientIdS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var dataUpdate models.Appointment
	err = c.BindJSON(&dataUpdate)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Appointment.Update(userId, clientId, dataUpdate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
