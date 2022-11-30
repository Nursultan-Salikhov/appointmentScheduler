package models

type Client struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	TGUsername  string `json:"tg_username"`
	Description string `json:"description"`
}
