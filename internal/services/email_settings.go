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
	return e.repo.Create(userId, s)
}

func (e *EmailSettingsService) Get(userId int) (models.NoticeTemplates, error) {
	return e.repo.Get(userId)
}

func (e *EmailSettingsService) Update(userId int, s models.UpdateEmailSettings) error {
	if err := s.Validate(); err != nil {
		return err
	}
	return e.repo.Update(userId, s)
}

func (e *EmailSettingsService) Delete(userId int) error {
	return e.repo.Delete(userId)
}
