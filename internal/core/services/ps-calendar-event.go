package services

import (
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type CalendarEventService struct {
	repo *repositories.CalendarEventRepository
}

func NewCalendarEventService(repo *repositories.CalendarEventRepository) *CalendarEventService {
	return &CalendarEventService{repo: repo}
}

// CreateCalendarEvent creates a new CalendarEvent
func (s *CalendarEventService) CreateCalendarEvent(req *models.CompanyCalendarEventRequest) (*models.CompanyCalendarEventResponse, error) {
	event := &domains.CompanyCalendarEvent{
		EventType: req.EventType,
		Title:     req.Title,
		IsAllDay:  req.IsAllDay,
		StartAt:   req.StartAt,
		EndAt:     req.EndAt,
		IsActive:  req.IsActive,
	}

	err := s.repo.CreateCalendarEvent(event)
	if err != nil {
		return nil, err
	}

	return &models.CompanyCalendarEventResponse{
		CalendarEventID: event.CalendarEventID,
		EventType:       event.EventType,
		Title:           event.Title,
		IsAllDay:        event.IsAllDay,
		StartAt:         event.StartAt,
		EndAt:           event.EndAt,
		IsActive:        event.IsActive,
	}, nil
}

// GetCalendarEventByID retrieves CalendarEvent by ID
func (s *CalendarEventService) GetCalendarEventByID(id int) (*models.CompanyCalendarEventResponse, error) {
	event, err := s.repo.GetCalendarEventByID(id)
	if err != nil {
		return nil, err
	}

	return &models.CompanyCalendarEventResponse{
		CalendarEventID: event.CalendarEventID,
		EventType:       event.EventType,
		Title:           event.Title,
		IsAllDay:        event.IsAllDay,
		StartAt:         event.StartAt,
		EndAt:           event.EndAt,
		IsActive:        event.IsActive,
	}, nil
}

// GetAllCalendarEvents retrieves all CalendarEvents with pagination
func (s *CalendarEventService) GetAllCalendarEvents(page, pageSize int) ([]models.CompanyCalendarEventResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	events, total, err := s.repo.GetAllCalendarEvents(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]models.CompanyCalendarEventResponse, len(events))
	for i, event := range events {
		resp[i] = models.CompanyCalendarEventResponse{
			CalendarEventID: event.CalendarEventID,
			EventType:       event.EventType,
			Title:           event.Title,
			IsAllDay:        event.IsAllDay,
			StartAt:         event.StartAt,
			EndAt:           event.EndAt,
			IsActive:        event.IsActive,
		}
	}

	return resp, total, nil
}

// UpdateCalendarEvent updates an existing CalendarEvent
func (s *CalendarEventService) UpdateCalendarEvent(id int, req *models.CompanyCalendarEventRequest) error {
	// Verify CalendarEvent exists
	_, err := s.repo.GetCalendarEventByID(id)
	if err != nil {
		return err
	}

	updates := make(map[string]interface{})

	if req.EventType != "" {
		updates["event_type"] = req.EventType
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	updates["is_all_day"] = req.IsAllDay
	updates["start_at"] = req.StartAt
	updates["end_at"] = req.EndAt
	updates["is_active"] = req.IsActive
	updates["updated_at"] = time.Now()

	return s.repo.UpdateCalendarEvent(id, updates)
}

// UpdateCalendarEventPartial updates only the specified fields
func (s *CalendarEventService) UpdateCalendarEventPartial(id int, req *models.UpdateCalendarEventRequest) error {
	// Verify CalendarEvent exists
	_, err := s.repo.GetCalendarEventByID(id)
	if err != nil {
		return err
	}

	updates := make(map[string]interface{})

	if req.EventType != nil && *req.EventType != "" {
		updates["event_type"] = *req.EventType
	}
	if req.Title != nil && *req.Title != "" {
		updates["title"] = *req.Title
	}
	if req.IsAllDay != nil {
		updates["is_all_day"] = *req.IsAllDay
	}
	if req.StartAt != nil {
		updates["start_at"] = *req.StartAt
	}
	if req.EndAt != nil {
		updates["end_at"] = *req.EndAt
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	updates["updated_at"] = time.Now()

	return s.repo.UpdateCalendarEvent(id, updates)
}

func (s *CalendarEventService) DeleteCalendarEvent(id int) error {
	// Verify CalendarEvent exists
	_, err := s.repo.GetCalendarEventByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteCalendarEvent(id)
}
