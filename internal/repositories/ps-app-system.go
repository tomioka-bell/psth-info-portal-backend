package repositories

import (
	"errors"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type AppSystemRepository struct {
	db *gorm.DB
}

func NewAppSystemRepository(db *gorm.DB) *AppSystemRepository {
	// AutoMigrate ได้แค่ AppSystem เท่านั้น (AppSystemMenu เป็น DTO ไม่ต้องใน DB)
	// if err := db.AutoMigrate(&domains.AppSystem{}); err != nil {
	// 	fmt.Printf("failed to auto migrate AppSystem: %v\n", err)
	// }
	return &AppSystemRepository{db: db}
}

// CreateAppSystem creates a new AppSystem
func (r *AppSystemRepository) CreateAppSystem(org *domains.AppSystem) error {
	result := r.db.Debug().Create(org)
	return result.Error
}

// GetAppSystemByID retrieves AppSystem by ID
func (r *AppSystemRepository) GetAppSystemByID(id int) (*domains.AppSystem, error) {
	var org domains.AppSystem
	result := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&org)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("AppSystem not found")
		}
		return nil, result.Error
	}
	return &org, nil
}

// GetAllAppSystems retrieves all AppSystems with pagination
func (r *AppSystemRepository) GetAllAppSystems(page, pageSize int) ([]domains.AppSystem, int64, error) {
	var AppSystems []domains.AppSystem
	var total int64

	// Get total count
	countResult := r.db.Where("deleted_at IS NULL").Model(&domains.AppSystem{}).Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated data
	offset := (page - 1) * pageSize
	result := r.db.Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&AppSystems)

	return AppSystems, total, result.Error
}

// GetAppSystemsByCategory retrieves AppSystems filtered by category
func (r *AppSystemRepository) GetAppSystemsByCategory(category string, page, pageSize int) ([]domains.AppSystem, int64, error) {
	var AppSystems []domains.AppSystem
	var total int64

	// Get total count
	countResult := r.db.Where("category = ? AND deleted_at IS NULL", category).Model(&domains.AppSystem{}).Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated data
	offset := (page - 1) * pageSize
	result := r.db.Where("category = ? AND deleted_at IS NULL", category).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&AppSystems)

	return AppSystems, total, result.Error
}

// UpdateAppSystem updates an existing AppSystem
func (r *AppSystemRepository) UpdateAppSystem(id int, updates map[string]interface{}) error {
	result := r.db.Model(&domains.AppSystem{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates)

	if result.RowsAffected == 0 {
		return errors.New("AppSystem not found or already deleted")
	}
	return result.Error
}

// DeleteAppSystem soft deletes an AppSystem
func (r *AppSystemRepository) DeleteAppSystem(id int) error {
	result := r.db.Model(&domains.AppSystem{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("GETUTCDATE()"))

	if result.RowsAffected == 0 {
		return errors.New("AppSystem not found or already deleted")
	}
	return result.Error
}

// DeleteAppSystemPermanently permanently deletes an AppSystem
func (r *AppSystemRepository) DeleteAppSystemPermanently(id int) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&domains.AppSystem{})
	if result.RowsAffected == 0 {
		return errors.New("AppSystem not found")
	}
	return result.Error
}

// SearchAppSystems searches AppSystems by name or description
func (r *AppSystemRepository) SearchAppSystems(keyword string, page, pageSize int) ([]domains.AppSystem, int64, error) {
	var AppSystems []domains.AppSystem
	var total int64

	// Get total count
	countResult := r.db.Where("(name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%").
		Model(&domains.AppSystem{}).
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
		Find(&AppSystems)

	return AppSystems, total, result.Error
}

// GetAllAppSystemsForTree retrieves all AppSystems (ไม่ pagination) เพื่อสร้าง tree structure
func (r *AppSystemRepository) GetAllAppSystemsForTree() ([]domains.AppSystem, error) {
	var systems []domains.AppSystem
	result := r.db.Where("deleted_at IS NULL").
		Order("sort_order ASC, created_at ASC").
		Find(&systems)
	return systems, result.Error
}

// GetAppSystemTree retrieves AppSystems as hierarchical tree structure
func (r *AppSystemRepository) GetAppSystemTree() ([]domains.AppSystemMenu, error) {
	systems, err := r.GetAllAppSystemsForTree()
	if err != nil {
		return nil, err
	}
	return domains.BuildAppSystemTree(systems), nil
}
