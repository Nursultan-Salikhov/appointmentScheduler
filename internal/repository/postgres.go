package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	tableUsers        = "users"
	tableSchedules    = "schedules"
	tableClients      = "clients"
	tableAppointments = "appointments"

	columnName            = "name"
	columnUsername        = "username"
	columnPasswordHash    = "password_hash"
	columnUserId          = "user_id"
	columnWorkDay         = "work_day"
	columnStartTime       = "start_time"
	columnEndTime         = "end_time"
	columnPhoneNumber     = "phone_number"
	columnEmail           = "email"
	columnTGUsername      = "tg_username"
	columnDescription     = "description"
	columnAppointmentDay  = "appointment_day"
	columnAppointmentTime = "appointment_time"
	columnClientId        = "client_id"
	columnAppointmentId   = "appointment_id"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	return db, err
}
