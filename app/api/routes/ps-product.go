package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/core/services"
	"backend/internal/handlers"
	"backend/internal/repositories"
)

func RoutesProduct(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	ProductRepository := repositories.NewProductRepositoryDB(db)
	ProductService := services.NewProductService(ProductRepository)
	ProductHandler := handlers.NewProductHandler(ProductService)

	app.Post("/create-product", ProductHandler.CreateProductFormHandler)
	app.Get("/get-products", ProductHandler.GetAllProductsHandler)
	app.Get("/get-recommended-products", ProductHandler.GetRecommendedProductsHandler)
	app.Put("/update-product/:id", ProductHandler.UpdateProductHandler)
	app.Delete("/delete-product/:id", ProductHandler.DeleteProductHandler)
	app.Get("/get-recommended-product", ProductHandler.GetRecommendedProductHandler)
	app.Get("/search-products", ProductHandler.SearchProductsByNameHandler)

	return app
}
