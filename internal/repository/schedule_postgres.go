package repository

import (
	"appointmentScheduler/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SchedulePostgres struct {
	db *sqlx.DB
}

func NewSchedulePostgres(db *sqlx.DB) *SchedulePostgres {
	return &SchedulePostgres{db: db}
}

func (s *SchedulePostgres) CreateWorkDay(userId int, workDay, startTime, endTime string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES ($1, $2, $3, $4) RETURNING id",
		tableSchedules, columnUserId, columnWorkDay, columnStartTime, columnEndTime)
	row := s.db.QueryRow(query, userId, workDay, startTime, endTime)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SchedulePostgres) GetSchedules(userId int) ([]models.Schedule, error) {
	var schedules []models.Schedule

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1",
		tableSchedules, columnUserId)

	err := s.db.Select(&schedules, query, userId)
	if err != nil {
		return nil, err
	}

	// Making the data more readable
	for id, elem := range schedules {
		before, _, found := strings.Cut(elem.WorkDay, "T")
		if found {
			schedules[id].WorkDay = before
		}

		_, after, found := strings.Cut(elem.StartTime, "T")
		if found {
			after = strings.TrimRight(after, "00Z")
			schedules[id].StartTime = strings.TrimRight(after, ":")
		}

		_, after, found = strings.Cut(elem.EndTime, "T")
		if found {
			after = strings.TrimRight(after, "00Z")
			schedules[id].EndTime = strings.TrimRight(after, ":")
		}
	}

	return schedules, nil
}

func (s *SchedulePostgres) Update(userId int, day string, input models.UpdateSchedule) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.StartTime != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", columnStartTime, argId))
		args = append(args, *input.StartTime)
		argId++
	}

	if input.EndTime != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", columnEndTime, argId))
		args = append(args, *input.EndTime)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE %s=$%d AND %s=$%d`,
		tableSchedules, setQuery, columnUserId, argId, columnWorkDay, argId+1)

	args = append(args, userId, day)

	_, err := s.db.Exec(query, args...)
	return err
}
