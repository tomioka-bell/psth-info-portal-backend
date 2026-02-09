package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
	uploader "backend/internal/pkgs/utils"
)

type SafetyDocumentHandler struct {
	SafetyDocumentSrv services.SafetyDocumentService
}

func NewSafetyDocumentHandler(insSrv services.SafetyDocumentService) *SafetyDocumentHandler {
	return &SafetyDocumentHandler{SafetyDocumentSrv: insSrv}
}

// Create safety document with single file upload
func (h *SafetyDocumentHandler) CreateSafetyDocumentHandler(c *fiber.Ctx) error {
	relPath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/safety_documents",
		AllowedMIMEs: []string{"application/pdf"},
		MaxSize:      50 << 20,
		BaseURL:      "",
		Required:     true,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := models.CreateSafetyDocumentRequest{
		SafetyDocumentName: c.FormValue("safety_document_name"),
		SafetyDocumentDesc: c.FormValue("safety_document_desc"),
		Category:           c.FormValue("category"),
		Department:         c.FormValue("department"),
		FileName:           fmt.Sprintf("[\"%s\"]", relPath),
	}

	if err := h.SafetyDocumentSrv.CreateSafetyDocumentService(req); err != nil {
		log.Println("Error creating safety document:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create safety document"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Safety document created successfully",
		"file":    relPath,
	})
}

// Get all safety documents
func (h *SafetyDocumentHandler) GetAllSafetyDocumentHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.SafetyDocumentSrv.GetAllSafetyDocumentService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch safety documents",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

// Get safety documents by category
func (h *SafetyDocumentHandler) GetSafetyDocumentByCategoryHandler(c *fiber.Ctx) error {
	category := c.Params("category")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.SafetyDocumentSrv.GetSafetyDocumentByCategoryService(category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch safety documents",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

// Get safety documents by department
func (h *SafetyDocumentHandler) GetSafetyDocumentByDepartmentHandler(c *fiber.Ctx) error {
	department := c.Params("department")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.SafetyDocumentSrv.GetSafetyDocumentByDepartmentService(department, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch safety documents",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

// Update safety document
func (h *SafetyDocumentHandler) UpdateSafetyDocumentHandler(c *fiber.Ctx) error {
	docIDStr := c.Params("safety_document_id")
	docID, err := strconv.Atoi(docIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid safety document ID",
		})
	}

	log.Printf("[UpdateSafetyDocumentHandler] Received ID: %d\n", docID)

	var relPath string
	// Handle optional file upload
	if c.Request().Header.Peek("Content-Type") != nil {
		log.Println("[UpdateSafetyDocumentHandler] Processing file upload...")
		relPath, _, err = uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/safety_documents",
			AllowedMIMEs: []string{"application/pdf"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err != nil {
			log.Printf("[UpdateSafetyDocumentHandler] File upload error: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		log.Printf("[UpdateSafetyDocumentHandler] File uploaded successfully: %s\n", relPath)
	}

	req := models.UpdateSafetyDocumentRequest{
		SafetyDocumentName: c.FormValue("safety_document_name"),
		SafetyDocumentDesc: c.FormValue("safety_document_desc"),
		Category:           c.FormValue("category"),
		Department:         c.FormValue("department"),
	}

	// Only set FileName if a new file was uploaded
	if relPath != "" {
		req.FileName = fmt.Sprintf("[\"%s\"]", relPath)
	}

	log.Printf("[UpdateSafetyDocumentHandler] Calling UpdateSafetyDocumentService with ID: %d\n", docID)
	if err := h.SafetyDocumentSrv.UpdateSafetyDocumentService(docID, req); err != nil {
		log.Printf("[UpdateSafetyDocumentHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update safety document"})
	}

	log.Println("[UpdateSafetyDocumentHandler] Update completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Safety document updated successfully",
	})
}

// Delete safety document
func (h *SafetyDocumentHandler) DeleteSafetyDocumentHandler(c *fiber.Ctx) error {
	docIDStr := c.Params("safety_document_id")
	docID, err := strconv.Atoi(docIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid safety document ID",
		})
	}

	log.Printf("[DeleteSafetyDocumentHandler] Received ID: %d\n", docID)

	if err := h.SafetyDocumentSrv.DeleteSafetyDocumentService(docID); err != nil {
		log.Printf("[DeleteSafetyDocumentHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete safety document"})
	}

	log.Println("[DeleteSafetyDocumentHandler] Delete completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Safety document deleted successfully",
	})
}

// Search safety documents
func (h *SafetyDocumentHandler) SearchSafetyDocumentHandler(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "keyword is required",
		})
	}

	docs, err := h.SafetyDocumentSrv.SearchSafetyDocumentService(keyword, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search safety documents",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}
