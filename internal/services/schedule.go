package services

import (
	"appointmentScheduler/internal/repository"
	"errors"
	"time"
)

const (
	dateFormat = "2006.01.02"
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

func checkDate(day string) error {
	now := time.Now()
	now.Add(-(time.Hour * 48))

	workDay, err := time.Parse(dateFormat, day)
	if err != nil {
		return err
	}

	if now.After(workDay) {
		return errors.New("entered a date that has already passed")
	}
	return nil
}
