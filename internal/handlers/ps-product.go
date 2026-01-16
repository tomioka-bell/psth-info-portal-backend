package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
)

type ProductHandler struct {
	ProductSrv services.ProductService
}

func NewProductHandler(insSrv services.ProductService) *ProductHandler {
	return &ProductHandler{ProductSrv: insSrv}
}

func (h *ProductHandler) CreateProductFormHandler(c *fiber.Ctx) error {
	productName := c.FormValue("product_name")
	category := c.FormValue("category")
	description := c.FormValue("description")
	recommendStr := c.FormValue("recommend")

	if productName == "" || category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_name and category are required",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	// ตั้งค่า
	uploadDir := "./uploads/product"
	allowedMIMEs := []string{"image/jpeg", "image/png", "image/webp"}
	maxFileSize := int64(10 << 20) // 10MB

	// รับภาพหลัก
	mainImageFiles := form.File["main_image"]
	mainImage := ""
	if len(mainImageFiles) > 0 {
		mainImageHeader := mainImageFiles[0]

		file, err := mainImageHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read main image file",
			})
		}
		defer file.Close()

		buf := make([]byte, 512)
		n, _ := io.ReadFull(file, buf)
		detected := http.DetectContentType(buf[:n])

		// ตรวจสอบ MIME type
		isAllowed := false
		for _, m := range allowedMIMEs {
			if strings.EqualFold(m, detected) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid main image type. Only JPEG, PNG, WebP are allowed",
			})
		}

		// ตรวจสอบขนาด
		if mainImageHeader.Size > maxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Main image too large. Maximum size is 10MB",
			})
		}

		file.Seek(0, 0)

		os.MkdirAll(uploadDir, 0755)

		randomStr := make([]byte, 6)
		rand.Read(randomStr)
		randomHex := hex.EncodeToString(randomStr)
		ext := filepath.Ext(mainImageHeader.Filename)
		filename := time.Now().Format("20060102_150405") + "_" + randomHex + ext

		filePath := filepath.Join(uploadDir, filename)
		dst, err := os.Create(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save main image",
			})
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to write main image",
			})
		}

		mainImage = filePath
	}

	// รับภาพอื่น ๆ
	imageFiles := form.File["images"]

	if len(imageFiles) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "At least one image is required",
		})
	}

	var imagePaths []string

	for _, fileHeader := range imageFiles {
		if fileHeader.Size > maxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "File too large. Maximum size is 10MB",
			})
		}

		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read file",
			})
		}
		defer file.Close()

		buf := make([]byte, 512)
		n, _ := io.ReadFull(file, buf)
		detected := http.DetectContentType(buf[:n])

		isAllowed := false
		for _, m := range allowedMIMEs {
			if strings.EqualFold(m, detected) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid file type. Only JPEG, PNG, WebP are allowed",
			})
		}
		file.Seek(0, 0)

		os.MkdirAll(uploadDir, 0755)

		randomStr := make([]byte, 6)
		rand.Read(randomStr)
		randomHex := hex.EncodeToString(randomStr)
		ext := filepath.Ext(fileHeader.Filename)
		filename := time.Now().Format("20060102_150405") + "_" + randomHex + ext

		filePath := filepath.Join(uploadDir, filename)
		dst, err := os.Create(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to write file",
			})
		}

		imagePaths = append(imagePaths, filePath)
	}

	imagePathsJSON, err := json.Marshal(imagePaths)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal image paths",
		})
	}

	req := models.CreateProductRequest{
		ProductName:       productName,
		Category:          category,
		Description:       description,
		ProductMainImages: mainImage,
		ProductImages:     string(imagePathsJSON),
		Recommend:         recommendStr == "true",
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	err = h.ProductSrv.CreateProductService(req)
	if err != nil {
		log.Println("Error creating product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) GetAllProductsHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 100)

	if limit <= 0 {
		limit = 100
	}

	if limit > 100 {
		limit = 100
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	products, err := h.ProductSrv.GetAllProductsService(limit)
	if err != nil {
		log.Println("Error getting products:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func (h *ProductHandler) UpdateProductHandler(c *fiber.Ctx) error {
	productID := c.Params("id")
	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_id is required",
		})
	}

	productName := c.FormValue("product_name")
	category := c.FormValue("category")
	description := c.FormValue("description")
	recommendStr := c.FormValue("recommend")

	if productName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_name is required",
		})
	}

	// ตั้งค่า
	uploadDir := "./uploads/product"
	allowedMIMEs := []string{"image/jpeg", "image/png", "image/webp"}
	maxFileSize := int64(10 << 20) // 10MB

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	mainImage := ""
	// รับภาพหลัก (ถ้ามี)
	mainImageFiles := form.File["main_image"]
	if len(mainImageFiles) > 0 {
		mainImageHeader := mainImageFiles[0]

		file, err := mainImageHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read main image file",
			})
		}
		defer file.Close()

		buf := make([]byte, 512)
		n, _ := io.ReadFull(file, buf)
		detected := http.DetectContentType(buf[:n])

		isAllowed := false
		for _, m := range allowedMIMEs {
			if strings.EqualFold(m, detected) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid main image type. Only JPEG, PNG, WebP are allowed",
			})
		}

		if mainImageHeader.Size > maxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Main image too large. Maximum size is 10MB",
			})
		}

		file.Seek(0, 0)

		os.MkdirAll(uploadDir, 0755)

		randomStr := make([]byte, 6)
		rand.Read(randomStr)
		randomHex := hex.EncodeToString(randomStr)
		ext := filepath.Ext(mainImageHeader.Filename)
		filename := time.Now().Format("20060102_150405") + "_" + randomHex + ext

		filePath := filepath.Join(uploadDir, filename)
		dst, err := os.Create(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save main image",
			})
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to write main image",
			})
		}

		mainImage = filePath
	}

	// รับภาพอื่น ๆ (ถ้ามี)
	imageFiles := form.File["images"]
	productImages := ""

	if len(imageFiles) > 0 {
		var imagePaths []string

		for _, fileHeader := range imageFiles {
			if fileHeader.Size > maxFileSize {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "File too large. Maximum size is 10MB",
				})
			}

			file, err := fileHeader.Open()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to read file",
				})
			}
			defer file.Close()

			buf := make([]byte, 512)
			n, _ := io.ReadFull(file, buf)
			detected := http.DetectContentType(buf[:n])

			isAllowed := false
			for _, m := range allowedMIMEs {
				if strings.EqualFold(m, detected) {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid file type. Only JPEG, PNG, WebP are allowed",
				})
			}
			file.Seek(0, 0)

			os.MkdirAll(uploadDir, 0755)

			randomStr := make([]byte, 6)
			rand.Read(randomStr)
			randomHex := hex.EncodeToString(randomStr)
			ext := filepath.Ext(fileHeader.Filename)
			filename := time.Now().Format("20060102_150405") + "_" + randomHex + ext

			filePath := filepath.Join(uploadDir, filename)
			dst, err := os.Create(filePath)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to save file",
				})
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to write file",
				})
			}

			imagePaths = append(imagePaths, filePath)
		}

		imagePathsJSON, err := json.Marshal(imagePaths)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to marshal image paths",
			})
		}
		productImages = string(imagePathsJSON)
	}

	req := models.UpdateProductRequest{
		ProductName:       productName,
		Category:          category,
		Description:       description,
		ProductMainImages: mainImage,
		ProductImages:     productImages,
	}

	// Handle recommend as pointer
	recommend := recommendStr == "true"
	req.Recommend = &recommend

	// Handle image deletion
	imagesToDeleteStr := c.FormValue("images_to_delete")
	if imagesToDeleteStr != "" {
		var imagesToDelete []string
		if err := json.Unmarshal([]byte(imagesToDeleteStr), &imagesToDelete); err == nil {
			req.ImagesToDelete = imagesToDelete
		}
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	err = h.ProductSrv.UpdateProductService(productID, req)
	if err != nil {
		log.Println("Error updating product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h *ProductHandler) DeleteProductHandler(c *fiber.Ctx) error {
	productID := c.Params("id")
	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_id is required",
		})
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	err := h.ProductSrv.DeleteProductService(productID)
	if err != nil {
		log.Println("Error deleting product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func (h *ProductHandler) GetRecommendedProductsHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 100)

	if limit <= 0 {
		limit = 100
	}

	if limit > 100 {
		limit = 100
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	products, err := h.ProductSrv.GetAllProductsService(limit)
	if err != nil {
		log.Println("Error getting recommended products:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve recommended products",
		})
	}

	// Filter only recommended products
	var recommendedProducts []models.ProductResponse
	for _, product := range products {
		if product.Recommend {
			recommendedProducts = append(recommendedProducts, product)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Recommended products retrieved successfully",
		"data":    recommendedProducts,
	})
}

func (h *ProductHandler) GetRecommendedProductHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 100)

	if limit <= 0 {
		limit = 100
	}

	if limit > 100 {
		limit = 100
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	products, err := h.ProductSrv.GetRecommendedProductService(limit)
	if err != nil {
		log.Println("Error getting recommended products:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve recommended products",
		})
	}

	var recommendedProducts []models.ProductResponse
	for _, product := range products {
		if product.Recommend {
			recommendedProducts = append(recommendedProducts, product)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Recommended products retrieved successfully",
		"data":    recommendedProducts,
	})
}

func (h *ProductHandler) SearchProductsByNameHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 100)

	if limit <= 0 {
		limit = 100
	}

	if limit > 100 {
		limit = 100
	}

	if h.ProductSrv == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Service is not available",
		})
	}

	products, err := h.ProductSrv.SearchProductsByName(c.Query("product_name", ""), limit)
	if err != nil {
		log.Println("Error searching products:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Products retrieved successfully",
		"data":    products,
	})
}
