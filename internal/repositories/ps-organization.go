package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	if err := db.AutoMigrate(&domains.Organization{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &OrganizationRepository{db: db}
}

// CreateOrganization creates a new organization
func (r *OrganizationRepository) CreateOrganization(org *domains.Organization) error {
	result := r.db.Create(org)
	return result.Error
}

// GetOrganizationByID retrieves organization by ID
func (r *OrganizationRepository) GetOrganizationByID(id int) (*domains.Organization, error) {
	var org domains.Organization
	result := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&org)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, result.Error
	}
	return &org, nil
}

// GetAllOrganizations retrieves all organizations with pagination
func (r *OrganizationRepository) GetAllOrganizations(page, pageSize int) ([]domains.Organization, int64, error) {
	var organizations []domains.Organization
	var total int64

	// Get total count
	countResult := r.db.Where("deleted_at IS NULL").Model(&domains.Organization{}).Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated data
	offset := (page - 1) * pageSize
	result := r.db.Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&organizations)

	return organizations, total, result.Error
}

// GetOrganizationsByCategory retrieves organizations filtered by category
func (r *OrganizationRepository) GetOrganizationsByCategory(category string, page, pageSize int) ([]domains.Organization, int64, error) {
	var organizations []domains.Organization
	var total int64

	// Get total count
	countResult := r.db.Where("category = ? AND deleted_at IS NULL", category).Model(&domains.Organization{}).Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated data
	offset := (page - 1) * pageSize
	result := r.db.Where("category = ? AND deleted_at IS NULL", category).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&organizations)

	return organizations, total, result.Error
}

// UpdateOrganization updates an existing organization
func (r *OrganizationRepository) UpdateOrganization(id int, updates map[string]interface{}) error {
	result := r.db.Model(&domains.Organization{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("organization not found or already deleted")
	}
	return result.Error
}

// DeleteOrganization soft deletes an organization
func (r *OrganizationRepository) DeleteOrganization(id int) error {
	result := r.db.Model(&domains.Organization{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("GETUTCDATE()"))

	if result.RowsAffected == 0 {
		return errors.New("organization not found or already deleted")
	}
	return result.Error
}

// DeleteOrganizationPermanently permanently deletes an organization
func (r *OrganizationRepository) DeleteOrganizationPermanently(id int) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&domains.Organization{})
	if result.RowsAffected == 0 {
		return errors.New("organization not found")
	}
	return result.Error
}

// SearchOrganizations searches organizations by name or description
func (r *OrganizationRepository) SearchOrganizations(keyword string, page, pageSize int) ([]domains.Organization, int64, error) {
	var organizations []domains.Organization
	var total int64

	// Get total count
	countResult := r.db.Where("(name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%").
		Model(&domains.Organization{}).
		Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated data
	offset := (page - 1) * pageSize
	result := r.db.Where("(name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&organizations)

	return organizations, total, result.Error
}
