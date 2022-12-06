package models

import "errors"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type SignInData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NoticeTemplates struct {
	Appointment string `json:"appointment" db:"appointment_template"`
	Reminder    string `json:"reminder" db:"reminder_template"`
}

type UpdateNoticeTemplates struct {
	Appointment *string `json:"appointment"`
	Reminder    *string `json:"reminder"`
}

func (u UpdateNoticeTemplates) Validate() error {
	if u.Appointment == nil && u.Reminder == nil {
		return errors.New("UpdateNoticeTemplates struct don't have values")
	}
	return nil
}
