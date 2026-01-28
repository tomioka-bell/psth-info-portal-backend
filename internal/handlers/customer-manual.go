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

type CustomerManualHandler struct {
	CustomerManualSrv services.CustomerManualService
}

func NewCustomerManualHandler(insSrv services.CustomerManualService) *CustomerManualHandler {
	return &CustomerManualHandler{CustomerManualSrv: insSrv}
}

// Create customer manual with single file upload
func (h *CustomerManualHandler) CreateCustomerManualHandler(c *fiber.Ctx) error {
	relPath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/customer_manual",
		AllowedMIMEs: []string{"application/pdf"},
		MaxSize:      50 << 20, // 50MB
		BaseURL:      "",
		Required:     true,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := models.CreateCustomerManualRequest{
		CustomerManualName: c.FormValue("customer_manual_name"),
		Desc:               c.FormValue("desc"),
		Category:           c.FormValue("category"),
		FileName:           fmt.Sprintf("[\"%s\"]", relPath), // Store as JSON array for consistency
	}

	if err := h.CustomerManualSrv.CreateCustomerManualService(req); err != nil {
		log.Println("Error creating customer manual:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create customer manual"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Customer manual created successfully",
		"file":    relPath,
	})
}

// Get all customer manuals
func (h *CustomerManualHandler) GetAllCustomerManualHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	manuals, err := h.CustomerManualSrv.GetAllCustomerManualService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch customer manuals",
		})
	}

	return c.Status(http.StatusOK).JSON(manuals)
}

// Update customer manual
func (h *CustomerManualHandler) UpdateCustomerManualHandler(c *fiber.Ctx) error {
	customerManualIDStr := c.Params("customer_manual_id")
	customerManualID, err := strconv.Atoi(customerManualIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer manual ID",
		})
	}

	log.Printf("[UpdateCustomerManualHandler] Received ID: %d\n", customerManualID)

	var relPath string
	// Handle optional file upload
	if c.Request().Header.Peek("Content-Type") != nil {
		log.Println("[UpdateCustomerManualHandler] Processing file upload...")
		relPath, _, err = uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/customer_manual",
			AllowedMIMEs: []string{"application/pdf"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err != nil {
			log.Printf("[UpdateCustomerManualHandler] File upload error: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		log.Printf("[UpdateCustomerManualHandler] File uploaded successfully: %s\n", relPath)
	}

	req := models.UpdateCustomerManualRequest{
		CustomerManualName: c.FormValue("customer_manual_name"),
		Desc:               c.FormValue("desc"),
		Category:           c.FormValue("category"),
	}

	// Only set FileName if a new file was uploaded
	if relPath != "" {
		req.FileName = fmt.Sprintf("[\"%s\"]", relPath)
	}

	log.Printf("[UpdateCustomerManualHandler] Calling UpdateCustomerManualService with ID: %d\n", customerManualID)
	if err := h.CustomerManualSrv.UpdateCustomerManualService(customerManualID, req); err != nil {
		log.Printf("[UpdateCustomerManualHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update customer manual"})
	}

	log.Println("[UpdateCustomerManualHandler] Update completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Customer manual updated successfully",
	})
}

// Delete customer manual
func (h *CustomerManualHandler) DeleteCustomerManualHandler(c *fiber.Ctx) error {
	customerManualIDStr := c.Params("customer_manual_id")
	customerManualID, err := strconv.Atoi(customerManualIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer manual ID",
		})
	}

	log.Printf("[DeleteCustomerManualHandler] Received ID: %d\n", customerManualID)

	if err := h.CustomerManualSrv.DeleteCustomerManualService(customerManualID); err != nil {
		log.Printf("[DeleteCustomerManualHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete customer manual"})
	}

	log.Println("[DeleteCustomerManualHandler] Delete completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Customer manual deleted successfully",
	})
}
