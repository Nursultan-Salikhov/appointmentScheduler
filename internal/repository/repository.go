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
}

type Repository struct {
	Authorization
	Schedule
	Appointment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Schedule:      NewSchedulePostgres(db),
		Appointment:   NewAppointmentPostgres(db),
	}
}
