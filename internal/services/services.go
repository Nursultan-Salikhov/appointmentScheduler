package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
	"errors"
	"time"
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
	Create(appDate models.AllAppointmentDate) (int, error)
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
	Get(userId int) (models.NoticeTemplates, error)
	Update(userId int, s models.UpdateEmailSettings) error
	Delete(userId int) error
}

type Notices interface {
	SendMessage(recipient, text string) error
	CreateReminder(recipient, text string, rTime time.Time) error
}

type Service struct {
	Authorization
	Schedule
	Appointment
	NoticeTemplates
	EmailSettings
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repo.Authorization),
		Schedule:        NewScheduleService(repo.Schedule),
		Appointment:     NewAppointmentService(repo.Appointment),
		NoticeTemplates: NewNoticeTemplatesService(repo.NoticeTemplates),
		EmailSettings:   NewEmailSettingsService(repo.EmailSettings),
	}
}

func checkDate(day string) error {
	now := time.Now()
	now = now.Add(-(time.Hour * 24))

	workDay, err := time.Parse(dateFormat, day)
	if err != nil {
		return err
	}

	if now.After(workDay) {
		return errors.New("entered a date that has already passed")
	}
	return nil
}
