package repositories

import (
	"errors"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type ProcedureManualRepository struct {
	db *gorm.DB
}

func NewProcedureManualRepository(db *gorm.DB) *ProcedureManualRepository {
	// if err := db.AutoMigrate(&domains.ProcedureManual{}); err != nil {
	// 	fmt.Printf("failed to auto migrate: %v", err)
	// }
	return &ProcedureManualRepository{db: db}
}

// CreateProcedureManual creates a new procedure manual
func (r *ProcedureManualRepository) CreateProcedureManual(manual *domains.ProcedureManual) error {
	result := r.db.Create(manual)
	return result.Error
}

// GetProcedureManualByID retrieves procedure manual by ID
func (r *ProcedureManualRepository) GetProcedureManualByID(id int) (*domains.ProcedureManual, error) {
	var manual domains.ProcedureManual
	result := r.db.Where("procedure_manual_id = ? AND deleted_at IS NULL", id).First(&manual)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("procedure manual not found")
		}
		return nil, result.Error
	}
	return &manual, nil
}

// GetAllProcedureManuals retrieves all procedure manuals with pagination
func (r *ProcedureManualRepository) GetAllProcedureManuals(page, pageSize int) ([]domains.ProcedureManual, int64, error) {
	var manuals []domains.ProcedureManual
	var total int64

	if err := r.db.Model(&domains.ProcedureManual{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("deleted_at IS NULL").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}

// GetProcedureManualsByCategory retrieves procedure manuals filtered by category
func (r *ProcedureManualRepository) GetProcedureManualsByCategory(category string, page, pageSize int) ([]domains.ProcedureManual, int64, error) {
	var manuals []domains.ProcedureManual
	var total int64

	if err := r.db.Model(&domains.ProcedureManual{}).Where("category = ? AND deleted_at IS NULL", category).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("category = ? AND deleted_at IS NULL", category).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}

// UpdateProcedureManual updates an existing procedure manual
func (r *ProcedureManualRepository) UpdateProcedureManual(id int, updates map[string]interface{}) error {
	result := r.db.Model(&domains.ProcedureManual{}).
		Where("procedure_manual_id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("procedure manual not found or already deleted")
	}
	return result.Error
}

// DeleteProcedureManual soft deletes a procedure manual
func (r *ProcedureManualRepository) DeleteProcedureManual(id int) error {
	result := r.db.Model(&domains.ProcedureManual{}).
		Where("procedure_manual_id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("GETUTCDATE()"))

	if result.RowsAffected == 0 {
		return errors.New("procedure manual not found or already deleted")
	}
	return result.Error
}

// SearchProcedureManuals searches procedure manuals by keyword
func (r *ProcedureManualRepository) SearchProcedureManuals(keyword string, page, pageSize int) ([]domains.ProcedureManual, int64, error) {
	var manuals []domains.ProcedureManual
	var total int64

	query := r.db.Where("(procedure_manual_name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Model(&domains.ProcedureManual{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&manuals).Error; err != nil {
		return nil, 0, err
	}

	return manuals, total, nil
}
