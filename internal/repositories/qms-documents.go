package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type QmsDocumentsRepository struct {
	db *gorm.DB
}

func NewQmsDocumentsRepository(db *gorm.DB) *QmsDocumentsRepository {
	if err := db.AutoMigrate(&domains.User{}, &domains.QmsDocuments{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &QmsDocumentsRepository{db: db}
}

func (r *QmsDocumentsRepository) CreateQmsDocument(doc *domains.QmsDocuments) error {
	result := r.db.Create(doc)
	return result.Error
}

func (r *QmsDocumentsRepository) GetQmsDocumentByID(id int) (*domains.QmsDocuments, error) {
	var doc domains.QmsDocuments
	// Use primary key lookup which respects GORM soft-delete behavior
	result := r.db.First(&doc, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("qms document not found")
		}
		return nil, result.Error
	}
	return &doc, nil
}

func (r *QmsDocumentsRepository) GetAllQmsDocuments(page, pageSize int) ([]domains.QmsDocuments, int64, error) {
	var docs []domains.QmsDocuments
	var total int64

	if err := r.db.Model(&domains.QmsDocuments{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("deleted_at IS NULL").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

func (r *QmsDocumentsRepository) UpdateQmsDocument(id int, updates map[string]interface{}) error {
	// First find the record by primary key (this will skip soft-deleted records)
	var doc domains.QmsDocuments
	res := r.db.First(&doc, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("qms document not found or already deleted")
		}
		return res.Error
	}

	// Apply updates to struct fields instead of using map->column Updates to avoid relying on exact DB column names
	if v, ok := updates["qms_documents_name"]; ok {
		if s, ok2 := v.(string); ok2 {
			doc.QmsDocumentsName = s
		}
	}
	// support both 'dqms_documents_desc' and generic 'desc' keys
	if v, ok := updates["dqms_documents_desc"]; ok {
		if s, ok2 := v.(string); ok2 {
			doc.DQmsDocumentsDesc = s
		}
	} else if v, ok := updates["desc"]; ok {
		if s, ok2 := v.(string); ok2 {
			doc.DQmsDocumentsDesc = s
		}
	}
	if v, ok := updates["category"]; ok {
		if s, ok2 := v.(string); ok2 {
			doc.Category = s
		}
	}
	if v, ok := updates["file_name"]; ok {
		if s, ok2 := v.(string); ok2 {
			doc.FileName = s
		}
	}

	result := r.db.Save(&doc)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("qms document not found or already deleted")
	}
	return nil
}

func (r *QmsDocumentsRepository) DeleteQmsDocument(id int) error {
	var doc domains.QmsDocuments
	res := r.db.First(&doc, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("qms document not found or already deleted")
		}
		return res.Error
	}

	result := r.db.Model(&doc).Update("deleted_at", gorm.Expr("GETUTCDATE()"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("qms document not found or already deleted")
	}
	return nil
}

func (r *QmsDocumentsRepository) SearchQmsDocuments(keyword string, page, pageSize int) ([]domains.QmsDocuments, int64, error) {
	var docs []domains.QmsDocuments
	var total int64

	// Search 'desc' column for description to align with existing schema naming
	query := r.db.Where("(qms_documents_name LIKE ? OR desc LIKE ?) AND deleted_at IS NULL", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Model(&domains.QmsDocuments{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}
