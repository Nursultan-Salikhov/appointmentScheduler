package models

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
	Appointment string `json:"appointment"`
	Reminder    string `json:"reminder"`
}

type UpdateNoticeTemplates struct {
	Appointment *string `json:"appointment"`
	Reminder    *string `json:"reminder"`
}

type EmailSettings struct {
	Status   bool   `json:"status" db:"status"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
}
