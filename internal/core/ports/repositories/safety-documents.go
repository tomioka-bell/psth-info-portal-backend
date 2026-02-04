package ports

import "backend/internal/core/domains"

type SafetyDocumentRepository interface {
	CreateSafetyDocument(doc *domains.SafetyDocument) error
	GetSafetyDocumentByID(id int) (*domains.SafetyDocument, error)
	GetAllSafetyDocuments(page, pageSize int) ([]domains.SafetyDocument, int64, error)
	GetSafetyDocumentsByCategory(category string, page, pageSize int) ([]domains.SafetyDocument, int64, error)
	GetSafetyDocumentsByDepartment(department string, page, pageSize int) ([]domains.SafetyDocument, int64, error)
	UpdateSafetyDocument(id int, updates map[string]interface{}) error
	DeleteSafetyDocument(id int) error
	SearchSafetyDocuments(keyword string, page, pageSize int) ([]domains.SafetyDocument, int64, error)
}
