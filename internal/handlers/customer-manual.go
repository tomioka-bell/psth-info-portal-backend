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

func (h *CustomerManualHandler) CreateCustomerManualNestedHandler(c *fiber.Ctx) error {
	var nestedReq []models.CreateCustomerManualRequest

	if err := c.BodyParser(&nestedReq); err != nil {
		var singleReq models.CreateCustomerManualRequest
		if err := c.BodyParser(&singleReq); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		nestedReq = []models.CreateCustomerManualRequest{singleReq}
	}

	if len(nestedReq) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Request body is empty",
		})
	}

	type CreatedManual struct {
		ID               int             `json:"id"`
		CustomerManualID int             `json:"customer_manual_id"`
		Name             string          `json:"name"`
		Children         []CreatedManual `json:"children,omitempty"`
	}

	var createdItems []CreatedManual

	// Recursive function สร้าง CustomerManual + children
	var createRecursive func(req models.CreateCustomerManualRequest, parentID *int) (*CreatedManual, error)
	createRecursive = func(req models.CreateCustomerManualRequest, parentID *int) (*CreatedManual, error) {
		createReq := models.CreateCustomerManualRequest{
			CustomerManualName: req.CustomerManualName,
			Desc:               req.Desc,
			Category:           req.Category,
			FileName:           req.FileName,
			ParentID:           parentID,
			SortOrder:          req.SortOrder,
		}

		// สร้าง parent item
		resp, err := h.CustomerManualSrv.CreateCustomerManualWithResponse(createReq)
		if err != nil {
			return nil, err
		}

		created := &CreatedManual{
			ID:               resp.CustomerManualID,
			CustomerManualID: resp.CustomerManualID,
			Name:             resp.CustomerManualName,
		}

		// สร้าง children recursively ด้วย parent ID ที่เพิ่งสร้าง
		for _, child := range req.Children {
			childCreated, err := createRecursive(child, &resp.CustomerManualID)
			if err != nil {
				return nil, err
			}
			created.Children = append(created.Children, *childCreated)
		}

		return created, nil
	}

	// สร้างทั้งหมด
	for _, item := range nestedReq {
		created, err := createRecursive(item, nil)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		createdItems = append(createdItems, *created)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Customer manuals created successfully",
		"data":    createdItems,
	})
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

// Get customer manuals as tree structure
func (h *CustomerManualHandler) GetCustomerManualTreeHandler(c *fiber.Ctx) error {
	tree, err := h.CustomerManualSrv.GetAllCustomerManualsTree()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch customer manuals tree",
		})
	}

	return c.Status(http.StatusOK).JSON(tree)
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

	sortOrder := 0
	if sortOrderStr := c.FormValue("sort_order"); sortOrderStr != "" {
		if so, err := strconv.Atoi(sortOrderStr); err == nil {
			sortOrder = so
		}
	}

	var parentID *int
	clearParentID := false
	if parentIDStr := c.FormValue("parent_id"); parentIDStr != "" {
		if pid, err := strconv.Atoi(parentIDStr); err == nil {
			parentID = &pid
		}
	} else if parentIDStr == "" && c.FormValue("parent_id") == "" {
		// Empty string means user wants to clear parent_id (set to NULL)
		clearParentID = true
	}

	req := models.UpdateCustomerManualRequest{
		CustomerManualName: c.FormValue("customer_manual_name"),
		Desc:               c.FormValue("desc"),
		Category:           c.FormValue("category"),
		SortOrder:          sortOrder,
		ParentID:           parentID,
		ClearParentID:      clearParentID,
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
