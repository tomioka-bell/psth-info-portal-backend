package services

import (
	"errors"
	"strings"
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type QmsDocumentsService struct {
	repo *repositories.QmsDocumentsRepository
}

func NewQmsDocumentsService(repo *repositories.QmsDocumentsRepository) *QmsDocumentsService {
	return &QmsDocumentsService{repo: repo}
}

func (s *QmsDocumentsService) CreateQmsDocument(req *models.CreateQmsDocumentRequest) (*models.QmsDocumentResponse, error) {
	if req.QmsDocumentsName == "" || req.DQmsDocumentsDesc == "" || req.Category == "" {
		return nil, errors.New("qms_documents_name, dqms_documents_desc, and category are required")
	}

	// optional: validate categories
	doc := &domains.QmsDocuments{
		QmsDocumentsName:  strings.TrimSpace(req.QmsDocumentsName),
		DQmsDocumentsDesc: strings.TrimSpace(req.DQmsDocumentsDesc),
		Category:          strings.ToLower(req.Category),
		FileName:          strings.TrimSpace(req.FileName),
	}

	if err := s.repo.CreateQmsDocument(doc); err != nil {
		return nil, err
	}

	return &models.QmsDocumentResponse{
		QmsDocumentsID:    doc.QmsDocumentsID,
		QmsDocumentsName:  doc.QmsDocumentsName,
		DQmsDocumentsDesc: doc.DQmsDocumentsDesc,
		Category:          doc.Category,
		FileName:          doc.FileName,
		CreatedAt:         doc.CreatedAt,
		UpdatedAt:         doc.UpdatedAt,
	}, nil
}

func (s *QmsDocumentsService) GetQmsDocumentByID(id int) (*models.QmsDocumentResponse, error) {
	doc, err := s.repo.GetQmsDocumentByID(id)
	if err != nil {
		return nil, err
	}
	return &models.QmsDocumentResponse{
		QmsDocumentsID:    doc.QmsDocumentsID,
		QmsDocumentsName:  doc.QmsDocumentsName,
		DQmsDocumentsDesc: doc.DQmsDocumentsDesc,
		Category:          doc.Category,
		FileName:          doc.FileName,
		CreatedAt:         doc.CreatedAt,
		UpdatedAt:         doc.UpdatedAt,
	}, nil
}

func (s *QmsDocumentsService) GetAllQmsDocuments(page, pageSize int) (*models.QmsDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	docs, total, err := s.repo.GetAllQmsDocuments(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.QmsDocumentListResponse{
		Data:       make([]models.QmsDocumentResponse, len(docs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, d := range docs {
		resp.Data[i] = models.QmsDocumentResponse{
			QmsDocumentsID:    d.QmsDocumentsID,
			QmsDocumentsName:  d.QmsDocumentsName,
			DQmsDocumentsDesc: d.DQmsDocumentsDesc,
			Category:          d.Category,
			FileName:          d.FileName,
			CreatedAt:         d.CreatedAt,
			UpdatedAt:         d.UpdatedAt,
		}
	}

	return resp, nil
}

func (s *QmsDocumentsService) UpdateQmsDocument(id int, req *models.UpdateQmsDocumentRequest) (*models.QmsDocumentResponse, error) {
	_, err := s.repo.GetQmsDocumentByID(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if req.QmsDocumentsName != "" {
		updates["qms_documents_name"] = strings.TrimSpace(req.QmsDocumentsName)
	}
	if req.DQmsDocumentsDesc != "" {
		// DB column for description follows the same pattern as procedure-manual ('desc')
		updates["desc"] = strings.TrimSpace(req.DQmsDocumentsDesc)
	}
	if req.Category != "" {
		updates["category"] = strings.ToLower(req.Category)
	}
	if req.FileName != "" {
		updates["file_name"] = strings.TrimSpace(req.FileName)
	}

	updates["updated_at"] = time.Now()

	if err := s.repo.UpdateQmsDocument(id, updates); err != nil {
		return nil, err
	}

	updated, err := s.repo.GetQmsDocumentByID(id)
	if err != nil {
		return nil, err
	}

	return &models.QmsDocumentResponse{
		QmsDocumentsID:    updated.QmsDocumentsID,
		QmsDocumentsName:  updated.QmsDocumentsName,
		DQmsDocumentsDesc: updated.DQmsDocumentsDesc,
		Category:          updated.Category,
		FileName:          updated.FileName,
		CreatedAt:         updated.CreatedAt,
		UpdatedAt:         updated.UpdatedAt,
	}, nil
}

func (s *QmsDocumentsService) DeleteQmsDocument(id int) error {
	return s.repo.DeleteQmsDocument(id)
}

func (s *QmsDocumentsService) SearchQmsDocuments(keyword string, page, pageSize int) (*models.QmsDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	docs, total, err := s.repo.SearchQmsDocuments(keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.QmsDocumentListResponse{
		Data:       make([]models.QmsDocumentResponse, len(docs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, d := range docs {
		resp.Data[i] = models.QmsDocumentResponse{
			QmsDocumentsID:    d.QmsDocumentsID,
			QmsDocumentsName:  d.QmsDocumentsName,
			DQmsDocumentsDesc: d.DQmsDocumentsDesc,
			Category:          d.Category,
			FileName:          d.FileName,
			CreatedAt:         d.CreatedAt,
			UpdatedAt:         d.UpdatedAt,
		}
	}

	return resp, nil
}

// Implement port interface wrappers if needed
func (s *QmsDocumentsService) CreateQmsDocumentService(req models.CreateQmsDocumentRequest) error {
	_, err := s.CreateQmsDocument(&req)
	return err
}

func (s *QmsDocumentsService) GetAllQmsDocumentService(limit, offset int) (models.QmsDocumentListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetAllQmsDocuments(page, limit)
	if err != nil {
		return models.QmsDocumentListResponse{}, err
	}
	return *resp, nil
}

func (s *QmsDocumentsService) GetQmsDocumentByIDService(id int) (*models.QmsDocumentResponse, error) {
	return s.GetQmsDocumentByID(id)
}

func (s *QmsDocumentsService) UpdateQmsDocumentService(id int, req models.UpdateQmsDocumentRequest) error {
	_, err := s.UpdateQmsDocument(id, &req)
	return err
}

func (s *QmsDocumentsService) DeleteQmsDocumentService(id int) error {
	return s.DeleteQmsDocument(id)
}

func (s *QmsDocumentsService) SearchQmsDocumentService(keyword string, limit, offset int) (models.QmsDocumentListResponse, error) {
	page := offset/limit + 1
	resp, err := s.SearchQmsDocuments(keyword, page, limit)
	if err != nil {
		return models.QmsDocumentListResponse{}, err
	}
	return *resp, nil
}
