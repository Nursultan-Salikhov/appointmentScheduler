package repository

import (
	"appointmentScheduler/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, passwordHash string) (models.User, error)
}

type Schedule interface {
	CreateWorkDay(userId int, workDay, startTime, endTime string) (int, error)
	GetSchedules(userId int) ([]models.Schedule, error)
	Update(userId int, day string, input models.UpdateSchedule) error
	Delete(userId int, day string) error
}

type Appointment interface {
	Create(appDate models.AllAppointmentDate) (int, error)
	CheckWorkDay(userId int, workDay string) bool
	Get(userId int, day string) ([]models.Appointment, error)
	GetClientInfo(userId int, day, time string) (models.Client, error)
	Update(userId, clientId int, newApp models.Appointment) error
	Delete(userId, clientId int) error
}

type NoticeTemplates interface {
	Create(userId int, nt models.NoticeTemplates) error
	Get(userId int) (models.NoticeTemplates, error)
	Update(userId int, nt models.UpdateNoticeTemplates) error
	Delete(userId int) error
}

type EmailSettings interface {
	Create(userId int, s models.EmailSettings) error
	Get(userId int) (models.EmailSettings, error)
	Update(userId int, s models.UpdateEmailSettings) error
	Delete(userId int) error
}

type Repository struct {
	Authorization
	Schedule
	Appointment
	NoticeTemplates
	EmailSettings
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		Schedule:        NewSchedulePostgres(db),
		Appointment:     NewAppointmentPostgres(db),
		NoticeTemplates: NewNoticeTemplatesPostgres(db),
		EmailSettings:   NewEmailSettingsPostgres(db),
	}
}
