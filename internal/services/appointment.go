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

func (a *AppointmentService) Get(userId int, day string) ([]models.Appointment, error) {
	if a.repo.CheckWorkDay(userId, day) {
		return a.repo.Get(userId, day)
	}
	return nil, errors.New("there is no data on the requested date (no working day was created)")
}

func (a *AppointmentService) GetClientInfo(userId int, day, time string) (models.Client, error) {
	return a.repo.GetClientInfo(userId, day, time)
}

func (a *AppointmentService) Update(userId, clientId int, newApp models.Appointment) error {

	if a.repo.CheckWorkDay(userId, newApp.AppDay) {
		return a.repo.Update(userId, clientId, newApp)
	}
	return errors.New("there is no data on the requested date(no working day was created")
}

func (a *AppointmentService) Delete(userId, clientId int) error {
	return a.repo.Delete(userId, clientId)
}
