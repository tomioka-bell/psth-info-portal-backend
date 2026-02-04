package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type CalendarEventRepository struct {
	db *gorm.DB
}

func NewCalendarEventRepository(db *gorm.DB) *CalendarEventRepository {
	if err := db.AutoMigrate(&domains.CompanyCalendarEvent{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &CalendarEventRepository{db: db}
}

// CreateCalendarEvent creates a new procedure manual
func (r *CalendarEventRepository) CreateCalendarEvent(manual *domains.CompanyCalendarEvent) error {
	result := r.db.Create(manual)
	return result.Error
}

// GetCalendarEventByID retrieves company calendar event by ID
func (r *CalendarEventRepository) GetCalendarEventByID(id int) (*domains.CompanyCalendarEvent, error) {
	var event domains.CompanyCalendarEvent
	result := r.db.Where("calendar_event_id = ? AND deleted_at IS NULL", id).First(&event)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("calendar event not found")
		}
		return nil, result.Error
	}
	return &event, nil
}

// GetAllCalendarEvents retrieves all company calendar events with pagination
func (r *CalendarEventRepository) GetAllCalendarEvents(page, pageSize int) ([]domains.CompanyCalendarEvent, int64, error) {
	var events []domains.CompanyCalendarEvent
	var total int64

	if err := r.db.Model(&domains.CompanyCalendarEvent{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("deleted_at IS NULL").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// UpdateCalendarEvent updates an existing company calendar event
func (r *CalendarEventRepository) UpdateCalendarEvent(id int, updates map[string]interface{}) error {
	result := r.db.Model(&domains.CompanyCalendarEvent{}).
		Where("calendar_event_id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("calendar event not found or already deleted")
	}
	return result.Error
}

// DeleteCalendarEvent soft deletes a company calendar event
func (r *CalendarEventRepository) DeleteCalendarEvent(id int) error {
	result := r.db.Model(&domains.CompanyCalendarEvent{}).
		Where("calendar_event_id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("GETUTCDATE()"))

	if result.RowsAffected == 0 {
		return errors.New("calendar event not found or already deleted")
	}
	return result.Error
}
