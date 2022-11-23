package services

import "appointmentScheduler/internal/models"

type Authorization interface {
	CreateUser(user models.Users) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Schedule interface {
}

type Note interface {
}

type Service struct {
	Authorization
	Schedule
	Note
}

func NewService() *Service {
	return &Service{}
}
