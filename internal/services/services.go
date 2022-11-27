package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	//ParseToken(accessToken string) (int, error)
}

type Schedule interface {
}

type Appointment interface {
}

type Service struct {
	Authorization
	Schedule
	Appointment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
