package services

import (
	"errors"
	"strings"
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type CustomerManualService struct {
	repo *repositories.CustomerManualRepository
}

func NewCustomerManualService(repo *repositories.CustomerManualRepository) *CustomerManualService {
	return &CustomerManualService{repo: repo}
}

// CreateCustomerManual creates a new customer manual
func (s *CustomerManualService) CreateCustomerManual(req *models.CreateCustomerManualRequest) (*models.CustomerManualResponse, error) {
	// Validation
	if req.CustomerManualName == "" || req.Desc == "" || req.Category == "" {
		return nil, errors.New("customer_manual_name, desc, and category are required")
	}

	// Validate category
	validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
	if !validCategories[strings.ToLower(req.Category)] {
		return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
	}

	manual := &domains.CustomerManual{
		CustomerManualName: strings.TrimSpace(req.CustomerManualName),
		Desc:               strings.TrimSpace(req.Desc),
		Category:           strings.ToLower(req.Category),
		FileName:           strings.TrimSpace(req.FileName),
		ParentID:           req.ParentID,
		SortOrder:          req.SortOrder,
	}

	err := s.repo.CreateCustomerManual(manual)
	if err != nil {
		return nil, err
	}

	return &models.CustomerManualResponse{
		CustomerManualID:   manual.CustomerManualID,
		CustomerManualName: manual.CustomerManualName,
		Desc:               manual.Desc,
		Category:           manual.Category,
		FileName:           manual.FileName,
		CreatedAt:          manual.CreatedAt,
		UpdatedAt:          manual.UpdatedAt,
	}, nil
}

// GetCustomerManualByID retrieves customer manual by ID
func (s *CustomerManualService) GetCustomerManualByID(id int) (*models.CustomerManualResponse, error) {
	manual, err := s.repo.GetCustomerManualByID(id)
	if err != nil {
		return nil, err
	}

	return &models.CustomerManualResponse{
		CustomerManualID:   manual.CustomerManualID,
		CustomerManualName: manual.CustomerManualName,
		Desc:               manual.Desc,
		Category:           manual.Category,
		FileName:           manual.FileName,
		CreatedAt:          manual.CreatedAt,
		UpdatedAt:          manual.UpdatedAt,
	}, nil
}

// GetAllCustomerManuals retrieves all customer manuals with pagination
func (s *CustomerManualService) GetAllCustomerManuals(page, pageSize int) (*models.CustomerManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.GetAllCustomerManuals(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.CustomerManualListResponse{
		Data:       make([]models.CustomerManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.CustomerManualResponse{
			CustomerManualID:   manual.CustomerManualID,
			CustomerManualName: manual.CustomerManualName,
			Desc:               manual.Desc,
			Category:           manual.Category,
			FileName:           manual.FileName,
			ParentID:           manual.ParentID,
			SortOrder:          manual.SortOrder,
			CreatedAt:          manual.CreatedAt,
			UpdatedAt:          manual.UpdatedAt,
		}
	}

	return resp, nil
}

// GetAllCustomerManualsTree retrieves all customer manuals as a hierarchical tree structure
func (s *CustomerManualService) GetAllCustomerManualsTree() ([]domains.CustomerManualMenu, error) {
	// Get all customer manuals without pagination
	page := 1
	pageSize := 10000 // Get all records

	manuals, _, err := s.repo.GetAllCustomerManuals(page, pageSize)
	if err != nil {
		return nil, err
	}

	// Build tree structure
	tree := domains.BuildCustomerManualTree(manuals)
	return tree, nil
}

// GetCustomerManualsByCategory retrieves customer manuals by category
func (s *CustomerManualService) GetCustomerManualsByCategory(category string, page, pageSize int) (*models.CustomerManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.GetCustomerManualsByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.CustomerManualListResponse{
		Data:       make([]models.CustomerManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.CustomerManualResponse{
			CustomerManualID:   manual.CustomerManualID,
			CustomerManualName: manual.CustomerManualName,
			Desc:               manual.Desc,
			Category:           manual.Category,
			FileName:           manual.FileName,
			CreatedAt:          manual.CreatedAt,
			UpdatedAt:          manual.UpdatedAt,
		}
	}

	return resp, nil
}

// UpdateCustomerManual updates an existing customer manual
func (s *CustomerManualService) UpdateCustomerManual(id int, req *models.UpdateCustomerManualRequest) (*models.CustomerManualResponse, error) {
	// Verify manual exists
	_, err := s.repo.GetCustomerManualByID(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if req.CustomerManualName != "" {
		updates["customer_manual_name"] = strings.TrimSpace(req.CustomerManualName)
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

	if req.ClearParentID {
		updates["parent_id"] = nil
	} else if req.ParentID != nil {
		updates["parent_id"] = *req.ParentID
	}
	updates["sort_order"] = req.SortOrder

	updates["updated_at"] = time.Now()

	err = s.repo.UpdateCustomerManual(id, updates)
	if err != nil {
		return nil, err
	}

	// Fetch updated manual
	updatedManual, err := s.repo.GetCustomerManualByID(id)
	if err != nil {
		return nil, err
	}

	return &models.CustomerManualResponse{
		CustomerManualID:   updatedManual.CustomerManualID,
		CustomerManualName: updatedManual.CustomerManualName,
		Desc:               updatedManual.Desc,
		Category:           updatedManual.Category,
		FileName:           updatedManual.FileName,
		CreatedAt:          updatedManual.CreatedAt,
		UpdatedAt:          updatedManual.UpdatedAt,
	}, nil
}

// DeleteCustomerManual soft deletes a customer manual
func (s *CustomerManualService) DeleteCustomerManual(id int) error {
	return s.repo.DeleteCustomerManual(id)
}

// SearchCustomerManuals searches customer manuals by keyword
func (s *CustomerManualService) SearchCustomerManuals(keyword string, page, pageSize int) (*models.CustomerManualListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	manuals, total, err := s.repo.SearchCustomerManuals(keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.CustomerManualListResponse{
		Data:       make([]models.CustomerManualResponse, len(manuals)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, manual := range manuals {
		resp.Data[i] = models.CustomerManualResponse{
			CustomerManualID:   manual.CustomerManualID,
			CustomerManualName: manual.CustomerManualName,
			Desc:               manual.Desc,
			Category:           manual.Category,
			FileName:           manual.FileName,
			CreatedAt:          manual.CreatedAt,
			UpdatedAt:          manual.UpdatedAt,
		}
	}

	return resp, nil
}

// CreateCustomerManualService implements the port interface
func (s *CustomerManualService) CreateCustomerManualService(req models.CreateCustomerManualRequest) error {
	_, err := s.CreateCustomerManual(&req)
	return err
}

// CreateCustomerManualWithResponse creates and returns the response with ID
func (s *CustomerManualService) CreateCustomerManualWithResponse(req models.CreateCustomerManualRequest) (*models.CustomerManualResponse, error) {
	return s.CreateCustomerManual(&req)
}

// GetAllCustomerManualService implements the port interface
func (s *CustomerManualService) GetAllCustomerManualService(limit, offset int) (models.CustomerManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetAllCustomerManuals(page, limit)
	if err != nil {
		return models.CustomerManualListResponse{}, err
	}
	return *resp, nil
}

// GetCustomerManualByCategoryService implements the port interface
func (s *CustomerManualService) GetCustomerManualByCategoryService(category string, limit, offset int) (models.CustomerManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetCustomerManualsByCategory(category, page, limit)
	if err != nil {
		return models.CustomerManualListResponse{}, err
	}
	return *resp, nil
}

// GetCustomerManualByIDService implements the port interface
func (s *CustomerManualService) GetCustomerManualByIDService(id int) (*models.CustomerManualResponse, error) {
	return s.GetCustomerManualByID(id)
}

// UpdateCustomerManualService implements the port interface
func (s *CustomerManualService) UpdateCustomerManualService(id int, req models.UpdateCustomerManualRequest) error {
	_, err := s.UpdateCustomerManual(id, &req)
	return err
}

// DeleteCustomerManualService implements the port interface
func (s *CustomerManualService) DeleteCustomerManualService(id int) error {
	return s.DeleteCustomerManual(id)
}

// SearchCustomerManualService implements the port interface
func (s *CustomerManualService) SearchCustomerManualService(keyword string, limit, offset int) (models.CustomerManualListResponse, error) {
	page := offset/limit + 1
	resp, err := s.SearchCustomerManuals(keyword, page, limit)
	if err != nil {
		return models.CustomerManualListResponse{}, err
	}
	return *resp, nil
}
