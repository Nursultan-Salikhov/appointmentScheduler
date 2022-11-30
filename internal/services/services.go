package services

import (
	"appointmentScheduler/internal/models"
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
	Create(appDate models.AllAppointmentDate) (int, error)
	Get(userId int, day string) ([]models.Appointment, error)
	GetClientInfo(userId int, day, time string) (models.Client, error)
}

type Service struct {
	Authorization
	Schedule
	Appointment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Schedule:      NewScheduleService(repo.Schedule),
		Appointment:   NewAppointmentService(repo.Appointment),
	}
}
