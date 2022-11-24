package main

import (
	"appointmentScheduler/internal/repository"
	"appointmentScheduler/internal/server"
	"appointmentScheduler/internal/services"
	"appointmentScheduler/internal/transport/handlers"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("Error initConfing: %s", err.Error())
	}
}

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("DB connect failed: %s", err.Error())
	}

	// database test
	var id int
	name := "Иван"
	username := "Ivan"
	password := "fuck_you"

	query := fmt.Sprintf("INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id")
	row := db.QueryRow(query, name, username, password)

	err = row.Scan(&id)
	if err != nil {
		logrus.Fatalf("DB don't return ID")
	} else {
		logrus.Println("user create succes, id = ", id)
	}

	// database test
	service := services.NewService()
	handler := handlers.NewHandler(service)
	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
