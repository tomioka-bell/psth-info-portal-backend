package domains

import (
	"time"

	"gorm.io/gorm"
)

type CustomerManual struct {
	CustomerManualID   int            `gorm:"primaryKey;autoIncrement"`
	CustomerManualName string         `gorm:"type:nvarchar(255);not null"`
	Desc               string         `gorm:"type:nvarchar(max);not null"`
	Category           string         `gorm:"type:varchar(50);not null"`
	FileName           string         `gorm:"type:nvarchar(max);not null"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (CustomerManual) TableName() string {
	return "customer_manuals"
}
