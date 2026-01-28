package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesProcedureManual(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	ProcedureManualRepository := repositories.NewProcedureManualRepository(db)
	ProcedureManualService := services.NewProcedureManualService(ProcedureManualRepository)
	ProcedureManualHandler := handlers.NewProcedureManualHandler(ProcedureManualService)

	app.Post("/create", ProcedureManualHandler.CreateProcedureManualHandler)
	app.Get("/list", ProcedureManualHandler.GetAllProcedureManualHandler)
	app.Put("/update/:procedure_manual_id", ProcedureManualHandler.UpdateProcedureManualHandler)
	app.Delete("/delete/:procedure_manual_id", ProcedureManualHandler.DeleteProcedureManualHandler)

	return app
}
