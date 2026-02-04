package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/internal/core/models"
	services "backend/internal/core/ports/services"
)

type CalendarEventHandler struct {
	CalendarEventSrv services.CalendarEventService
}

func NewCalendarEventHandler(insSrv services.CalendarEventService) *CalendarEventHandler {
	return &CalendarEventHandler{CalendarEventSrv: insSrv}
}

func (h *CalendarEventHandler) CreateCalendarEventHandler(c *fiber.Ctx) error {
	var req models.CompanyCalendarEventRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
	}

	if req.EventType == "" || req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "EventType และ Title จำเป็นต้องกรอก",
		})
	}

	resp, err := h.CalendarEventSrv.CreateCalendarEvent(&req)
	if err != nil {
		log.Println("Error creating calendar event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถสร้างกิจกรรมได้",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "สร้างกิจกรรมสำเร็จ",
		"data":    resp,
	})
}

// GetCalendarEventByIDHandler gets a calendar event by ID
func (h *CalendarEventHandler) GetCalendarEventByIDHandler(c *fiber.Ctx) error {
	eventIDStr := c.Params("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event ID ไม่ถูกต้อง",
		})
	}

	event, err := h.CalendarEventSrv.GetCalendarEventByID(eventID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ไม่พบกิจกรรม",
		})
	}

	return c.Status(http.StatusOK).JSON(event)
}

// GetAllCalendarEventsHandler gets all calendar events
func (h *CalendarEventHandler) GetAllCalendarEventsHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 500)

	events, total, err := h.CalendarEventSrv.GetAllCalendarEvents(page, pageSize)
	if err != nil {
		log.Println("Error fetching calendar events:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถดึงข้อมูลกิจกรรมได้",
		})
	}

	totalPages := (int(total) + pageSize - 1) / pageSize

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":        events,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	})
}

// UpdateCalendarEventHandler updates a calendar event
func (h *CalendarEventHandler) UpdateCalendarEventHandler(c *fiber.Ctx) error {
	eventIDStr := c.Params("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event ID ไม่ถูกต้อง",
		})
	}

	var req models.UpdateCalendarEventRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ข้อมูลไม่ถูกต้อง",
		})
	}

	err = h.CalendarEventSrv.UpdateCalendarEventPartial(eventID, &req)
	if err != nil {
		log.Println("Error updating calendar event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถอัปเดตกิจกรรมได้",
		})
	}

	// Fetch updated event to return
	updatedEvent, err := h.CalendarEventSrv.GetCalendarEventByID(eventID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถดึงข้อมูลกิจกรรมได้",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "อัปเดตกิจกรรมสำเร็จ",
		"data":    updatedEvent,
	})
}

// DeleteCalendarEventHandler deletes a calendar event
func (h *CalendarEventHandler) DeleteCalendarEventHandler(c *fiber.Ctx) error {
	eventIDStr := c.Params("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event ID ไม่ถูกต้อง",
		})
	}

	err = h.CalendarEventSrv.DeleteCalendarEvent(eventID)
	if err != nil {
		log.Println("Error deleting calendar event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถลบกิจกรรมได้",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "ลบกิจกรรมสำเร็จ",
	})
}
