package models

type Schedule struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	WorkDay   string `json:"work_day" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

type Appointment struct {
	Id      int
	AppDay  string
	AppTime string
}
