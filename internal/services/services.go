package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/notices"
	"appointmentScheduler/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Schedule interface {
	CreateWorkDay(userId int, workDay, startTime, endTime string) (int, error)
	GetSchedules(userId int) ([]models.Schedule, error)
	Update(userId int, day string, input models.UpdateSchedule) error
	Delete(userId int, day string) error
}

type Appointment interface {
	Create(appDate models.AllAppointmentData) (int, error)
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

type Notices interface {
	Send(nd NoticeData) error
}

type settings struct {
	NoticeTemplates
	EmailSettings
}

type Service struct {
	Authorization
	Schedule
	Appointment
	Notices
	Settings settings
}

func NewService(repo *repository.Repository, n *notices.Notice) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Schedule:      NewScheduleService(repo.Schedule),
		Appointment:   NewAppointmentService(repo.Appointment),
		Notices:       NewNoticeService(n),
		Settings: settings{
			NoticeTemplates: NewNoticeTemplatesService(repo.NoticeTemplates),
			EmailSettings:   NewEmailSettingsService(repo.EmailSettings),
		},
	}
}
