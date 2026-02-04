package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesAppSystem(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	AppSystemRepository := repositories.NewAppSystemRepository(db)
	AppSystemService := services.NewAppSystemService(AppSystemRepository)
	AppSystemHandler := handlers.NewAppSystemHandler(AppSystemService)

	// CRUD routes
	app.Post("/create", AppSystemHandler.CreateAppSystemHandler)
	app.Get("/list", AppSystemHandler.GetAllAppSystemsHandler)
	app.Get("/tree", AppSystemHandler.GetAppSystemsTreeHandler)
	app.Get("/category/:category", AppSystemHandler.GetAppSystemsByCategoryHandler)
	app.Get("/search", AppSystemHandler.SearchAppSystemsHandler)
	app.Get("/:id", AppSystemHandler.GetAppSystemHandler)
	app.Put("/update/:id", AppSystemHandler.UpdateAppSystemHandler)
	app.Delete("/:id", AppSystemHandler.DeleteAppSystemHandler)

	app.Post("/create-multiple", AppSystemHandler.CreateAppSystemsHandler)

	return app
}
