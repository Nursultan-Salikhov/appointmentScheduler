package models

import "errors"

type EmailSettings struct {
	Status   bool   `json:"status" db:"status"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
}

type UpdateEmailSettings struct {
	Status   *bool   `json:"status"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Host     *string `json:"host"`
	Port     *string `json:"port"`
}

func (u UpdateEmailSettings) Validate() error {
	if u.Status == nil && u.Email == nil && u.Password == nil &&
		u.Host == nil && u.Port == nil {
		return errors.New("UpdateEmailSettings struct don't have values")
	}
	return nil
}
