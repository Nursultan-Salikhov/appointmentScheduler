package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
)

type NoticeTemplatesService struct {
	repo repository.NoticeTemplates
}

func NewNoticeTemplatesService(repo repository.NoticeTemplates) *NoticeTemplatesService {
	return &NoticeTemplatesService{repo: repo}
}

func (n NoticeTemplatesService) Create(userId int, nt models.NoticeTemplates) error {
	return n.repo.Create(userId, nt)
}

func (n NoticeTemplatesService) Get(userId int) (models.NoticeTemplates, error) {
	return n.repo.Get(userId)
}

func (n NoticeTemplatesService) Update(userId int, nt models.UpdateNoticeTemplates) error {
	if err := nt.Validate(); err != nil {
		return err
	}
	return n.repo.Update(userId, nt)
}

func (n NoticeTemplatesService) Delete(userId int) error {
	return n.repo.Delete(userId)
}
