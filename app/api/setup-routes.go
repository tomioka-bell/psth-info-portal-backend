package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	routes "backend/app/api/routes"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	if db == nil {
		panic("Database connection is nil")
	}

	api := app.Group("/api", logger.New())
	api.Mount("/user", routes.RoutesUser(db))
	api.Mount("/product", routes.RoutesProduct(db))
	api.Mount("/company-news", routes.RoutesCompanyNews(db))
	api.Mount("/file", routes.RoutesFile())
	api.Mount("/organization", routes.RoutesOrganization(db))
	api.Mount("/app-system", routes.RoutesAppSystem(db))
	api.Mount("/procedure-manual", routes.RoutesProcedureManual(db))
	api.Mount("/customer-manual", routes.RoutesCustomerManual(db))
	api.Mount("/organization-docs", routes.RoutesOrganizationDoc(db))
	api.Mount("/questionnaire", routes.RoutesQuestionnaire(db))
}
