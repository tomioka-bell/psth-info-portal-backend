package services

import (
	"errors"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	portRepositories "backend/internal/core/ports/repositories"
	portServices "backend/internal/core/ports/services"
)

type SafetyDocumentService struct {
	repo portRepositories.SafetyDocumentRepository
}

func NewSafetyDocumentService(repo portRepositories.SafetyDocumentRepository) portServices.SafetyDocumentService {
	return &SafetyDocumentService{repo: repo}
}

func (s *SafetyDocumentService) CreateSafetyDocument(req *models.CreateSafetyDocumentRequest) (*models.SafetyDocumentResponse, error) {
	// Validate required fields
	if req.SafetyDocumentName == "" {
		return nil, errors.New("safety_document_name is required")
	}
	if req.SafetyDocumentDesc == "" {
		return nil, errors.New("safety_document_desc is required")
	}
	if req.Category == "" {
		return nil, errors.New("category is required")
	}
	if req.Department == "" {
		return nil, errors.New("department is required")
	}

	doc := domains.SafetyDocument{
		SafetyDocumentName: req.SafetyDocumentName,
		SafetyDocumentDesc: req.SafetyDocumentDesc,
		Category:           req.Category,
		Department:         req.Department,
		FileName:           req.FileName,
	}

	if err := s.repo.CreateSafetyDocument(&doc); err != nil {
		return nil, err
	}

	return &models.SafetyDocumentResponse{
		SafetyDocumentID:   doc.SafetyDocumentID,
		SafetyDocumentName: doc.SafetyDocumentName,
		SafetyDocumentDesc: doc.SafetyDocumentDesc,
		Category:           doc.Category,
		Department:         doc.Department,
		FileName:           doc.FileName,
		CreatedAt:          doc.CreatedAt,
		UpdatedAt:          doc.UpdatedAt,
	}, nil
}

func (s *SafetyDocumentService) GetSafetyDocumentByID(id int) (*models.SafetyDocumentResponse, error) {
	doc, err := s.repo.GetSafetyDocumentByID(id)
	if err != nil {
		return nil, err
	}

	return &models.SafetyDocumentResponse{
		SafetyDocumentID:   doc.SafetyDocumentID,
		SafetyDocumentName: doc.SafetyDocumentName,
		SafetyDocumentDesc: doc.SafetyDocumentDesc,
		Category:           doc.Category,
		Department:         doc.Department,
		FileName:           doc.FileName,
		CreatedAt:          doc.CreatedAt,
		UpdatedAt:          doc.UpdatedAt,
	}, nil
}

