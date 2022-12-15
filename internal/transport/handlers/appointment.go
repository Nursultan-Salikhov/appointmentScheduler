package handlers

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	noticeDisable = "disable"
	noticeMail    = "mail"
)

func (h *Handler) CreateAppointment(c *gin.Context) {

	var allAppData models.AllAppointmentData
	err := c.BindJSON(&allAppData)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	allAppData.Client.UserId, err = getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.Appointment.Create(allAppData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	noticeStatus := sendNotice(id, &allAppData, h)

	c.JSON(http.StatusOK, gin.H{"status": "done",
		"id": id, "notice_status": noticeStatus})
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

func sendNotice(clientId int, allAppData *models.AllAppointmentData, h *Handler) string {
	var nd services.NoticeData

	nt, err := h.service.Settings.NoticeTemplates.Get(allAppData.Client.UserId)
	if err != nil {
		logrus.Errorf("failed to get notice templates")
		logrus.Error(err.Error())
		return "internal error"
	}
	nd.Text = nt.Appointment

	nd.ClientName = allAppData.Client.Name

	switch allAppData.NoticeSource {

	case noticeDisable:
		return noticeDisable

	case noticeMail:
		es, err := h.service.Settings.EmailSettings.Get(allAppData.Client.UserId)
		if err != nil {
			logrus.Errorf("error receiving mail settings, user_id - %d", allAppData.Client.UserId)
			logrus.Error(err.Error())
			return "internal error"
		}

		if es.Status == false {
			logrus.Warnf("the selected notification sending source is not available: user_id -%d, source - %s",
				allAppData.Client.UserId, noticeMail)
			return "not available"
		}

		nd.SourceData = es
		nd.Source = services.SourceMail
		nd.Recipient = allAppData.Client.Email

	default:
		logrus.Warnf("unknown notice source - %s : user_id - %d", allAppData.NoticeSource, allAppData.Client.UserId)
		return fmt.Sprintf("unknown source: %s", allAppData.NoticeSource)

	}

	err = h.service.Notices.Send(nd)
	if err != nil {
		logrus.Errorf("notice send failed: user_id - %d, client_id - %d", allAppData.Client.UserId, clientId)
		logrus.Error(err.Error())
		return "input error"
	}

	return "done"
}
