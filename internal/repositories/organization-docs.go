package repositories

import (
	"errors"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type OrganizationDocRepository struct {
	db *gorm.DB
}

func NewOrganizationDocRepository(db *gorm.DB) *OrganizationDocRepository {
	// Auto-migrate
	db.AutoMigrate(&domains.OrganizationDoc{})
	return &OrganizationDocRepository{db: db}
}

// CreateOrganizationDoc creates a new organization doc
func (r *OrganizationDocRepository) CreateOrganizationDoc(doc *domains.OrganizationDoc) error {
	if err := r.db.Create(doc).Error; err != nil {
		return err
	}
	return nil
}

// GetOrganizationDocByID retrieves organization doc by ID
func (r *OrganizationDocRepository) GetOrganizationDocByID(id int) (*domains.OrganizationDoc, error) {
	var doc domains.OrganizationDoc
	if err := r.db.Where("organization_doc_id = ?", id).First(&doc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization doc not found")
		}
		return nil, err
	}
	return &doc, nil
}

// GetAllOrganizationDocs retrieves all organization docs with pagination
func (r *OrganizationDocRepository) GetAllOrganizationDocs(limit, offset int) ([]domains.OrganizationDoc, int64, error) {
	var docs []domains.OrganizationDoc
	var total int64

	if err := r.db.Model(&domains.OrganizationDoc{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

// GetOrganizationDocsByDepartment retrieves organization docs by department
func (r *OrganizationDocRepository) GetOrganizationDocsByDepartment(department string, limit, offset int) ([]domains.OrganizationDoc, int64, error) {
	var docs []domains.OrganizationDoc
	var total int64

	if err := r.db.Model(&domains.OrganizationDoc{}).Where("department = ?", department).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("department = ?", department).Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

// UpdateOrganizationDoc updates an existing organization doc
func (r *OrganizationDocRepository) UpdateOrganizationDoc(id int, updates map[string]interface{}) error {
	if err := r.db.Model(&domains.OrganizationDoc{}).Where("organization_doc_id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrganizationDoc soft deletes an organization doc
func (r *OrganizationDocRepository) DeleteOrganizationDoc(id int) error {
	if err := r.db.Where("organization_doc_id = ?", id).Delete(&domains.OrganizationDoc{}).Error; err != nil {
		return err
	}
	return nil
}

// SearchOrganizationDocs searches organization docs by keyword
func (r *OrganizationDocRepository) SearchOrganizationDocs(keyword string, limit, offset int) ([]domains.OrganizationDoc, int64, error) {
	var docs []domains.OrganizationDoc
	var total int64

	query := r.db.Model(&domains.OrganizationDoc{}).
		Where("name ILIKE ? OR desc ILIKE ? OR department ILIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}
