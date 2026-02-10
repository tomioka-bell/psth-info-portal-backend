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

type QmsDocumentsHandler struct {
	QmsDocumentsSrv services.QmsDocumentsService
}

func NewQmsDocumentsHandler(insSrv services.QmsDocumentsService) *QmsDocumentsHandler {
	return &QmsDocumentsHandler{QmsDocumentsSrv: insSrv}
}

func (h *QmsDocumentsHandler) CreateQmsDocumentHandler(c *fiber.Ctx) error {
	relPath, _, err := uploader.UploadFromForm(c, "file", uploader.Options{
		Dir:          "./uploads/qms_documents",
		AllowedMIMEs: []string{"application/pdf"},
		MaxSize:      50 << 20,
		BaseURL:      "",
		Required:     true,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := models.CreateQmsDocumentRequest{
		QmsDocumentsName:  c.FormValue("qms_documents_name"),
		DQmsDocumentsDesc: c.FormValue("dqms_documents_desc"),
		Category:          c.FormValue("category"),
		FileName:          fmt.Sprintf("[\"%s\"]", relPath),
	}

	if err := h.QmsDocumentsSrv.CreateQmsDocumentService(req); err != nil {
		log.Println("Error creating qms document:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create qms document"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "QMS document created successfully", "file": relPath})
}

func (h *QmsDocumentsHandler) GetAllQmsDocumentsHandler(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	docs, err := h.QmsDocumentsSrv.GetAllQmsDocumentService(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch qms documents"})
	}

	return c.Status(http.StatusOK).JSON(docs)
}

func (h *QmsDocumentsHandler) UpdateQmsDocumentHandler(c *fiber.Ctx) error {
	idStr := c.Params("qms_documents_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid qms document ID"})
	}

	var relPath string
	if c.Request().Header.Peek("Content-Type") != nil {
		relPath, _, err = uploader.UploadFromForm(c, "file", uploader.Options{
			Dir:          "./uploads/qms_documents",
			AllowedMIMEs: []string{"application/pdf"},
			MaxSize:      50 << 20,
			BaseURL:      "",
			Required:     false,
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	req := models.UpdateQmsDocumentRequest{
		QmsDocumentsName:  c.FormValue("qms_documents_name"),
		DQmsDocumentsDesc: c.FormValue("dqms_documents_desc"),
		Category:          c.FormValue("category"),
	}
	if relPath != "" {
		req.FileName = fmt.Sprintf("[\"%s\"]", relPath)
	}

	if err := h.QmsDocumentsSrv.UpdateQmsDocumentService(id, req); err != nil {
		log.Printf("Error updating qms document: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update qms document"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "QMS document updated successfully"})
}

func (h *QmsDocumentsHandler) DeleteQmsDocumentHandler(c *fiber.Ctx) error {
	idStr := c.Params("qms_documents_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid qms document ID"})
	}

	if err := h.QmsDocumentsSrv.DeleteQmsDocumentService(id); err != nil {
		log.Printf("Error deleting qms document: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete qms document"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "QMS document deleted successfully"})
}
