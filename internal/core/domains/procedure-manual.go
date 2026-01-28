package domains

import (
	"time"

	"gorm.io/gorm"
)

type ProcedureManual struct {
	ProcedureManualID   int            `gorm:"primaryKey;autoIncrement"`
	ProcedureManualName string         `gorm:"type:nvarchar(255);not null"`
	Desc                string         `gorm:"type:nvarchar(max);not null"` // Description in Thai
	Category            string         `gorm:"type:varchar(50);not null"`   // office, production, quality, support, document
	FileName            string         `gorm:"type:nvarchar(max);not null"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

func (ProcedureManual) TableName() string {
	return "ps_procedure_manuals"
}
