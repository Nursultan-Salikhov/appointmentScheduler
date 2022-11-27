package repository

import (
	"appointmentScheduler/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3) RETURNING id",
		tableUsers, columnName, columnUsername, columnPasswordHash)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (models.User, error) {
	var output models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1 AND %s=$2",
		tableUsers, columnUsername, columnPasswordHash)

	err := r.db.Get(&output, query, username, passwordHash)
	return output, err
}
