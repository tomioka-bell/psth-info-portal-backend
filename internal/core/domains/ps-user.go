package domains

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    string         `gorm:"type:uniqueidentifier;primaryKey;default=NEWID()" json:"user_id"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Status    string         `json:"status"`
	Role      string         `json:"role"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Role struct {
	RoleID          string         `gorm:"type:uuid;primaryKey" json:"role_id"`
	RoleName        string         `json:"role_name"`
	RoleDescription string         `json:"role_description"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
