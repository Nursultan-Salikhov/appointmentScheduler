package repository

type Authorization interface {
	CreateUser() (int, error)
	GetUser(username, password string) error
}

type Repository struct {
	Authorization
}
