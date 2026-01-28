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

type ProcedureManualHandler struct {
	ProcedureManualSrv services.ProcedureManualService
}

func NewProcedureManualHandler(insSrv services.ProcedureManualService) *ProcedureManualHandler {
	return &ProcedureManualHandler{ProcedureManualSrv: insSrv}
}

func (h *ProcedureManualHandler) CreateProcedureManualHandler(c *fiber.Ctx) error {
	relPath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/procedure_manual",
		AllowedMIMEs: []string{"application/pdf"},
		MaxSize:      50 << 20, // 50MB
		BaseURL:      "",
		Required:     true,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := models.CreateProcedureManualRequest{
		ProcedureManualName: c.FormValue("procedure_manual_name"),
		Desc:                c.FormValue("desc"),
		Category:            c.FormValue("category"),
		FileName:            fmt.Sprintf("[\"%s\"]", relPath),
	}

	if err := h.ProcedureManualSrv.CreateProcedureManualService(req); err != nil {
		log.Println("Error creating procedure manual:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create procedure manual"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Procedure manual created successfully",
		"file":    relPath,
	})
}

// Get all procedure manuals
func (h *ProcedureManualHandler) GetAllProcedureManualHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	manuals, err := h.ProcedureManualSrv.GetAllProcedureManualService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch procedure manuals",
		})
	}

	return c.Status(http.StatusOK).JSON(manuals)
}

// Update procedure manual
func (h *ProcedureManualHandler) UpdateProcedureManualHandler(c *fiber.Ctx) error {
	procedureManualIDStr := c.Params("procedure_manual_id")
	procedureManualID, err := strconv.Atoi(procedureManualIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid procedure manual ID",
		})
	}

	log.Printf("[UpdateProcedureManualHandler] Received ID: %d\n", procedureManualID)

	var relPath string
	// Handle optional file upload
	if c.Request().Header.Peek("Content-Type") != nil {
		log.Println("[UpdateProcedureManualHandler] Processing file upload...")
		relPath, _, err = uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/procedure_manual",
			AllowedMIMEs: []string{"application/pdf"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err != nil {
			log.Printf("[UpdateProcedureManualHandler] File upload error: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		log.Printf("[UpdateProcedureManualHandler] File uploaded successfully: %s\n", relPath)
	}

	req := models.UpdateProcedureManualRequest{
		ProcedureManualName: c.FormValue("procedure_manual_name"),
		Desc:                c.FormValue("desc"),
		Category:            c.FormValue("category"),
	}

	// Only set FileName if a new file was uploaded
	if relPath != "" {
		req.FileName = fmt.Sprintf("[\"%s\"]", relPath)
	}

	log.Printf("[UpdateProcedureManualHandler] Calling UpdateProcedureManualService with ID: %d\n", procedureManualID)
	if err := h.ProcedureManualSrv.UpdateProcedureManualService(procedureManualID, req); err != nil {
		log.Printf("[UpdateProcedureManualHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update procedure manual"})
	}

	log.Println("[UpdateProcedureManualHandler] Update completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Procedure manual updated successfully",
	})
}

// Delete procedure manual
func (h *ProcedureManualHandler) DeleteProcedureManualHandler(c *fiber.Ctx) error {
	procedureManualIDStr := c.Params("procedure_manual_id")
	procedureManualID, err := strconv.Atoi(procedureManualIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid procedure manual ID",
		})
	}

	log.Printf("[DeleteProcedureManualHandler] Received ID: %d\n", procedureManualID)

	if err := h.ProcedureManualSrv.DeleteProcedureManualService(procedureManualID); err != nil {
		log.Printf("[DeleteProcedureManualHandler] Service error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete procedure manual"})
	}

	log.Println("[DeleteProcedureManualHandler] Delete completed successfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Procedure manual deleted successfully",
	})
}
