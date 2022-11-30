package models

type Client struct {
	Id          int    `json:"id" db:"id"'`
	UserId      int    `json:"user_id" db:"user_id"`
	Name        string `json:"name" db:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	TGUsername  string `json:"tg_username" db:"tg_username"`
	Description string `json:"description" db:"description"`
}
