package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/notices"
	"appointmentScheduler/internal/notices/email"
	"errors"
)

type NoticeService struct {
	notice *notices.Notice
}

type NoticeData struct {
	ClientName string
	Recipient  string
	Source     int
	Text       string
	SourceData interface{}
}

const (
	SourceMail = notices.MailMessage
)

func NewNoticeService(n *notices.Notice) *NoticeService {
	return &NoticeService{notice: n}
}

func (n NoticeService) Send(nd NoticeData) error {

	switch nd.Source {

	case SourceMail:
		es, ok := nd.SourceData.(models.EmailSettings)
		if !ok {
			return errors.New("service notice: the data type does not match the source (mail)")
		}

		var en email.EmailNotice

		en.From = es.Email
		en.Password = es.Password
		en.Host = es.Host
		en.Port = es.Port

		err := n.notice.SetType(SourceMail, en)
		if err != nil {
			return err
		}

	default:
		return errors.New("service notice: unknown service type")
	}

	return n.notice.SendMessage(nd.Recipient, nd.Text)
}
