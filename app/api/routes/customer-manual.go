package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesCustomerManual(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	CustomerManualRepository := repositories.NewCustomerManualRepository(db)
	CustomerManualService := services.NewCustomerManualService(CustomerManualRepository)
	CustomerManualHandler := handlers.NewCustomerManualHandler(CustomerManualService)

	app.Post("/create", CustomerManualHandler.CreateCustomerManualHandler)
	app.Post("/create-nested", CustomerManualHandler.CreateCustomerManualNestedHandler)
	app.Get("/list", CustomerManualHandler.GetAllCustomerManualHandler)
	app.Get("/tree", CustomerManualHandler.GetCustomerManualTreeHandler)
	app.Put("/update/:customer_manual_id", CustomerManualHandler.UpdateCustomerManualHandler)
	app.Delete("/delete/:customer_manual_id", CustomerManualHandler.DeleteCustomerManualHandler)

	return app
}
