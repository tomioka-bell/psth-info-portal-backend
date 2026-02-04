package models

import (
	"time"
)

type CompanyCalendarEventRequest struct {
	CalendarEventID uint      `json:"calendar_event_id"`
	EventType       string    `json:"event_type"`
	Title           string    `json:"title"`
	IsAllDay        bool      `json:"is_all_day"`
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
	IsActive        bool      `json:"is_active"`
}

type CompanyCalendarEventResponse struct {
	CalendarEventID uint      `json:"calendar_event_id"`
	EventType       string    `json:"event_type"`
	Title           string    `json:"title"`
	IsAllDay        bool      `json:"is_all_day"`
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
	IsActive        bool      `json:"is_active"`
}

type UpdateCalendarEventRequest struct {
	EventType *string    `json:"event_type,omitempty"`
	Title     *string    `json:"title,omitempty"`
	IsAllDay  *bool      `json:"is_all_day,omitempty"`
	StartAt   *time.Time `json:"start_at,omitempty"`
	EndAt     *time.Time `json:"end_at,omitempty"`
	IsActive  *bool      `json:"is_active,omitempty"`
}
