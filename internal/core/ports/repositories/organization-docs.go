package ports

import "backend/internal/core/domains"

type OrganizationDocRepository interface {
	CreateOrganizationDoc(doc *domains.OrganizationDoc) error
	GetOrganizationDocByID(id int) (*domains.OrganizationDoc, error)
	GetAllOrganizationDocs(page, pageSize int) ([]domains.OrganizationDoc, int64, error)
	GetOrganizationDocsByDepartment(department string, page, pageSize int) ([]domains.OrganizationDoc, int64, error)
	UpdateOrganizationDoc(id int, updates map[string]interface{}) error
	DeleteOrganizationDoc(id int) error
	SearchOrganizationDocs(keyword string, page, pageSize int) ([]domains.OrganizationDoc, int64, error)
}
