package ports

import "backend/internal/core/domains"

type OrganizationRepository interface {
	CreateOrganization(org *domains.Organization) error
	GetOrganizationByID(id int) (*domains.Organization, error)
	GetAllOrganizations(page, pageSize int) ([]domains.Organization, int64, error)
	GetOrganizationsByCategory(category string, page, pageSize int) ([]domains.Organization, int64, error)
	UpdateOrganization(id int, updates map[string]interface{}) error
	DeleteOrganization(id int) error
	SearchOrganizations(keyword string, page, pageSize int) ([]domains.Organization, int64, error)
}
