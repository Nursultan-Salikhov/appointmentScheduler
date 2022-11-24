package models

type Schedule struct {
	Id        int
	UserId    int
	WorkDay   string
	StartTime string
	EndTime   string
}

type Appointment struct {
	Id      int
	AppDay  string
	AppTime string
}
