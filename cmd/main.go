package main

import (
	"appointmentScheduler/internal/server"
	"appointmentScheduler/internal/services"
	"appointmentScheduler/internal/transport/handlers"
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
