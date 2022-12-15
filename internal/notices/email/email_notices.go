package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type EmailNotice struct {
	From     string
	Password string
	Host     string
	Port     int
}

func (e EmailNotice) SendMessage(recipient, text string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", e.From)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "appointment reminder")

	m.SetBody("text/plain", text)

	d := gomail.NewDialer(e.Host, e.Port, e.From, e.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}
