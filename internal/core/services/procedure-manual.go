package services

import (
	"errors"
	"strings"
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type ProcedureManualService struct {
	repo *repositories.ProcedureManualRepository
}

func NewProcedureManualService(repo *repositories.ProcedureManualRepository) *ProcedureManualService {
	return &ProcedureManualService{repo: repo}
}

// CreateProcedureManual creates a new procedure manual
func (s *ProcedureManualService) CreateProcedureManual(req *models.CreateProcedureManualRequest) (*models.ProcedureManualResponse, error) {
	// Validation
	if req.ProcedureManualName == "" || req.Desc == "" || req.Category == "" {
		return nil, errors.New("procedure_manual_name, desc, and category are required")
	}

	// Validate category
	validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
	if !validCategories[strings.ToLower(req.Category)] {
		return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
	}

	manual := &domains.ProcedureManual{
		ProcedureManualName: strings.TrimSpace(req.ProcedureManualName),
		Desc:                strings.TrimSpace(req.Desc),
		Category:            strings.ToLower(req.Category),
		FileName:            strings.TrimSpace(req.FileName),
	}

	err := s.repo.CreateProcedureManual(manual)
	if err != nil {
		return nil, err
	}

	return &models.ProcedureManualResponse{
		ProcedureManualID:   manual.ProcedureManualID,
		ProcedureManualName: manual.ProcedureManualName,
		Desc:                manual.Desc,
		Category:            manual.Category,
		FileName:            manual.FileName,
		CreatedAt:           manual.CreatedAt,
		UpdatedAt:           manual.UpdatedAt,
	}, nil
}

// GetProcedureManualByID retrieves procedure manual by ID
func (s *ProcedureManualService) GetProcedureManualByID(id int) (*models.ProcedureManualResponse, error) {
	manual, err := s.repo.GetProcedureManualByID(id)
	if err != nil {
		return nil, err
	}

	return &models.ProcedureManualResponse{
		ProcedureManualID:   manual.ProcedureManualID,
		ProcedureManualName: manual.ProcedureManualName,
		Desc:                manual.Desc,
		Category:            manual.Category,
		FileName:            manual.FileName,
		CreatedAt:           manual.CreatedAt,
		UpdatedAt:           manual.UpdatedAt,
	}, nil
}

// GetAllProcedureManuals retrieves all procedure manuals with pagination
func (s *ProcedureManualService) GetAllProcedureManuals(page, pageSize int) (*models.ProcedureManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.GetAllProcedureManuals(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.ProcedureManualListResponse{
		Data:       make([]models.ProcedureManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.ProcedureManualResponse{
			ProcedureManualID:   manual.ProcedureManualID,
			ProcedureManualName: manual.ProcedureManualName,
			Desc:                manual.Desc,
			Category:            manual.Category,
			FileName:            manual.FileName,
			CreatedAt:           manual.CreatedAt,
			UpdatedAt:           manual.UpdatedAt,
		}
	}

	return resp, nil
}

// GetProcedureManualsByCategory retrieves procedure manuals by category
func (s *ProcedureManualService) GetProcedureManualsByCategory(category string, page, pageSize int) (*models.ProcedureManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.GetProcedureManualsByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.ProcedureManualListResponse{
		Data:       make([]models.ProcedureManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.ProcedureManualResponse{
			ProcedureManualID:   manual.ProcedureManualID,
			ProcedureManualName: manual.ProcedureManualName,
			Desc:                manual.Desc,
			Category:            manual.Category,
			FileName:            manual.FileName,
			CreatedAt:           manual.CreatedAt,
			UpdatedAt:           manual.UpdatedAt,
		}
	}

	return resp, nil
}

// UpdateProcedureManual updates an existing procedure manual
func (s *ProcedureManualService) UpdateProcedureManual(id int, req *models.UpdateProcedureManualRequest) (*models.ProcedureManualResponse, error) {
	// Verify manual exists
	_, err := s.repo.GetProcedureManualByID(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if req.ProcedureManualName != "" {
		updates["procedure_manual_name"] = strings.TrimSpace(req.ProcedureManualName)
	}
	if req.Desc != "" {
		updates["desc"] = strings.TrimSpace(req.Desc)
	}
	if req.Category != "" {
		validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
		category := strings.ToLower(req.Category)
		if !validCategories[category] {
			return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
		}
		updates["category"] = category
	}
	if req.FileName != "" {
		updates["file_name"] = strings.TrimSpace(req.FileName)
	}

	updates["updated_at"] = time.Now()

	err = s.repo.UpdateProcedureManual(id, updates)
	if err != nil {
		return nil, err
	}

	// Fetch updated manual
	updatedManual, err := s.repo.GetProcedureManualByID(id)
	if err != nil {
		return nil, err
	}

	return &models.ProcedureManualResponse{
		ProcedureManualID:   updatedManual.ProcedureManualID,
		ProcedureManualName: updatedManual.ProcedureManualName,
		Desc:                updatedManual.Desc,
		Category:            updatedManual.Category,
		FileName:            updatedManual.FileName,
		CreatedAt:           updatedManual.CreatedAt,
		UpdatedAt:           updatedManual.UpdatedAt,
	}, nil
}

// DeleteProcedureManual soft deletes a procedure manual
func (s *ProcedureManualService) DeleteProcedureManual(id int) error {
	return s.repo.DeleteProcedureManual(id)
}

// SearchProcedureManuals searches procedure manuals by keyword
func (s *ProcedureManualService) SearchProcedureManuals(keyword string, page, pageSize int) (*models.ProcedureManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.SearchProcedureManuals(keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.ProcedureManualListResponse{
		Data:       make([]models.ProcedureManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.ProcedureManualResponse{
			ProcedureManualID:   manual.ProcedureManualID,
			ProcedureManualName: manual.ProcedureManualName,
			Desc:                manual.Desc,
			Category:            manual.Category,
			FileName:            manual.FileName,
			CreatedAt:           manual.CreatedAt,
			UpdatedAt:           manual.UpdatedAt,
		}
	}

	return resp, nil
}

// CreateProcedureManualService implements the port interface
func (s *ProcedureManualService) CreateProcedureManualService(req models.CreateProcedureManualRequest) error {
	_, err := s.CreateProcedureManual(&req)
	return err
}

// GetAllProcedureManualService implements the port interface
func (s *ProcedureManualService) GetAllProcedureManualService(limit, offset int) (models.ProcedureManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetAllProcedureManuals(page, limit)
	if err != nil {
		return models.ProcedureManualListResponse{}, err
	}
	return *resp, nil
}

// GetProcedureManualByCategoryService implements the port interface
func (s *ProcedureManualService) GetProcedureManualByCategoryService(category string, limit, offset int) (models.ProcedureManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetProcedureManualsByCategory(category, page, limit)
	if err != nil {
		return models.ProcedureManualListResponse{}, err
	}
	return *resp, nil
}

// GetProcedureManualByIDService implements the port interface
func (s *ProcedureManualService) GetProcedureManualByIDService(id int) (*models.ProcedureManualResponse, error) {
	return s.GetProcedureManualByID(id)
}

// UpdateProcedureManualService implements the port interface
func (s *ProcedureManualService) UpdateProcedureManualService(id int, req models.UpdateProcedureManualRequest) error {
	_, err := s.UpdateProcedureManual(id, &req)
	return err
}

// DeleteProcedureManualService implements the port interface
func (s *ProcedureManualService) DeleteProcedureManualService(id int) error {
	return s.DeleteProcedureManual(id)
}

// SearchProcedureManualService implements the port interface
func (s *ProcedureManualService) SearchProcedureManualService(keyword string, limit, offset int) (models.ProcedureManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.SearchProcedureManuals(keyword, page, limit)
	if err != nil {
		return models.ProcedureManualListResponse{}, err
	}
	return *resp, nil
}
