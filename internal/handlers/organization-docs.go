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

type OrganizationDocHandler struct {
	OrganizationDocSrv services.OrganizationDocService
}

func NewOrganizationDocHandler(insSrv services.OrganizationDocService) *OrganizationDocHandler {
	return &OrganizationDocHandler{OrganizationDocSrv: insSrv}
}

// Create organization doc with single file upload
func (h *OrganizationDocHandler) CreateOrganizationDocHandler(c *fiber.Ctx) error {
	relPath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/organization_docs",
		AllowedMIMEs: []string{"application/pdf"},
		MaxSize:      50 << 20, // 50MB
		BaseURL:      "",
		Required:     true,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := models.CreateOrganizationDocRequest{
		Name:       c.FormValue("name"),
		Desc:       c.FormValue("desc"),
		Department: c.FormValue("department"),
		FileName:   fmt.Sprintf("[\"%s\"]", relPath),
	}

	if err := h.OrganizationDocSrv.CreateOrganizationDocService(req); err != nil {
		log.Println("Error creating organization doc:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create organization doc"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Organization doc created successfully",
		"file":    relPath,
	})
}

// Get all organization docs
func (h *OrganizationDocHandler) GetAllOrganizationDocHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.OrganizationDocSrv.GetAllOrganizationDocService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch organization docs",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

// Get organization docs by department
func (h *OrganizationDocHandler) GetOrganizationDocByDepartmentHandler(c *fiber.Ctx) error {
	department := c.Params("department")
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.OrganizationDocSrv.GetOrganizationDocByDepartmentService(department, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch organization docs",
		})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

// Update organization doc
func (h *OrganizationDocHandler) UpdateOrganizationDocHandler(c *fiber.Ctx) error {
	docIDStr := c.Params("organization_doc_id")
	docID, err := strconv.Atoi(docIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid organization doc ID",
		})
	}

	log.Printf("[UpdateOrganizationDocHandler] Received ID: %d\n", docID)

	var relPath string
	// Handle optional file upload
	if c.Request().Header.Peek("Content-Type") != nil {
		log.Println("[UpdateOrganizationDocHandler] Processing file upload...")
		relPath, _, err = uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/organization_docs",
			AllowedMIMEs: []string{"application/pdf"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err != nil {
			log.Printf("[UpdateOrganizationDocHandler] File upload error: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		log.Printf("[UpdateOrganizationDocHandler] File uploaded successfully: %s\n", relPath)
	}

	req := models.UpdateOrganizationDocRequest{
		Name:       c.FormValue("name"),
		Desc:       c.FormValue("desc"),
		Department: c.FormValue("department"),
	}

	// Only set FileName if a new file was uploaded
	if relPath != "" {
		req.FileName = fmt.Sprintf("[\"%s\"]", relPath)
	}

	log.Printf("[UpdateOrganizationDocHandler] Calling UpdateOrganizationDocService with ID: %d\n", docID)
	if err := h.OrganizationDocSrv.UpdateOrganizationDocService(docID, req); err != nil {
		log.Printf("[UpdateOrganizationDocHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update organization doc"})
	}

	log.Println("[UpdateOrganizationDocHandler] Update completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Organization doc updated successfully",
	})
}

// Delete organization doc
func (h *OrganizationDocHandler) DeleteOrganizationDocHandler(c *fiber.Ctx) error {
	docIDStr := c.Params("organization_doc_id")
	docID, err := strconv.Atoi(docIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid organization doc ID",
		})
	}

	log.Printf("[DeleteOrganizationDocHandler] Received ID: %d\n", docID)

	if err := h.OrganizationDocSrv.DeleteOrganizationDocService(docID); err != nil {
		log.Printf("[DeleteOrganizationDocHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete organization doc"})
	}

	log.Println("[DeleteOrganizationDocHandler] Delete completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Organization doc deleted successfully",
	})
}
