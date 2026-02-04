package domains

import (
	"time"

	"gorm.io/gorm"
)

type WelfareBenefit struct {
	WelfareBenefitID int            `gorm:"column:welfare_benefit_id;primaryKey;autoIncrement"`
	Title            string         `gorm:"column:title;type:varchar(255);not null"`
	Description      string         `gorm:"column:description;type:text;not null"`
	ImageURL         string         `gorm:"column:image_url;type:varchar(255)"`
	FileName         string         `gorm:"column:file_name;type:varchar(255)"`
	Category         string         `gorm:"type:varchar(50);not null"` // Permanent employee, Contract employee
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (WelfareBenefit) TableName() string {
	return "welfare_benefits"
}
