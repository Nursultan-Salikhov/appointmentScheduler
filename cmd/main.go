package main

import (
	"appointmentScheduler/internal/repository"
	"appointmentScheduler/internal/server"
	"appointmentScheduler/internal/services"
	"appointmentScheduler/internal/transport/handlers"
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

	rep := repository.NewRepository(db)
	service := services.NewService(rep)
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
