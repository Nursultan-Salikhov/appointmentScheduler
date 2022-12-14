package handlers

import (
	"appointmentScheduler/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateEmailSet(c *gin.Context) {
	var es models.EmailSettings
	err := c.BindJSON(&es)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.service.Settings.EmailSettings.Create(userId, es)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (h *Handler) GetEmailSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	es, err := h.service.Settings.EmailSettings.Get(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, es)
}

func (h *Handler) UpdateEmailSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var ues models.UpdateEmailSettings
	err = c.BindJSON(&ues)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Settings.EmailSettings.Update(userId, ues)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (h *Handler) DeleteEmailSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.service.Settings.EmailSettings.Delete(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}
