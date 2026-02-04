package repositories

import (
	"errors"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type SafetyDocumentRepository struct {
	db *gorm.DB
}

func NewSafetyDocumentRepository(db *gorm.DB) *SafetyDocumentRepository {
	// db.AutoMigrate(&domains.SafetyDocument{})
	return &SafetyDocumentRepository{db: db}
}

// CreateSafetyDocument creates a new safety document
func (r *SafetyDocumentRepository) CreateSafetyDocument(doc *domains.SafetyDocument) error {
	if err := r.db.Create(doc).Error; err != nil {
		return err
	}
	return nil
}

// GetSafetyDocumentByID retrieves safety document by ID
func (r *SafetyDocumentRepository) GetSafetyDocumentByID(id int) (*domains.SafetyDocument, error) {
	var doc domains.SafetyDocument
	if err := r.db.Where("safety_document_id = ?", id).First(&doc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("safety document not found")
		}
		return nil, err
	}
	return &doc, nil
}

// GetAllSafetyDocuments retrieves all safety documents with pagination
func (r *SafetyDocumentRepository) GetAllSafetyDocuments(limit, offset int) ([]domains.SafetyDocument, int64, error) {
	var docs []domains.SafetyDocument
	var total int64

	if err := r.db.Model(&domains.SafetyDocument{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

// GetSafetyDocumentsByCategory retrieves safety documents by category
func (r *SafetyDocumentRepository) GetSafetyDocumentsByCategory(category string, limit, offset int) ([]domains.SafetyDocument, int64, error) {
	var docs []domains.SafetyDocument
	var total int64

	if err := r.db.Model(&domains.SafetyDocument{}).Where("category = ?", category).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("category = ?", category).Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

// GetSafetyDocumentsByDepartment retrieves safety documents by department
func (r *SafetyDocumentRepository) GetSafetyDocumentsByDepartment(department string, limit, offset int) ([]domains.SafetyDocument, int64, error) {
	var docs []domains.SafetyDocument
	var total int64

	if err := r.db.Model(&domains.SafetyDocument{}).Where("department = ?", department).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("department = ?", department).Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

// UpdateSafetyDocument updates an existing safety document
func (r *SafetyDocumentRepository) UpdateSafetyDocument(id int, updates map[string]interface{}) error {
	if err := r.db.Model(&domains.SafetyDocument{}).Where("safety_document_id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSafetyDocument soft deletes a safety document
func (r *SafetyDocumentRepository) DeleteSafetyDocument(id int) error {
	if err := r.db.Where("safety_document_id = ?", id).Delete(&domains.SafetyDocument{}).Error; err != nil {
		return err
	}
	return nil
}

// SearchSafetyDocuments searches safety documents by keyword
func (r *SafetyDocumentRepository) SearchSafetyDocuments(keyword string, limit, offset int) ([]domains.SafetyDocument, int64, error) {
	var docs []domains.SafetyDocument
	var total int64

	query := r.db.Model(&domains.SafetyDocument{}).
		Where("safety_document_name ILIKE ? OR safety_document_desc ILIKE ? OR department ILIKE ? OR category ILIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}
