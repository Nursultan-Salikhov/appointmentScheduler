package repository

import (
	"appointmentScheduler/internal/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
		schedules[id].WorkDay = correctDateFormat(elem.WorkDay)
		schedules[id].StartTime = correctTimeFormat(elem.StartTime)
		schedules[id].EndTime = correctTimeFormat(elem.EndTime)
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

	res, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	numb, err := res.RowsAffected()
	if err != nil {
		return err
	} else if numb == 0 {
		return errors.New("update is not possible because there is no element")
	}

	return nil
}

func (s *SchedulePostgres) Delete(userId int, day string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1 AND %s=$2",
		tableSchedules, columnUserId, columnWorkDay)

	res, err := s.db.Exec(query, userId, day)
	if err != nil {
		return err
	}

	numb, err := res.RowsAffected()
	if err != nil {
		logrus.Error("RowsAffected failed")
		return err
	}

	if numb == 0 {
		return errors.New("deletion is not possible because there is no element")
	}

	return nil
}
