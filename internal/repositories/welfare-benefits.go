package repositories

import (
	"errors"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type WelfareBenefitRepository struct {
	db *gorm.DB
}

func NewWelfareBenefitRepository(db *gorm.DB) *WelfareBenefitRepository {
	db.AutoMigrate(&domains.WelfareBenefit{})
	return &WelfareBenefitRepository{db: db}
}

// CreateWelfareBenefit creates a new welfare benefit
func (r *WelfareBenefitRepository) CreateWelfareBenefit(benefit *domains.WelfareBenefit) error {
	if err := r.db.Create(benefit).Error; err != nil {
		return err
	}
	return nil
}

// GetWelfareBenefitByID retrieves welfare benefit by ID
func (r *WelfareBenefitRepository) GetWelfareBenefitByID(id int) (*domains.WelfareBenefit, error) {
	var benefit domains.WelfareBenefit
	if err := r.db.Where("welfare_benefit_id = ?", id).First(&benefit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("welfare benefit not found")
		}
		return nil, err
	}
	return &benefit, nil
}

// GetAllWelfareBenefits retrieves all welfare benefits with pagination
func (r *WelfareBenefitRepository) GetAllWelfareBenefits(limit, offset int) ([]domains.WelfareBenefit, int64, error) {
	var benefits []domains.WelfareBenefit
	var total int64

	if err := r.db.Model(&domains.WelfareBenefit{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&benefits).Error; err != nil {
		return nil, 0, err
	}

	return benefits, total, nil
}

// GetWelfareBenefitsByCategory retrieves welfare benefits by category
func (r *WelfareBenefitRepository) GetWelfareBenefitsByCategory(category string, limit, offset int) ([]domains.WelfareBenefit, int64, error) {
	var benefits []domains.WelfareBenefit
	var total int64

	if err := r.db.Model(&domains.WelfareBenefit{}).Where("category = ?", category).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("category = ?", category).Limit(limit).Offset(offset).Order("created_at DESC").Find(&benefits).Error; err != nil {
		return nil, 0, err
	}

	return benefits, total, nil
}

// UpdateWelfareBenefit updates an existing welfare benefit
func (r *WelfareBenefitRepository) UpdateWelfareBenefit(id int, updates map[string]interface{}) error {
	if err := r.db.Model(&domains.WelfareBenefit{}).Where("welfare_benefit_id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

// DeleteWelfareBenefit soft deletes a welfare benefit
func (r *WelfareBenefitRepository) DeleteWelfareBenefit(id int) error {
	if err := r.db.Where("welfare_benefit_id = ?", id).Delete(&domains.WelfareBenefit{}).Error; err != nil {
		return err
	}
	return nil
}

// SearchWelfareBenefits searches welfare benefits by keyword
func (r *WelfareBenefitRepository) SearchWelfareBenefits(keyword string, limit, offset int) ([]domains.WelfareBenefit, int64, error) {
	var benefits []domains.WelfareBenefit
	var total int64

	query := r.db.Model(&domains.WelfareBenefit{}).
		Where("title ILIKE ? OR description ILIKE ? OR category ILIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&benefits).Error; err != nil {
		return nil, 0, err
	}

	return benefits, total, nil
}
