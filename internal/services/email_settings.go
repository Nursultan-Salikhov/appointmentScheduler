package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
)

type EmailSettingsService struct {
	repo repository.EmailSettings
}

func NewEmailSettingsService(repo repository.EmailSettings) *EmailSettingsService {
	return &EmailSettingsService{repo: repo}
}

func (e *EmailSettingsService) Create(userId int, s models.EmailSettings) error {
	var err error
	s.Password, err = encryptAES(s.Password)
	if err != nil {
		return err
	}

	return e.repo.Create(userId, s)
}

func (e *EmailSettingsService) Get(userId int) (models.EmailSettings, error) {
	esData, err := e.repo.Get(userId)
	if err != nil {
		return models.EmailSettings{}, err
	}

	esData.Password, err = decryptAES([]byte(esData.Password))
	if err != nil {
		return models.EmailSettings{}, err
	}

	return esData, nil
}

func (e *EmailSettingsService) Update(userId int, s models.UpdateEmailSettings) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	if s.Password != nil {
		*s.Password, err = encryptAES(*s.Password)
		if err != nil {
			return err
		}
	}

	return e.repo.Update(userId, s)
}

func (e *EmailSettingsService) Delete(userId int) error {
	return e.repo.Delete(userId)
}
