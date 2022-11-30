package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
	"errors"
)

type AppointmentService struct {
	repo repository.Appointment
}

func NewAppointmentService(repo repository.Appointment) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (a *AppointmentService) Create(appData models.AllAppointmentDate) (int, error) {
	if a.repo.CheckWorkDay(appData.Client.UserId, appData.AppData.AppDay) {
		return a.repo.Create(appData)
	}
	return 0, errors.New("no date for appointment has been created")
}
