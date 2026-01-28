package ports

import "backend/internal/core/models"

type OrganizationService interface {
	CreateOrganizationService(req models.CreateOrganizationRequest) error
	GetAllOrganizationsService(limit, offset int) (models.OrganizationListResponse, error)
	GetOrganizationsByCategoryService(category string, limit, offset int) (models.OrganizationListResponse, error)
	GetOrganizationByIDService(id int) (*models.OrganizationResponse, error)
	UpdateOrganizationService(id int, req models.UpdateOrganizationRequest) error
	DeleteOrganizationService(id int) error
	SearchOrganizationsService(keyword string, limit, offset int) (models.OrganizationListResponse, error)
}
