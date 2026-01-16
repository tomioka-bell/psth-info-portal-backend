package domains

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ProductID         string `gorm:"type:char(36);primaryKey;default:(newid())"` // รหัสผลิตภัณฑ์
	ProductName       string
	Category          string
	Description       string
	Recommend         bool
	ProductMainImages string
	ProductImages     string `gorm:"type:text"` // รูปภาพหลายรูป (JSON array เช่น ["uploads/company_news/20251103_103916_a2ba9c7d430a2768.JPG","uploads/company_news/20251103_103916_a2ba9c7d430a2769.JPG"])
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
