package handlers

import (
	"appointmentScheduler/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	user := router.Group("/user", h.userIdentity)
	{
		schedules := user.Group("/schedules")
		{
			schedules.POST("/", h.CreateWorkDay)

			schedules.GET("/", h.GetSchedules)

			schedules.GET("/:day", h.GetDetailedSchedule)

			schedules.PUT("/:day", h.UpdateDay)

			schedules.DELETE("/:day", h.DeleteDay)
		}

		appointments := user.Group("/appointments")
		{
			appointments.POST("/", h.CreateAppointment)
			appointments.GET("/", h.GetAppointments)
			appointments.DELETE("/", h.DeleteAppointment)
		}
	}

	appointment := router.Group("/:id/appointment", h.AppointmentIdentity)
	{
		appointment.POST("/", h.CreateAppointment)
		appointment.GET("/available", h.GetAvailableAppointment)
	}

	return router
}