func (s *SafetyDocumentService) GetAllSafetyDocuments(limit, offset int) (*models.SafetyDocumentListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.GetAllSafetyDocuments(limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.SafetyDocumentResponse
	for _, doc := range docs {
		responses = append(responses, models.SafetyDocumentResponse{
			SafetyDocumentID:   doc.SafetyDocumentID,
			SafetyDocumentName: doc.SafetyDocumentName,
			SafetyDocumentDesc: doc.SafetyDocumentDesc,
			Category:           doc.Category,
			Department:         doc.Department,
			FileName:           doc.FileName,
			CreatedAt:          doc.CreatedAt,
			UpdatedAt:          doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.SafetyDocumentListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *SafetyDocumentService) GetSafetyDocumentByCategory(category string, limit, offset int) (*models.SafetyDocumentListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.GetSafetyDocumentsByCategory(category, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.SafetyDocumentResponse
	for _, doc := range docs {
		responses = append(responses, models.SafetyDocumentResponse{
			SafetyDocumentID:   doc.SafetyDocumentID,
			SafetyDocumentName: doc.SafetyDocumentName,
			SafetyDocumentDesc: doc.SafetyDocumentDesc,
			Category:           doc.Category,
			Department:         doc.Department,
			FileName:           doc.FileName,
			CreatedAt:          doc.CreatedAt,
			UpdatedAt:          doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.SafetyDocumentListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *SafetyDocumentService) GetSafetyDocumentByDepartment(department string, limit, offset int) (*models.SafetyDocumentListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.GetSafetyDocumentsByDepartment(department, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.SafetyDocumentResponse
	for _, doc := range docs {
		responses = append(responses, models.SafetyDocumentResponse{
			SafetyDocumentID:   doc.SafetyDocumentID,
			SafetyDocumentName: doc.SafetyDocumentName,
			SafetyDocumentDesc: doc.SafetyDocumentDesc,
			Category:           doc.Category,
			Department:         doc.Department,
			FileName:           doc.FileName,
			CreatedAt:          doc.CreatedAt,
			UpdatedAt:          doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.SafetyDocumentListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *SafetyDocumentService) UpdateSafetyDocument(id int, req *models.UpdateSafetyDocumentRequest) (*models.SafetyDocumentResponse, error) {
	// Validate
	if req.SafetyDocumentName == "" {
		return nil, errors.New("safety_document_name is required")
	}
	if req.SafetyDocumentDesc == "" {
		return nil, errors.New("safety_document_desc is required")
	}
	if req.Category == "" {
		return nil, errors.New("category is required")
	}
	if req.Department == "" {
		return nil, errors.New("department is required")
	}

	updates := map[string]interface{}{
		"safety_document_name": req.SafetyDocumentName,
		"safety_document_desc": req.SafetyDocumentDesc,
		"category":             req.Category,
		"department":           req.Department,
	}

	if req.FileName != "" {
		updates["file_name"] = req.FileName
	}

	if err := s.repo.UpdateSafetyDocument(id, updates); err != nil {
		return nil, err
	}

	// Get updated doc
	return s.GetSafetyDocumentByID(id)
}

func (s *SafetyDocumentService) DeleteSafetyDocument(id int) error {
	return s.repo.DeleteSafetyDocument(id)
}

func (s *SafetyDocumentService) SearchSafetyDocuments(keyword string, limit, offset int) (*models.SafetyDocumentListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.SearchSafetyDocuments(keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.SafetyDocumentResponse
	for _, doc := range docs {
		responses = append(responses, models.SafetyDocumentResponse{
			SafetyDocumentID:   doc.SafetyDocumentID,
			SafetyDocumentName: doc.SafetyDocumentName,
			SafetyDocumentDesc: doc.SafetyDocumentDesc,
			Category:           doc.Category,
			Department:         doc.Department,
			FileName:           doc.FileName,
			CreatedAt:          doc.CreatedAt,
			UpdatedAt:          doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.SafetyDocumentListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

// ===== Port Interface Implementation =====

// CreateSafetyDocumentService implements the port interface
func (s *SafetyDocumentService) CreateSafetyDocumentService(req models.CreateSafetyDocumentRequest) error {
	_, err := s.CreateSafetyDocument(&req)
	return err
}

// GetAllSafetyDocumentService implements the port interface
func (s *SafetyDocumentService) GetAllSafetyDocumentService(page, pageSize int) (*models.SafetyDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetAllSafetyDocuments(pageSize, offset)
}

// GetSafetyDocumentByIDService implements the port interface
func (s *SafetyDocumentService) GetSafetyDocumentByIDService(id int) (*models.SafetyDocumentResponse, error) {
	return s.GetSafetyDocumentByID(id)
}

// GetSafetyDocumentByCategoryService implements the port interface
func (s *SafetyDocumentService) GetSafetyDocumentByCategoryService(category string, page, pageSize int) (*models.SafetyDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetSafetyDocumentByCategory(category, pageSize, offset)
}

// GetSafetyDocumentByDepartmentService implements the port interface
func (s *SafetyDocumentService) GetSafetyDocumentByDepartmentService(department string, page, pageSize int) (*models.SafetyDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetSafetyDocumentByDepartment(department, pageSize, offset)
}

// UpdateSafetyDocumentService implements the port interface
func (s *SafetyDocumentService) UpdateSafetyDocumentService(id int, req models.UpdateSafetyDocumentRequest) error {
	_, err := s.UpdateSafetyDocument(id, &req)
	return err
}

// DeleteSafetyDocumentService implements the port interface
func (s *SafetyDocumentService) DeleteSafetyDocumentService(id int) error {
	return s.DeleteSafetyDocument(id)
}

// SearchSafetyDocumentService implements the port interface
func (s *SafetyDocumentService) SearchSafetyDocumentService(keyword string, page, pageSize int) (*models.SafetyDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.SearchSafetyDocuments(keyword, pageSize, offset)
}
