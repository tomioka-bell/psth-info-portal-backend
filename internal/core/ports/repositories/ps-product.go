package ports

import (
	"backend/internal/core/domains"
)

type ProductRepository interface {
	CreateProduct(n *domains.Product) error
	GetAllProducts(limit int) ([]domains.Product, error)
	UpdateProduct(productID string, product *domains.Product) error
	DeleteProduct(productID string) error
	GetProductByID(productID string) (*domains.Product, error)
	GetRecommendedProducts(limit int) ([]domains.Product, error)
	SearchProductsByName(keyword string, limit int) ([]domains.Product, error)
}
