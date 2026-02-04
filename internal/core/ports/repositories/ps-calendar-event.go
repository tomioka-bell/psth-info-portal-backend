package ports

import "backend/internal/core/domains"

type CalendarEventRepository interface {
	GetCalendarEventByID(id int) (*domains.CompanyCalendarEvent, error)
	GetAllCalendarEvents(page, pageSize int) ([]domains.CompanyCalendarEvent, int64, error)
	CreateCalendarEvent(manual *domains.CompanyCalendarEvent) error
	UpdateCalendarEvent(id int, updates map[string]interface{}) error
	DeleteCalendarEvent(id int) error
}
