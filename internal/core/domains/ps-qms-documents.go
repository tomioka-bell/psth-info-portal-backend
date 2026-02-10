package domains

import (
	"time"

	"gorm.io/gorm"
)

type QmsDocuments struct {
	QmsDocumentsID    int            `gorm:"primaryKey;autoIncrement"`
	QmsDocumentsName  string         `gorm:"type:nvarchar(255);not null"`
	DQmsDocumentsDesc string         `gorm:"type:nvarchar(max);not null"` // Description in Thai
	Category          string         `gorm:"type:varchar(50);not null"`
	FileName          string         `gorm:"type:nvarchar(max);not null"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (QmsDocuments) TableName() string {
	return "ps_qms_documents"
}
