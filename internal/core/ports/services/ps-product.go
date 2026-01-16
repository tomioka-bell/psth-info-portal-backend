package ports

import "backend/internal/core/models"

type ProductService interface {
	CreateProductService(req models.CreateProductRequest) error
	GetAllProductsService(limit int) ([]models.ProductResponse, error)
	UpdateProductService(productID string, req models.UpdateProductRequest) error
	DeleteProductService(productID string) error
	GetRecommendedProductService(limit int) ([]models.ProductResponse, error)
	SearchProductsByName(keyword string, limit int) ([]models.ProductResponse, error)
}
