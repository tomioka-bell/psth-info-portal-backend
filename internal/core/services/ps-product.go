package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	ports "backend/internal/core/ports/repositories"
	servicesports "backend/internal/core/ports/services"
	"backend/internal/pkgs/logs"
)

type productService struct {
	productRepo ports.ProductRepository
}

func NewProductService(productRepo ports.ProductRepository) servicesports.ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProductService(req models.CreateProductRequest) error {

	newID := uuid.New()

	domainProduct := domains.Product{
		ProductID:         newID.String(),
		ProductName:       req.ProductName,
		Category:          req.Category,
		Description:       req.Description,
		ProductMainImages: req.ProductMainImages,
		ProductImages:     req.ProductImages,
		Recommend:         req.Recommend,
	}

	if s.productRepo == nil {
		log.Println("ProductRepo is nil")
		return fmt.Errorf("product repository is not initialized")
	}

	err := s.productRepo.CreateProduct(&domainProduct)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (s *productService) GetAllProductsService(limit int) ([]models.ProductResponse, error) {
	domainProducts, err := s.productRepo.GetAllProducts(limit)
	if err != nil {
		logs.Error(err)
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	var productResponses []models.ProductResponse
	for _, dp := range domainProducts {
		productResponse := models.ProductResponse{
			ProductID:         dp.ProductID,
			ProductName:       dp.ProductName,
			Category:          dp.Category,
			Description:       dp.Description,
			ProductMainImages: dp.ProductMainImages,
			ProductImages:     dp.ProductImages,
			Recommend:         dp.Recommend,
			CreatedAt:         dp.CreatedAt,
			UpdatedAt:         dp.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (s *productService) GetRecommendedProductService(limit int) ([]models.ProductResponse, error) {
	domainProducts, err := s.productRepo.GetRecommendedProducts(limit)
	if err != nil {
		logs.Error(err)
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	var productResponses []models.ProductResponse
	for _, dp := range domainProducts {
		productResponse := models.ProductResponse{
			ProductID:         dp.ProductID,
			ProductName:       dp.ProductName,
			Category:          dp.Category,
			Description:       dp.Description,
			ProductMainImages: dp.ProductMainImages,
			// ProductImages:     dp.ProductImages,
			Recommend: dp.Recommend,
			CreatedAt: dp.CreatedAt,
			UpdatedAt: dp.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (s *productService) UpdateProductService(productID string, req models.UpdateProductRequest) error {
	if s.productRepo == nil {
		log.Println("ProductRepo is nil")
		return fmt.Errorf("product repository is not initialized")
	}

	// Get existing product
	existingProduct, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("product not found: %w", err)
	}

	// Update basic fields
	if req.ProductName != "" {
		existingProduct.ProductName = req.ProductName
	}
	if req.Category != "" {
		existingProduct.Category = req.Category
	}
	if req.Description != "" {
		existingProduct.Description = req.Description
	}

	// Update recommend only if provided
	if req.Recommend != nil {
		existingProduct.Recommend = *req.Recommend
	}

	// Handle main image update (but not delete)
	if req.ProductMainImages != "" {
		existingProduct.ProductMainImages = req.ProductMainImages
	}

	// Handle product images - delete selected and add new ones
	if len(req.ImagesToDelete) > 0 || req.ProductImages != "" {
		// Parse existing images
		var existingImages []string
		if existingProduct.ProductImages != "" {
			if err := json.Unmarshal([]byte(existingProduct.ProductImages), &existingImages); err != nil {
				logs.Error(err)
				existingImages = []string{}
			}
		}

		// Create a map of images to delete for quick lookup
		deleteMap := make(map[string]bool)
		for _, img := range req.ImagesToDelete {
			deleteMap[img] = true
		}

		// Filter out deleted images
		var filteredImages []string
		for _, img := range existingImages {
			if !deleteMap[img] {
				filteredImages = append(filteredImages, img)
			}
		}

		// Add new images
		if req.ProductImages != "" {
			var newImages []string
			if err := json.Unmarshal([]byte(req.ProductImages), &newImages); err == nil {
				filteredImages = append(filteredImages, newImages...)
			}
		}

		// Marshal back to JSON
		updatedImagesJSON, err := json.Marshal(filteredImages)
		if err != nil {
			logs.Error(err)
			return fmt.Errorf("failed to marshal images: %w", err)
		}
		existingProduct.ProductImages = string(updatedImagesJSON)
	}

	err = s.productRepo.UpdateProduct(productID, existingProduct)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (s *productService) DeleteProductService(productID string) error {
	if s.productRepo == nil {
		log.Println("ProductRepo is nil")
		return fmt.Errorf("product repository is not initialized")
	}

	err := s.productRepo.DeleteProduct(productID)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (s *productService) SearchProductsByName(keyword string, limit int) ([]models.ProductResponse, error) {
	domainProducts, err := s.productRepo.SearchProductsByName(keyword, limit)
	if err != nil {
		logs.Error(err)
		return nil, fmt.Errorf("failed to search products: %w", err)
	}

	var productResponses []models.ProductResponse
	for _, dp := range domainProducts {
		productResponse := models.ProductResponse{
			ProductID:         dp.ProductID,
			ProductName:       dp.ProductName,
			Category:          dp.Category,
			Description:       dp.Description,
			ProductMainImages: dp.ProductMainImages,
			ProductImages:     dp.ProductImages,
			Recommend:         dp.Recommend,
			CreatedAt:         dp.CreatedAt,
			UpdatedAt:         dp.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
