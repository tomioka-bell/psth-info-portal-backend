package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesOrganization(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	OrganizationRepository := repositories.NewOrganizationRepository(db)
	OrganizationService := services.NewOrganizationService(OrganizationRepository)
	OrganizationHandler := handlers.NewOrganizationHandler(OrganizationService)

	// CRUD routes
	app.Post("/create", OrganizationHandler.CreateOrganizationHandler)
	app.Get("/list", OrganizationHandler.GetAllOrganizationsHandler)
	app.Get("/category/:category", OrganizationHandler.GetOrganizationsByCategoryHandler)
	app.Get("/search", OrganizationHandler.SearchOrganizationsHandler)
	app.Get("/:id", OrganizationHandler.GetOrganizationHandler)
	app.Put("/:id", OrganizationHandler.UpdateOrganizationHandler)
	app.Delete("/:id", OrganizationHandler.DeleteOrganizationHandler)

	app.Post("/create-multiple", OrganizationHandler.CreateOrganizationsHandler)

	return app
}
