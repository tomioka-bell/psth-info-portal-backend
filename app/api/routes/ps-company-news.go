package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesCompanyNews(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	CompanyNewsRepository := repositories.NewCompanyNewsRepositoryDB(db)
	CompanyNewsService := services.NewCompanyNewsService(CompanyNewsRepository)
	CompanyNewsHandler := handlers.NewCompanyNewsHandler(CompanyNewsService)

	app.Post("/create-company-news", CompanyNewsHandler.CreateCompanyNewsFormHandler)
	app.Get("/get-company-news", CompanyNewsHandler.GetCompanyNewsHandler)
	app.Get("/get-company-news-by-title", CompanyNewsHandler.GetCompanyNewsByTitleHandler)
	app.Put("/update-company-news/:company_news_id", CompanyNewsHandler.UpdateCompanyNewsFormHandler)
	app.Delete("/delete-company-news/:company_news_id", CompanyNewsHandler.DeleteCompanyNewsHandler)
	return app
}
