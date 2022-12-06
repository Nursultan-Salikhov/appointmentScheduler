package repository

import (
	"appointmentScheduler/internal/models"
	"github.com/jmoiron/sqlx"
)

type EmailSettingsPostgres struct {
	db *sqlx.DB
}

func NewEmailSettingsPostgres(db *sqlx.DB) *EmailSettingsPostgres {
	return &EmailSettingsPostgres{db: db}
}

func (e EmailSettingsPostgres) Create(userId int, s models.EmailSettings) error {
	//TODO implement me
	panic("implement me")
}

func (e EmailSettingsPostgres) Get(userId int) (models.NoticeTemplates, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmailSettingsPostgres) Update(userId int, s models.UpdateEmailSettings) error {
	//TODO implement me
	panic("implement me")
}

func (e EmailSettingsPostgres) Delete(userId int) error {
	//TODO implement me
	panic("implement me")
}
