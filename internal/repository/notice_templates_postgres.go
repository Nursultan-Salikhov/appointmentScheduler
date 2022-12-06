package repository

import (
	"appointmentScheduler/internal/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type NoticeTemplatesPostgres struct {
	db *sqlx.DB
}

func NewNoticeTemplatesPostgres(db *sqlx.DB) *NoticeTemplatesPostgres {
	return &NoticeTemplatesPostgres{db: db}
}

func (n NoticeTemplatesPostgres) Create(userId int, nt models.NoticeTemplates) error {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3)",
		tableNoticesTemplates, columnUserId, columnAppointmentTemplate, columnReminderTemplate)

	_, err := n.db.Exec(query, userId, nt.Appointment, nt.Reminder)
	return err
}

func (n NoticeTemplatesPostgres) Get(userId int) (models.NoticeTemplates, error) {
	var nt models.NoticeTemplates

	query := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s=$1",
		columnAppointmentTemplate, columnReminderTemplate, tableNoticesTemplates, columnUserId)

	err := n.db.Get(&nt, query, userId)
	if err != nil {
		return models.NoticeTemplates{}, err
	}
	return nt, nil
}

func (n NoticeTemplatesPostgres) Update(userId int, nt models.UpdateNoticeTemplates) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if nt.Appointment != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", columnAppointmentTemplate, argId))
		args = append(args, *nt.Appointment)
		argId++
	}

	if nt.Reminder != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", columnReminderTemplate, argId))
		args = append(args, *nt.Reminder)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE %s=$%d`, tableNoticesTemplates, setQuery,
		columnUserId, argId)

	args = append(args, userId)

	_, err := n.db.Exec(query, args...)
	return err
}

func (n NoticeTemplatesPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", tableNoticesTemplates, columnUserId)

	res, err := n.db.Exec(query, userId)
	if err != nil {
		return err
	}

	numb, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if numb == 0 {
		return errors.New("notice_templates: deletion is not possible, because there is no element")
	}

	return nil
}
