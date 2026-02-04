package domains

import (
	"time"

	"gorm.io/gorm"
)

type CompanyCalendarEvent struct {
	CalendarEventID uint `gorm:"primaryKey"`

	// HOLIDAY | SPECIAL | PAYDAY
	EventType string `gorm:"size:20;index;not null;check:event_type IN ('HOLIDAY','SPECIAL','PAYDAY')"`

	Title string `gorm:"size:255;not null"`

	IsAllDay bool `gorm:"not null;default:true"`

	StartAt time.Time `gorm:"index;not null"`
	EndAt   time.Time `gorm:"index;not null"`

	IsActive bool `gorm:"not null;default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
