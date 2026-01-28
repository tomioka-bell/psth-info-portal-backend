package ports

import "backend/internal/core/models"

type OrganizationDocService interface {
	CreateOrganizationDocService(req models.CreateOrganizationDocRequest) error
	GetAllOrganizationDocService(page, pageSize int) (*models.OrganizationDocListResponse, error)
	GetOrganizationDocByIDService(id int) (*models.OrganizationDocResponse, error)
	GetOrganizationDocByDepartmentService(department string, page, pageSize int) (*models.OrganizationDocListResponse, error)
	UpdateOrganizationDocService(id int, req models.UpdateOrganizationDocRequest) error
	DeleteOrganizationDocService(id int) error
	SearchOrganizationDocService(keyword string, page, pageSize int) (*models.OrganizationDocListResponse, error)
}
