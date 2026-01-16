package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesUser(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	UserRepository := repositories.NewUserRepositoryDB(db)
	UserService := services.NewUserService(UserRepository)
	UserHandler := handlers.NewUserHandler(UserService)

	app.Post("/create-user", UserHandler.CreateUserHandler)
	app.Post("/sign-out", UserHandler.LogoutDBHandler)
	app.Get("/get-by-id/:user_id", UserHandler.GetUserByIDHandler)
	app.Get("/get-all-user", UserHandler.GetAllUserHandler)
	app.Patch("/update-user/:user_id", UserHandler.UpdateUserHandler)

	app.Post("/login", UserHandler.LoginHandler)
	app.Get("/get-by-profile", UserHandler.GetProfileHandler)
	app.Get("/get-user-count", UserHandler.GetUserCountHandler)

	app.Get("/get-all-employees", UserHandler.GetAllEmployeesHandler)
	return app
}
