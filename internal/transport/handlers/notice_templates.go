package handlers

import (
	"appointmentScheduler/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateNoticeTemplates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var nt models.NoticeTemplates
	err = c.BindJSON(&nt)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.NoticeTemplates.Create(userId, nt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (h *Handler) GetNoticeTemplates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	nt, err := h.service.NoticeTemplates.Get(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nt)
}

func (h *Handler) UpdateNoticeTemplates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var unt models.UpdateNoticeTemplates
	err = c.BindJSON(&unt)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.NoticeTemplates.Update(userId, unt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (h *Handler) DeleteNoticeTemplates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.service.NoticeTemplates.Delete(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}
