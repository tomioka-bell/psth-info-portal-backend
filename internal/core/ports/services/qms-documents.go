package ports

import "backend/internal/core/models"

type QmsDocumentsService interface {
	CreateQmsDocumentService(req models.CreateQmsDocumentRequest) error
	GetAllQmsDocumentService(limit, offset int) (models.QmsDocumentListResponse, error)
	GetQmsDocumentByIDService(id int) (*models.QmsDocumentResponse, error)
	UpdateQmsDocumentService(id int, req models.UpdateQmsDocumentRequest) error
	DeleteQmsDocumentService(id int) error
	SearchQmsDocumentService(keyword string, limit, offset int) (models.QmsDocumentListResponse, error)
}
