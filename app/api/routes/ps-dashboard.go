package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesDashboard(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	DashboardRepository := repositories.NewDashboardRepository(db)
	DashboardService := services.NewDashboardService(DashboardRepository)
	DashboardHandler := handlers.NewDashboardHandler(DashboardService)

	app.Get("/stats", DashboardHandler.GetDashboardStats)
	app.Get("/table/:table", DashboardHandler.GetTableStats)
	app.Get("/counts", DashboardHandler.GetAllCounts)
	app.Get("/category/app-systems", DashboardHandler.GetAppSystemsCategory)
	app.Get("/category/organizations", DashboardHandler.GetOrganizationsCategory)
	app.Get("/category/safety-documents", DashboardHandler.GetSafetyDocsCategory)
	app.Get("/category/safety-documents-department", DashboardHandler.GetSafetyDocsDepartment)
	app.Get("/category/welfare-benefits", DashboardHandler.GetWelfareBenefitsCategory)
	app.Get("/category/company-news", DashboardHandler.GetNewsCategory)

	return app
}
