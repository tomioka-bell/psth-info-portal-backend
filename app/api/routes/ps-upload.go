package api

import (
	"github.com/gofiber/fiber/v2"

	"backend/internal/handlers"
)

func RoutesFile() *fiber.App {
	app := fiber.New()

	app.Get("/get-file", handlers.ServeUploadFile)

	return app
}
