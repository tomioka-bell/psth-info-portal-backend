package ports

import "backend/internal/core/models"

type SafetyDocumentService interface {
	CreateSafetyDocumentService(req models.CreateSafetyDocumentRequest) error
	GetAllSafetyDocumentService(page, pageSize int) (*models.SafetyDocumentListResponse, error)
	GetSafetyDocumentByIDService(id int) (*models.SafetyDocumentResponse, error)
	GetSafetyDocumentByCategoryService(category string, page, pageSize int) (*models.SafetyDocumentListResponse, error)
	GetSafetyDocumentByDepartmentService(department string, page, pageSize int) (*models.SafetyDocumentListResponse, error)
	UpdateSafetyDocumentService(id int, req models.UpdateSafetyDocumentRequest) error
	DeleteSafetyDocumentService(id int) error
	SearchSafetyDocumentService(keyword string, page, pageSize int) (*models.SafetyDocumentListResponse, error)
}
