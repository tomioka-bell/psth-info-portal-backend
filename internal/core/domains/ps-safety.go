package domains

import (
	"time"

	"gorm.io/gorm"
)

type SafetyDocument struct {
	SafetyDocumentID   int            `gorm:"primaryKey;autoIncrement"`
	SafetyDocumentName string         `gorm:"type:nvarchar(255);not null"`
	SafetyDocumentDesc string         `gorm:"type:nvarchar(max);not null"` // Description in Thai
	Category           string         `gorm:"type:varchar(50);not null"`   // office, production, quality, support
	FileName           string         `gorm:"type:nvarchar(255)"`
	Department         string         `gorm:"type:nvarchar(255)"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for the SafetyDocument model
func (SafetyDocument) TableName() string {
	return "ps_safetys_documents"
}
