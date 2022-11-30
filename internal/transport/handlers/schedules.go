package handlers

import (
	"appointmentScheduler/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateWorkDay(c *gin.Context) {
	var input models.Schedule
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := h.service.Schedule.CreateWorkDay(userId, input.WorkDay, input.StartTime, input.EndTime)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": "Done",
	})
}

func (h *Handler) GetSchedules(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	schedules, err := h.service.Schedule.GetSchedules(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (h *Handler) UpdateDay(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	day := c.Param("day")

	var input models.UpdateSchedule
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Schedule.Update(userId, day, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func (h *Handler) DeleteDay(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	day := c.Param("day")

	err = h.service.Schedule.Delete(userId, day)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Done"})
}
