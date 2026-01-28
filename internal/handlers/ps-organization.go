package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
)

type OrganizationHandler struct {
	OrganizationSrv services.OrganizationService
}

func NewOrganizationHandler(insSrv services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{OrganizationSrv: insSrv}
}

// @Router /api/organizations [post]
func (h *OrganizationHandler) CreateOrganizationHandler(c *fiber.Ctx) error {
	var req models.CreateOrganizationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.OrganizationSrv.CreateOrganizationService(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Organization created successfully",
	})
}

// @Router /api/organizations/{id} [get]
func (h *OrganizationHandler) GetOrganizationHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid organization ID",
		})
	}

	org, err := h.OrganizationSrv.GetOrganizationByIDService(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	return c.Status(http.StatusOK).JSON(org)
}

func (h *OrganizationHandler) CreateOrganizationsHandler(c *fiber.Ctx) error {
	var reqs []models.CreateOrganizationRequest

	// parse body
	if err := c.BodyParser(&reqs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if len(reqs) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Request body is empty",
		})
	}

	// call service for each organization
	failedCount := 0
	for _, req := range reqs {
		if err := h.OrganizationSrv.CreateOrganizationService(req); err != nil {
			failedCount++
		}
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":      "Organizations created successfully",
		"count":        len(reqs),
		"failed_count": failedCount,
	})
}

// @Router /api/organizations [get]
func (h *OrganizationHandler) GetAllOrganizationsHandler(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "50"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	orgs, err := h.OrganizationSrv.GetAllOrganizationsService(pageSize, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch organizations",
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}

// @Router /api/organizations/category/{category} [get]
func (h *OrganizationHandler) GetOrganizationsByCategoryHandler(c *fiber.Ctx) error {
	category := c.Params("category")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	orgs, err := h.OrganizationSrv.GetOrganizationsByCategoryService(category, pageSize, offset)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}

// @Router /api/organizations/{id} [put]
func (h *OrganizationHandler) UpdateOrganizationHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid organization ID",
		})
	}

	var req models.UpdateOrganizationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.OrganizationSrv.UpdateOrganizationService(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Organization updated successfully",
	})
}

// @Router /api/organizations/{id} [delete]
func (h *OrganizationHandler) DeleteOrganizationHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid organization ID",
		})
	}

	err = h.OrganizationSrv.DeleteOrganizationService(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// @Router /api/organizations/search [get]
func (h *OrganizationHandler) SearchOrganizationsHandler(c *fiber.Ctx) error {
	keyword := c.Query("q", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	if keyword == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Search keyword is required",
		})
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	orgs, err := h.OrganizationSrv.SearchOrganizationsService(keyword, pageSize, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search organizations",
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}
