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
}

type Repository struct {
	Authorization
	Schedule
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Schedule:      NewSchedulePostgres(db),
	}
}
