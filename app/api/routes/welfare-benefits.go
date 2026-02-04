package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesWelfareBenefit(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	WelfareBenefitRepository := repositories.NewWelfareBenefitRepository(db)
	WelfareBenefitService := services.NewWelfareBenefitService(WelfareBenefitRepository)
	WelfareBenefitHandler := handlers.NewWelfareBenefitHandler(WelfareBenefitService)

	app.Post("/create", WelfareBenefitHandler.CreateWelfareBenefitHandler)
	app.Get("/list", WelfareBenefitHandler.GetAllWelfareBenefitHandler)
	app.Get("/category/:category", WelfareBenefitHandler.GetWelfareBenefitByCategoryHandler)
	app.Put("/update/:welfare_benefit_id", WelfareBenefitHandler.UpdateWelfareBenefitHandler)
	app.Delete("/delete/:welfare_benefit_id", WelfareBenefitHandler.DeleteWelfareBenefitHandler)
	app.Get("/search", WelfareBenefitHandler.SearchWelfareBenefitHandler)

	return app

}
