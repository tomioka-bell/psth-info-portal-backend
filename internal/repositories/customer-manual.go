package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type CustomerManualRepository struct {
	db *gorm.DB
}

func NewCustomerManualRepository(db *gorm.DB) *CustomerManualRepository {
	if err := db.AutoMigrate(&domains.CustomerManual{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &CustomerManualRepository{db: db}
}

// CreateCustomerManual creates a new customer manual
func (r *CustomerManualRepository) CreateCustomerManual(manual *domains.CustomerManual) error {
	result := r.db.Create(manual)
	return result.Error
}

// GetCustomerManualByID retrieves customer manual by ID
func (r *CustomerManualRepository) GetCustomerManualByID(id int) (*domains.CustomerManual, error) {
	var manual domains.CustomerManual
	result := r.db.Where("customer_manual_id = ? AND deleted_at IS NULL", id).First(&manual)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer manual not found")
		}
		return nil, result.Error
	}
	return &manual, nil
}

// GetAllCustomerManuals retrieves all customer manuals with pagination
func (r *CustomerManualRepository) GetAllCustomerManuals(page, pageSize int) ([]domains.CustomerManual, int64, error) {
	var manuals []domains.CustomerManual
	var total int64

	if err := r.db.Debug().Model(&domains.CustomerManual{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("deleted_at IS NULL").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}

// GetCustomerManualsByCategory retrieves customer manuals filtered by category
func (r *CustomerManualRepository) GetCustomerManualsByCategory(category string, page, pageSize int) ([]domains.CustomerManual, int64, error) {
	var manuals []domains.CustomerManual
	var total int64

	if err := r.db.Model(&domains.CustomerManual{}).Where("category = ? AND deleted_at IS NULL", category).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("category = ? AND deleted_at IS NULL", category).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}

// UpdateCustomerManual updates an existing customer manual
func (r *CustomerManualRepository) UpdateCustomerManual(id int, updates map[string]interface{}) error {
	result := r.db.Model(&domains.CustomerManual{}).
		Where("customer_manual_id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("customer manual not found or already deleted")
	}
	return result.Error
}

// DeleteCustomerManual soft deletes a customer manual
func (r *CustomerManualRepository) DeleteCustomerManual(id int) error {
	result := r.db.Model(&domains.CustomerManual{}).
		Where("customer_manual_id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("GETUTCDATE()"))

	if result.RowsAffected == 0 {
		return errors.New("customer manual not found or already deleted")
	}
	return result.Error
}

// SearchCustomerManuals searches customer manuals by keyword
func (r *CustomerManualRepository) SearchCustomerManuals(keyword string, page, pageSize int) ([]domains.CustomerManual, int64, error) {
	var manuals []domains.CustomerManual
	var total int64

	query := r.db.Where("(customer_manual_name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Model(&domains.CustomerManual{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}
