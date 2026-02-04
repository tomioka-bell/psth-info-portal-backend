package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesSafetyDocument(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	SafetyDocumentRepository := repositories.NewSafetyDocumentRepository(db)
	SafetyDocumentService := services.NewSafetyDocumentService(SafetyDocumentRepository)
	SafetyDocumentHandler := handlers.NewSafetyDocumentHandler(SafetyDocumentService)

	app.Post("/create", SafetyDocumentHandler.CreateSafetyDocumentHandler)
	app.Get("/list", SafetyDocumentHandler.GetAllSafetyDocumentHandler)
	app.Get("/category/:category", SafetyDocumentHandler.GetSafetyDocumentByCategoryHandler)
	app.Get("/department/:department", SafetyDocumentHandler.GetSafetyDocumentByDepartmentHandler)
	app.Put("/update/:safety_document_id", SafetyDocumentHandler.UpdateSafetyDocumentHandler)
	app.Delete("/delete/:safety_document_id", SafetyDocumentHandler.DeleteSafetyDocumentHandler)
	app.Get("/search", SafetyDocumentHandler.SearchSafetyDocumentHandler)

	return app
}
