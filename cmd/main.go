package main

import (
	"appointmentScheduler/internal/server"
	"appointmentScheduler/internal/services"
	"appointmentScheduler/internal/transport/handlers"
	"log"
)

func main() {
	service := services.NewService()
	handler := handlers.NewHandler(service)
	srv := new(server.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
