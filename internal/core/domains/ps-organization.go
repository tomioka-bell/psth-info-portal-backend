package domains

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	Name       string         `gorm:"type:nvarchar(255);not null"`
	Desc       string         `gorm:"type:nvarchar(max);not null"` // Description in Thai
	Category   string         `gorm:"type:varchar(50);not null"`   // office, production, quality, support
	Href       string         `gorm:"type:nvarchar(max);not null"`
	Icon       string         `gorm:"type:nvarchar(max)"` // React icon name or base64 encoded icon
	FileName   string         `gorm:"type:nvarchar(255)"`
	Department string         `gorm:"type:nvarchar(255)"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for the Organization model
func (Organization) TableName() string {
	return "ps_organizations"
}
