package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesOrganizationDoc(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	OrganizationDocRepository := repositories.NewOrganizationDocRepository(db)
	OrganizationDocService := services.NewOrganizationDocService(OrganizationDocRepository)
	OrganizationDocHandler := handlers.NewOrganizationDocHandler(OrganizationDocService)

	app.Post("/create", OrganizationDocHandler.CreateOrganizationDocHandler)
	app.Get("/list", OrganizationDocHandler.GetAllOrganizationDocHandler)
	app.Get("/department/:department", OrganizationDocHandler.GetOrganizationDocByDepartmentHandler)
	app.Put("/update/:organization_doc_id", OrganizationDocHandler.UpdateOrganizationDocHandler)
	app.Delete("/delete/:organization_doc_id", OrganizationDocHandler.DeleteOrganizationDocHandler)

	return app
}
