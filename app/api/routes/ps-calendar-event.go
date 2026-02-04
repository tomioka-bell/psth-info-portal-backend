package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesCompanyCalendarEvent(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	CompanyCalendarEventRepository := repositories.NewCalendarEventRepository(db)
	CompanyCalendarEventService := services.NewCalendarEventService(CompanyCalendarEventRepository)
	CompanyCalendarEventHandler := handlers.NewCalendarEventHandler(CompanyCalendarEventService)

	app.Post("/create-calendar-event", CompanyCalendarEventHandler.CreateCalendarEventHandler)
	app.Get("/get-calendar-event/:event_id", CompanyCalendarEventHandler.GetCalendarEventByIDHandler)
	app.Get("/get-calendar-events", CompanyCalendarEventHandler.GetAllCalendarEventsHandler)
	app.Patch("/update-calendar-event/:event_id", CompanyCalendarEventHandler.UpdateCalendarEventHandler)
	app.Delete("/delete-calendar-event/:event_id", CompanyCalendarEventHandler.DeleteCalendarEventHandler)
	return app
}
