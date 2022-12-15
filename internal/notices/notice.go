package notices

import (
	"appointmentScheduler/internal/notices/email"
	"errors"
)

const (
	MailMessage = iota
)

type Message interface {
	SendMessage(recipient, text string) error
}

type Notice struct {
	Message
}

func NewNotice() *Notice {
	return &Notice{}
}

func (n *Notice) SetType(noticeType int, data interface{}) error {
	switch noticeType {
	case MailMessage:
		d, ok := data.(email.EmailNotice)
		if !ok {
			return errors.New("notices: email data structure transfer error ")
		}

		n.Message = d
		return nil

	default:
		return errors.New("notices: non-existent notice type")
	}
}
