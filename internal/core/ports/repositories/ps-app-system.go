package ports

import "backend/internal/core/domains"

type AppSystemRepository interface {
	CreateAppSystem(org *domains.AppSystem) error
	GetAppSystemByID(id int) (*domains.AppSystem, error)
	GetAllAppSystems(page, pageSize int) ([]domains.AppSystem, int64, error)
	GetAppSystemsByCategory(category string, page, pageSize int) ([]domains.AppSystem, int64, error)
	UpdateAppSystem(id int, updates map[string]interface{}) error
	DeleteAppSystem(id int) error
	SearchAppSystems(keyword string, page, pageSize int) ([]domains.AppSystem, int64, error)
}
