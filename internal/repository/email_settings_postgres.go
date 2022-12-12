package repository

import (
	"appointmentScheduler/internal/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmailSettingsPostgres struct {
	db *sqlx.DB
}

func NewEmailSettingsPostgres(db *sqlx.DB) *EmailSettingsPostgres {
	return &EmailSettingsPostgres{db: db}
}

func (e EmailSettingsPostgres) Create(userId int, s models.EmailSettings) error {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6)",
		tableEmailSettings, columnUserId, columnStatus, columnEmail, columnPassword, columnHOST, columnPort)

	_, err := e.db.Exec(query, userId, s.Status, s.Email, []byte(s.Password), s.Host, s.Port)
	return err
}

func (e EmailSettingsPostgres) Get(userId int) (models.EmailSettings, error) {
	var es models.EmailSettings

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s=$1",
		columnStatus, columnEmail, columnPassword, columnHOST, columnPort,
		tableEmailSettings, columnUserId)

	err := e.db.Get(&es, query, userId)
	if err != nil {
		return models.EmailSettings{}, err
	}

	return es, nil
}

func (e EmailSettingsPostgres) Update(userId int, s models.UpdateEmailSettings) error {
	//TODO implement me
	panic("implement me")
}

func (e EmailSettingsPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1",
		tableEmailSettings, columnUserId)

	res, err := e.db.Exec(query, userId)
	if err != nil {
		return err
	}

	numb, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if numb == 0 {
		return errors.New("email_settings: deletion is not possible, because there is no element")
	}

	return nil
}
