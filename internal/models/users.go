package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type SignInData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmailSettings struct {
	Status   bool   `json:"status" db:"status"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
