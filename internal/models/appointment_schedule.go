package models

import "errors"

type Schedule struct {
	Id        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	WorkDay   string `json:"work_day" db:"work_day" binding:"required" `
	StartTime string `json:"start_time" db:"start_time" binding:"required"`
	EndTime   string `json:"end_time" db:"end_time" binding:"required"`
}

type UpdateSchedule struct {
	StartTime *string `json:"start_time"`
	EndTime   *string `json:"end_time"`
}

type Appointment struct {
	Id      int    `json:"id" db:"id"`
	AppDay  string `json:"app_day" binding:"required" db:"appointment_day"`
	AppTime string `json:"app_time" binding:"required" db:"appointment_time"`
}

type AllAppointmentData struct {
	Client       Client      `json:"client"`
	AppData      Appointment `json:"app_data"`
	NoticeSource string      `json:"notice_source"`
}

func (u UpdateSchedule) Validate() error {
	if u.StartTime == nil && u.EndTime == nil {
		return errors.New("UpdateSchedule struct don't have values")
	}
	return nil
}
