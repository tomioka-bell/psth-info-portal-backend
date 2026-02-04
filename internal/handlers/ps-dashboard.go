package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/services"
)

type DashboardHandler struct {
	DashboardSrv *services.DashboardService
}

func NewDashboardHandler(dashboardSrv *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		DashboardSrv: dashboardSrv,
	}
}

// GetDashboardStats returns all statistics and monthly data for dashboard
func (h *DashboardHandler) GetDashboardStats(c *fiber.Ctx) error {
	stats, err := h.DashboardSrv.GetDashboardStats()
	if err != nil {
		log.Printf("[GetDashboardStats] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch dashboard stats",
		})
	}

	return c.Status(http.StatusOK).JSON(stats)
}

// GetTableStats returns statistics for a specific table
func (h *DashboardHandler) GetTableStats(c *fiber.Ctx) error {
	tableName := c.Params("table")
	if tableName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Table name is required",
		})
	}

	monthlyStats, err := h.DashboardSrv.GetTableMonthlyStats(tableName)
	if err != nil {
		log.Printf("[GetTableStats] Error for table %s: %v\n", tableName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch table stats",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"table":         tableName,
		"monthly_stats": monthlyStats,
	})
}

// GetAllCounts returns counts from all tables
func (h *DashboardHandler) GetAllCounts(c *fiber.Ctx) error {
	counts, err := h.DashboardSrv.GetDashboardStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch counts",
		})
	}

	return c.Status(http.StatusOK).JSON(counts)
}

// GetAppSystemsCategory returns app systems grouped by category
func (h *DashboardHandler) GetAppSystemsCategory(c *fiber.Ctx) error {
	categories, err := h.DashboardSrv.GetAppSystemsCategory()
	if err != nil {
		log.Printf("[GetAppSystemsCategory] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch app systems by category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}

// GetOrganizationsCategory returns organizations grouped by category
func (h *DashboardHandler) GetOrganizationsCategory(c *fiber.Ctx) error {
	categories, err := h.DashboardSrv.GetOrganizationsCategory()
	if err != nil {
		log.Printf("[GetOrganizationsCategory] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch organizations by category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}

// GetSafetyDocsCategory returns safety documents grouped by category
func (h *DashboardHandler) GetSafetyDocsCategory(c *fiber.Ctx) error {
	categories, err := h.DashboardSrv.GetSafetyDocsCategory()
	if err != nil {
		log.Printf("[GetSafetyDocsCategory] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch safety documents by category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}

// GetSafetyDocsDepartment returns safety documents grouped by department
func (h *DashboardHandler) GetSafetyDocsDepartment(c *fiber.Ctx) error {
	departments, err := h.DashboardSrv.GetSafetyDocsDepartment()
	if err != nil {
		log.Printf("[GetSafetyDocsDepartment] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch safety documents by department",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"departments": departments,
	})
}

// GetWelfareBenefitsCategory returns welfare benefits grouped by category
func (h *DashboardHandler) GetWelfareBenefitsCategory(c *fiber.Ctx) error {
	categories, err := h.DashboardSrv.GetWelfareBenefitsCategory()
	if err != nil {
		log.Printf("[GetWelfareBenefitsCategory] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch welfare benefits by category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}

// GetNewsCategory returns company news grouped by category
func (h *DashboardHandler) GetNewsCategory(c *fiber.Ctx) error {
	categories, err := h.DashboardSrv.GetNewsCategory()
	if err != nil {
		log.Printf("[GetNewsCategory] Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch company news by category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}
