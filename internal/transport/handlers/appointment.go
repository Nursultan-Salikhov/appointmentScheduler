package handlers

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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

	noticeStatus := "done"
	nt, err := h.service.Settings.NoticeTemplates.Get(allAppData.Client.UserId)
	if err != nil {
		logrus.Errorf("failed to get notice templates")
		logrus.Error(err.Error())
		noticeStatus = "internal error"
	} else {
		noticeStatus = sendNotice(nt.Appointment, allAppData.NoticeSource, &allAppData.Client, h)

		// for reminder
		go func(userId int, noticeSource, day, hour string) {
			reminderTime := day + " " + hour + " +0300"

			t, err := time.Parse("2006-01-02 15:04 -0700", reminderTime)
			if err != nil {
				logrus.Error("reminder don't start")
				logrus.Error(err.Error())
				return
			}

			// A reminder an hour before the appointment
			t = t.Add(-(time.Hour))
			time.Sleep(time.Until(t))

			go func() {
				nt, err := h.service.Settings.NoticeTemplates.Get(userId)
				if err != nil {
					logrus.Errorf("failed to get notice templates for reminder")
					logrus.Error(err.Error())
					return
				}

				client, err := h.service.Appointment.GetClientInfo(userId, day, hour)
				if err != nil {
					logrus.Error("failed to get client info")
					logrus.Error(err.Error())
					return
				}

				logrus.Infoln("sendResult - ", sendNotice(nt.Reminder, noticeSource, &client, h), hour)
			}()

		}(allAppData.Client.UserId, allAppData.NoticeSource, allAppData.AppData.AppDay, allAppData.AppData.AppTime)
	}

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
	hour := c.Param("time")

	clientInfo, err := h.service.Appointment.GetClientInfo(userId, day, hour)
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

func sendNotice(text, noticeSource string, client *models.Client, h *Handler) string {
	var nd services.NoticeData

	nd.ClientName = client.Name
	nd.Text = text

	switch noticeSource {

	case noticeDisable:
		return noticeDisable

	case noticeMail:
		es, err := h.service.Settings.EmailSettings.Get(client.UserId)
		if err != nil {
			logrus.Errorf("error receiving mail settings, user_id - %d", client.UserId)
			logrus.Error(err.Error())
			return "internal error"
		}

		if es.Status == false {
			logrus.Warnf("the selected notification sending source is not available: user_id -%d, source - %s",
				client.UserId, noticeMail)
			return "not available"
		}

		nd.SourceData = es
		nd.Source = services.SourceMail
		nd.Recipient = client.Email

	default:
		logrus.Warnf("unknown notice source - %s : user_id - %d", noticeSource, client.UserId)
		return fmt.Sprintf("unknown source: %s", noticeSource)

	}

	err := h.service.Notices.Send(nd)
	if err != nil {
		logrus.Errorf("notice send failed: user_id - %d", client.UserId)
		logrus.Error(err.Error())
		return "input error"
	}

	return "done"
}
