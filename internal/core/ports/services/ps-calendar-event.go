package ports

import "backend/internal/core/models"

type CalendarEventService interface {
	CreateCalendarEvent(req *models.CompanyCalendarEventRequest) (*models.CompanyCalendarEventResponse, error)
	GetCalendarEventByID(id int) (*models.CompanyCalendarEventResponse, error)
	GetAllCalendarEvents(page, pageSize int) ([]models.CompanyCalendarEventResponse, int64, error)
	UpdateCalendarEvent(id int, req *models.CompanyCalendarEventRequest) error
	UpdateCalendarEventPartial(id int, req *models.UpdateCalendarEventRequest) error
	DeleteCalendarEvent(id int) error
}
