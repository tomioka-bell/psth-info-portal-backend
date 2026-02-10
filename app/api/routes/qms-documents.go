package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesQmsDocuments(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	repo := repositories.NewQmsDocumentsRepository(db)
	srv := services.NewQmsDocumentsService(repo)
	h := handlers.NewQmsDocumentsHandler(srv)

	app.Post("/create", h.CreateQmsDocumentHandler)
	app.Get("/list", h.GetAllQmsDocumentsHandler)
	app.Put("/update/:qms_documents_id", h.UpdateQmsDocumentHandler)
	app.Delete("/delete/:qms_documents_id", h.DeleteQmsDocumentHandler)

	return app
}
