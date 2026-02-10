package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
	uploader "backend/internal/pkgs/utils"
)

type WelfareBenefitHandler struct {
	WelfareBenefitSrv services.WelfareBenefitService
}

func NewWelfareBenefitHandler(insSrv services.WelfareBenefitService) *WelfareBenefitHandler {
	return &WelfareBenefitHandler{WelfareBenefitSrv: insSrv}
}

// Create welfare benefit with file and/or image upload
func (h *WelfareBenefitHandler) CreateWelfareBenefitHandler(c *fiber.Ctx) error {
	var fileRelPath, imageRelPath string

	// Upload file (PDF, DOC, etc.)
	filePath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/welfare_benefits",
		AllowedMIMEs: []string{"application/pdf", "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		MaxSize:      50 << 20, // 50MB
		BaseURL:      "",
		Required:     false,
	})
	if err == nil && filePath != "" {
		fileRelPath = filePath
	}

	// Upload image (JPEG, PNG, WebP)
	imgPath, _, err := uploader.UploadFromForm(c, "image", uploader.Options{
		Dir:          "./uploads/welfare_benefits",
		AllowedMIMEs: []string{"image/jpeg", "image/png", "image/webp"},
		MaxSize:      10 << 20, // 10MB
		BaseURL:      "",
		Required:     false,
	})
	if err == nil && imgPath != "" {
		imageRelPath = imgPath
	}

	if fileRelPath == "" && imageRelPath == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please upload at least one file (document or image)"})
	}

	req := models.CreateWelfareBenefitRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Category:    c.FormValue("category"),
	}

	// Set FileName if file was uploaded
	if fileRelPath != "" {
		req.FileName = fmt.Sprintf(`["%s"]`, fileRelPath)
	}

	// Set ImageURL if image was uploaded
	if imageRelPath != "" {
		req.ImageURL = imageRelPath
	}

	if err := h.WelfareBenefitSrv.CreateWelfareBenefitService(req); err != nil {
		log.Println("Error creating welfare benefit:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create welfare benefit"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Welfare benefit created successfully",
		"file":    fileRelPath,
		"image":   imageRelPath,
	})
}

// Get all welfare benefits
func (h *WelfareBenefitHandler) GetAllWelfareBenefitHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	benefits, err := h.WelfareBenefitSrv.GetAllWelfareBenefitService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch welfare benefits",
		})
	}

	return c.Status(http.StatusOK).JSON(benefits)
}

// Get total welfare benefits count
func (h *WelfareBenefitHandler) GetWelfareBenefitsCountHandler(c *fiber.Ctx) error {
	total, err := h.WelfareBenefitSrv.GetWelfareBenefitsCountService()
	if err != nil {
		log.Printf("[GetWelfareBenefitsCountHandler] error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get welfare benefits count"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"total": total})
}

// Get welfare benefits by category
func (h *WelfareBenefitHandler) GetWelfareBenefitByCategoryHandler(c *fiber.Ctx) error {
	category := c.Params("category")
	limit := c.QueryInt("limit", 250)
	offset := c.QueryInt("offset", 0)

	// If user requests "all-employee" (case-insensitive), return full list
	if strings.TrimSpace(strings.ToLower(category)) == "all-employee" {
		fmt.Println("เข้าเงื่อนไข")
		benefits, err := h.WelfareBenefitSrv.GetAllWelfareBenefitsNoPaginationService()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch welfare benefits",
			})
		}

		return c.Status(http.StatusOK).JSON(benefits)
	}

	benefits, err := h.WelfareBenefitSrv.GetWelfareBenefitByCategoryService(category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch welfare benefits",
		})
	}

	return c.Status(http.StatusOK).JSON(benefits)
}

// Update welfare benefit with optional file and/or image upload
func (h *WelfareBenefitHandler) UpdateWelfareBenefitHandler(c *fiber.Ctx) error {
	benefitIDStr := c.Params("welfare_benefit_id")
	benefitID, err := strconv.Atoi(benefitIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid welfare benefit ID",
		})
	}

	log.Printf("[UpdateWelfareBenefitHandler] Received ID: %d\n", benefitID)

	var fileRelPath, imageRelPath string

	// Handle optional file upload
	if c.Request().Header.Peek("Content-Type") != nil {
		log.Println("[UpdateWelfareBenefitHandler] Processing file uploads...")

		// Upload file (PDF, DOC, etc.)
		filePath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/welfare_benefits",
			AllowedMIMEs: []string{"application/pdf", "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err == nil && filePath != "" {
			fileRelPath = filePath
			log.Printf("[UpdateWelfareBenefitHandler] File uploaded: %s\n", fileRelPath)
		}

		// Upload image (JPEG, PNG, WebP)
		imgPath, _, err := uploader.UploadFromForm(c, "image", uploader.Options{
			Dir:          "./uploads/welfare_benefits",
			AllowedMIMEs: []string{"image/jpeg", "image/png", "image/webp"},
			MaxSize:      10 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err == nil && imgPath != "" {
			imageRelPath = imgPath
			log.Printf("[UpdateWelfareBenefitHandler] Image uploaded: %s\n", imageRelPath)
		}
	}

	req := models.UpdateWelfareBenefitRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Category:    c.FormValue("category"),
	}

	// Set FileName if new file was uploaded
	if fileRelPath != "" {
		req.FileName = fmt.Sprintf(`["%s"]`, fileRelPath)
	}

	// Set ImageURL if new image was uploaded
	if imageRelPath != "" {
		req.ImageURL = imageRelPath
	}

	log.Printf("[UpdateWelfareBenefitHandler] Calling UpdateWelfareBenefitService with ID: %d\n", benefitID)
	if err := h.WelfareBenefitSrv.UpdateWelfareBenefitService(benefitID, req); err != nil {
		log.Printf("[UpdateWelfareBenefitHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update welfare benefit"})
	}

	log.Println("[UpdateWelfareBenefitHandler] Update completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Welfare benefit updated successfully",
	})
}

// Delete welfare benefit
func (h *WelfareBenefitHandler) DeleteWelfareBenefitHandler(c *fiber.Ctx) error {
	benefitIDStr := c.Params("welfare_benefit_id")
	benefitID, err := strconv.Atoi(benefitIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid welfare benefit ID",
		})
	}

	log.Printf("[DeleteWelfareBenefitHandler] Received ID: %d\n", benefitID)

	if err := h.WelfareBenefitSrv.DeleteWelfareBenefitService(benefitID); err != nil {
		log.Printf("[DeleteWelfareBenefitHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete welfare benefit"})
	}

	log.Println("[DeleteWelfareBenefitHandler] Delete completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Welfare benefit deleted successfully",
	})
}

// Search welfare benefits
func (h *WelfareBenefitHandler) SearchWelfareBenefitHandler(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "keyword is required",
		})
	}

	benefits, err := h.WelfareBenefitSrv.SearchWelfareBenefitService(keyword, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search welfare benefits",
		})
	}

	return c.Status(http.StatusOK).JSON(benefits)
}
