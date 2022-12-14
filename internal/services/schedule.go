package services

import (
	"appointmentScheduler/internal/models"
	"appointmentScheduler/internal/repository"
)

const (
	dateFormat = "2006-01-02"
	timeFormat = "15:04"
)

type ScheduleService struct {
	repo repository.Schedule
}

func NewScheduleService(repo repository.Schedule) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) CreateWorkDay(userId int, workDay, startTime, endTime string) (int, error) {
	err := checkDate(workDay)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateWorkDay(userId, workDay, startTime, endTime)
}

func (s *ScheduleService) GetSchedules(userId int) ([]models.Schedule, error) {
	return s.repo.GetSchedules(userId)
}

func (s *ScheduleService) Update(userId int, day string, input models.UpdateSchedule) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, day, input)
}

func (s *ScheduleService) Delete(userId int, day string) error {
	return s.repo.Delete(userId, day)
}
