package domains

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationDoc struct {
	OrganizationDocID int            `gorm:"primaryKey;autoIncrement" json:"organization_doc_id"`
	Name              string         `json:"name"`
	Desc              string         `json:"desc"`
	Department        string         `json:"department"`
	FileName          string         `json:"file_name"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OrganizationDoc) TableName() string {
	return "organization_docs"
}
