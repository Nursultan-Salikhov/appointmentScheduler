package repository

import (
	"appointmentScheduler/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AppointmentPostgres struct {
	db *sqlx.DB
}

func NewAppointmentPostgres(db *sqlx.DB) *AppointmentPostgres {
	return &AppointmentPostgres{db: db}
}

func (a *AppointmentPostgres) Create(appDate models.AllAppointmentDate) (int, error) {

	//Check if a date is available for appointment
	appointmentId, err := a.getAppointmentId(appDate.Client.UserId, appDate.AppData.AppDay, appDate.AppData.AppTime)
	if err != nil {
		return 0, err
	}

	// start transaction
	tx, err := a.db.Begin()
	if err != nil {
		logrus.Errorln("Can't start transaction")
		return 0, err
	}

	//Create client query
	var clientId int
	query := fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		tableClients, columnUserId, columnName, columnPhoneNumber, columnEmail, columnTGUsername, columnDescription)

	row := tx.QueryRow(query, appDate.Client.UserId, appDate.Client.Name, appDate.Client.PhoneNumber,
		appDate.Client.Email, appDate.Client.TGUsername, appDate.Client.Description)
	err = row.Scan(&clientId)
	if err != nil {
		logrus.Errorln("Failed to create client")
		tx.Rollback()
		return 0, err
	}

	//Create client_id and app_id in clients_appointments table
	query = fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING id",
		tableClientsApps, columnClientId, columnAppointmentId)

	var id int
	row = tx.QueryRow(query, clientId, appointmentId)
	err = row.Scan(&id)
	if err != nil {
		logrus.Errorf("Failed creating data in %s", tableClientsApps)
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (a *AppointmentPostgres) CheckWorkDay(userId int, workDay string) bool {

	query := fmt.Sprintf("SELECT id FROM %s WHERE %s=$1 AND %s=$2",
		tableSchedules, columnUserId, columnWorkDay)

	row := a.db.QueryRow(query, userId, workDay)
	err := row.Err()
	if err != nil {
		return false
	}
	return true
}

func (a *AppointmentPostgres) getAppointmentId(userId int, day, time string) (int, error) {
	var id int
	//check if such a date exists in the database
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s=$1 AND %s=$2", tableAppointments,
		columnAppointmentDay, columnAppointmentTime)

	row := a.db.QueryRow(query, day, time)

	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		logrus.Infoln("missing date in", tableAppointments)
		logrus.Infof("%s", err.Error())

		//Most likely such a date does not exist, so we need to create it
		query = fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING id",
			tableAppointments, columnAppointmentDay, columnAppointmentTime)

		row = a.db.QueryRow(query, day, time)
		err = row.Scan(&id)
		if err != nil {
			logrus.Errorf("Failed create date in %s, appointment_id - %d", tableAppointments, id)
			return 0, err
		}
		return id, nil

	} else if err != nil {
		return 0, err
	} else {
		//check if the time is free for appointment
		//To do this, we check if the data in the three tables are related:
		//appointments (id=app_id) clients_appointments (client_id=id) clients
		logrus.Infoln("Check free appointment")
		query = fmt.Sprintf(`SELECT a.id FROM %s a INNER JOIN %s c_a ON c_a.%s = a.id 
							INNER JOIN %s c ON c.id = c_a.%s WHERE c.%s=$1 AND a.%s=$2 AND a.%s=$3`,
			tableAppointments, tableClientsApps, columnAppointmentId, tableClients, columnClientId,
			columnUserId, columnAppointmentDay, columnAppointmentTime)

		row = a.db.QueryRow(query, userId, day, time)
		var checkId int
		err = row.Scan(&checkId)

		logrus.Infof("error = %d", checkId)

		if err == sql.ErrNoRows {
			return id, nil
		} else if err == nil {
			return 0, errors.New("work time busy")
		} else if err != nil {
			return 0, err
		}
	}

	return 0, errors.New("unknown error")
}

func (a *AppointmentPostgres) Get(userId int, day string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	query := fmt.Sprintf(`SELECT a.id, a.%s, a.%s FROM %s a INNER JOIN %s ca ON a.id=ca.%s
  						INNER JOIN %s c ON ca.%s=c.id WHERE c.%s=$1 AND a.%s=$2`,
		columnAppointmentDay, columnAppointmentTime, tableAppointments, tableClientsApps,
		columnAppointmentId, tableClients, columnClientId, columnUserId, columnAppointmentDay)

	err := a.db.Select(&appointments, query, userId, day)
	if err != nil {
		return nil, err
	}

	for id, elem := range appointments {
		appointments[id].AppDay = correctDateFormat(elem.AppDay)
		appointments[id].AppTime = correctTimeFormat(elem.AppTime)
	}

	return appointments, nil
}

func (a *AppointmentPostgres) GetClientInfo(userId int, day, time string) (models.Client, error) {
	var clientInfo models.Client

	query := fmt.Sprintf(`SELECT c.id, c.%s, c.%s, c.%s, c.%s, c.%s, c.%s FROM %s c INNER JOIN %s ca ON c.id=ca.%s 
 						INNER JOIN %s a ON ca.%s=a.id WHERE c.%s=$1 AND a.%s=$2 AND a.%s=$3`,
		columnUserId, columnName, columnPhoneNumber, columnEmail, columnTGUsername, columnDescription,
		tableClients, tableClientsApps, columnClientId, tableAppointments,
		columnAppointmentId, columnUserId, columnAppointmentDay, columnAppointmentTime)

	err := a.db.Get(&clientInfo, query, userId, day, time)
	if err != nil {
		return models.Client{}, err
	}

	return clientInfo, nil
}
