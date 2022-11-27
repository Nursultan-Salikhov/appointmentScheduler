package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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
