package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
)

type AppSystemHandler struct {
	AppSystemSrv services.AppSystemService
}

func NewAppSystemHandler(insSrv services.AppSystemService) *AppSystemHandler {
	return &AppSystemHandler{AppSystemSrv: insSrv}
}

// @Router /api/AppSystems [post]
func (h *AppSystemHandler) CreateAppSystemHandler(c *fiber.Ctx) error {
	// ลองรับ nested structure ก่อน
	var nestedReq []models.CreateAppSystemNestedRequest

	if err := c.BodyParser(&nestedReq); err != nil {
		// ถ้า parse ไม่ได้ ลองรับแบบ single object
		var singleReq models.CreateAppSystemNestedRequest
		if err := c.BodyParser(&singleReq); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		nestedReq = []models.CreateAppSystemNestedRequest{singleReq}
	}

	if len(nestedReq) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Request body is empty",
		})
	}

	// Recursive function สร้าง AppSystem + children
	var createRecursive func(req models.CreateAppSystemNestedRequest, parentID *int) error
	createRecursive = func(req models.CreateAppSystemNestedRequest, parentID *int) error {
		createReq := models.CreateAppSystemRequest{
			Name:      req.Name,
			Desc:      req.Desc,
			Category:  req.Category,
			Href:      req.Href,
			Icon:      req.Icon,
			SortOrder: req.SortOrder,
			ParentID:  parentID,
		}

		// สร้าง parent item
		resp, err := h.AppSystemSrv.CreateAppSystemWithResponse(createReq)
		if err != nil {
			return err
		}

		// สร้าง children recursively ด้วย parent ID ที่เพิ่งสร้าง
		for _, child := range req.Children {
			if err := createRecursive(child, &resp.ID); err != nil {
				return err
			}
		}

		return nil
	}

	// สร้างทั้งหมด
	for _, item := range nestedReq {
		if err := createRecursive(item, nil); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "AppSystem created successfully",
	})
}

// @Router /api/AppSystems/{id} [get]
func (h *AppSystemHandler) GetAppSystemHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid AppSystem ID",
		})
	}

	org, err := h.AppSystemSrv.GetAppSystemByIDService(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "AppSystem not found",
		})
	}

	return c.Status(http.StatusOK).JSON(org)
}

func (h *AppSystemHandler) CreateAppSystemsHandler(c *fiber.Ctx) error {
	var body []models.AppSystemCategoryRequest

	// parse body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if len(body) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Request body is empty",
		})
	}

	failedCount := 0
	successCount := 0
	var errors []string

	// flatten category -> systems[]
	for _, group := range body {
		for _, sys := range group.Systems {
			sys.Category = group.Category // inject category

			if err := h.AppSystemSrv.CreateAppSystemService(sys); err != nil {
				failedCount++
				errors = append(errors, err.Error())
				continue
			}
			successCount++
		}
	}

	response := fiber.Map{
		"message":       "AppSystems processed",
		"success_count": successCount,
		"failed_count":  failedCount,
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// @Router /api/AppSystems [get]
func (h *AppSystemHandler) GetAllAppSystemsHandler(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "50"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	orgs, err := h.AppSystemSrv.GetAllAppSystemsService(pageSize, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch AppSystems",
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}

// @Router /api/app-system/tree [get]
func (h *AppSystemHandler) GetAppSystemsTreeHandler(c *fiber.Ctx) error {
	tree, err := h.AppSystemSrv.GetAllAppSystemsTree()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch AppSystems tree",
		})
	}

	return c.Status(http.StatusOK).JSON(tree)
}

// @Router /api/AppSystems/category/{category} [get]
func (h *AppSystemHandler) GetAppSystemsByCategoryHandler(c *fiber.Ctx) error {
	category := c.Params("category")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	orgs, err := h.AppSystemSrv.GetAppSystemsByCategoryService(category, pageSize, offset)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}

// @Router /api/AppSystems/{id} [put]
func (h *AppSystemHandler) UpdateAppSystemHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid AppSystem ID",
		})
	}

	var req models.UpdateAppSystemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.AppSystemSrv.UpdateAppSystemService(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "AppSystem updated successfully",
	})
}

// @Router /api/AppSystems/{id} [delete]
func (h *AppSystemHandler) DeleteAppSystemHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid AppSystem ID",
		})
	}

	err = h.AppSystemSrv.DeleteAppSystemService(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "AppSystem not found",
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// @Router /api/AppSystems/search [get]
func (h *AppSystemHandler) SearchAppSystemsHandler(c *fiber.Ctx) error {
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

	orgs, err := h.AppSystemSrv.SearchAppSystemsService(keyword, pageSize, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search AppSystems",
		})
	}

	return c.Status(http.StatusOK).JSON(orgs)
}
