package ports

import (
	"backend/internal/core/domains"
	"backend/internal/core/models"
)

type AppSystemService interface {
	CreateAppSystemService(req models.CreateAppSystemRequest) error
	CreateAppSystemWithResponse(req models.CreateAppSystemRequest) (*models.AppSystemResponse, error)
	GetAllAppSystemsService(limit, offset int) (models.AppSystemListResponse, error)
	GetAllAppSystemsTree() ([]domains.AppSystemMenu, error)
	GetAppSystemsByCategoryService(category string, limit, offset int) (models.AppSystemListResponse, error)
	GetAppSystemByIDService(id int) (*models.AppSystemResponse, error)
	UpdateAppSystemService(id int, req models.UpdateAppSystemRequest) error
	DeleteAppSystemService(id int) error
	SearchAppSystemsService(keyword string, limit, offset int) (models.AppSystemListResponse, error)
}
