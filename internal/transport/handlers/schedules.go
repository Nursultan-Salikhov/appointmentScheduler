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
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
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

}

func (h *Handler) GetDetailedSchedule(c *gin.Context) {

}

func (h *Handler) UpdateDay(c *gin.Context) {

}

func (h *Handler) DeleteDay(c *gin.Context) {

}
