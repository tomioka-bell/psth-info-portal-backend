package services

import (
	"errors"
	"strings"
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type AppSystemService struct {
	repo *repositories.AppSystemRepository
}

func NewAppSystemService(repo *repositories.AppSystemRepository) *AppSystemService {
	return &AppSystemService{repo: repo}
}

// CreateAppSystem creates a new AppSystem
func (s *AppSystemService) CreateAppSystem(req *models.CreateAppSystemRequest) (*models.AppSystemResponse, error) {
	// Validation
	if req.Name == "" || req.Desc == "" || req.Category == "" || req.Href == "" {
		return nil, errors.New("name, desc, category, and href are required")
	}

	// Validate category is not empty
	if strings.TrimSpace(req.Category) == "" {
		return nil, errors.New("category is required")
	}

	org := &domains.AppSystem{
		Name:      strings.TrimSpace(req.Name),
		Desc:      strings.TrimSpace(req.Desc),
		Category:  strings.ToLower(req.Category),
		Href:      strings.TrimSpace(req.Href),
		Icon:      strings.TrimSpace(req.Icon),
		ParentID:  req.ParentID,
		SortOrder: req.SortOrder,
	}

	err := s.repo.CreateAppSystem(org)
	if err != nil {
		return nil, err
	}

	return &models.AppSystemResponse{
		ID:        org.ID,
		Name:      org.Name,
		Desc:      org.Desc,
		Category:  org.Category,
		Href:      org.Href,
		Icon:      org.Icon,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}, nil
}

// GetAppSystemByID retrieves AppSystem by ID
func (s *AppSystemService) GetAppSystemByID(id int) (*models.AppSystemResponse, error) {
	org, err := s.repo.GetAppSystemByID(id)
	if err != nil {
		return nil, err
	}

	return &models.AppSystemResponse{
		ID:        org.ID,
		Name:      org.Name,
		Desc:      org.Desc,
		Category:  org.Category,
		Href:      org.Href,
		Icon:      org.Icon,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}, nil
}

// GetAllAppSystems retrieves all AppSystems with pagination
func (s *AppSystemService) GetAllAppSystems(page, pageSize int) (*models.AppSystemListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	orgs, total, err := s.repo.GetAllAppSystems(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.AppSystemListResponse{
		Data:       make([]models.AppSystemResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		var parentID int
		if org.ParentID != nil {
			parentID = *org.ParentID
		} else {
			parentID = 0 // root
		}
		resp.Data[i] = models.AppSystemResponse{
			ID:        org.ID,
			Name:      org.Name,
			Desc:      org.Desc,
			Category:  org.Category,
			Href:      org.Href,
			Icon:      org.Icon,
			ParentID:  parentID,
			SortOrder: org.SortOrder,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		}
	}

	return resp, nil
}

// GetAllAppSystemsTree retrieves all AppSystems as a hierarchical tree structure
func (s *AppSystemService) GetAllAppSystemsTree() ([]domains.AppSystemMenu, error) {
	// Get all AppSystems without pagination
	page := 1
	pageSize := 10000 // Get all records

	orgs, _, err := s.repo.GetAllAppSystems(page, pageSize)
	if err != nil {
		return nil, err
	}

	// Build tree structure
	tree := domains.BuildAppSystemTree(orgs)
	return tree, nil
}

// GetAppSystemsByCategory retrieves AppSystems by category
func (s *AppSystemService) GetAppSystemsByCategory(category string, page, pageSize int) (*models.AppSystemListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Validate category is not empty
	category = strings.ToLower(strings.TrimSpace(category))
	if category == "" {
		return nil, errors.New("category is required")
	}

	orgs, total, err := s.repo.GetAppSystemsByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.AppSystemListResponse{
		Data:       make([]models.AppSystemResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		resp.Data[i] = models.AppSystemResponse{
			ID:        org.ID,
			Name:      org.Name,
			Desc:      org.Desc,
			Category:  org.Category,
			Href:      org.Href,
			Icon:      org.Icon,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		}
	}

	return resp, nil
}

// UpdateAppSystem updates an existing AppSystem
func (s *AppSystemService) UpdateAppSystem(id int, req *models.UpdateAppSystemRequest) (*models.AppSystemResponse, error) {
	// Verify AppSystem exists
	_, err := s.repo.GetAppSystemByID(id)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = strings.TrimSpace(req.Name)
	}
	if req.Desc != "" {
		updates["desc"] = strings.TrimSpace(req.Desc)
	}
	if req.Category != "" {
		category := strings.ToLower(strings.TrimSpace(req.Category))
		if category != "" {
			updates["category"] = category
		}
	}
	if req.Href != "" {
		updates["href"] = strings.TrimSpace(req.Href)
	}
	if req.Icon != "" {
		updates["icon"] = strings.TrimSpace(req.Icon)
	}

	// Add parent_id and sort_order updates
	// If ClearParentID is true, set parent_id to NULL; otherwise update if ParentID is not nil
	if req.ClearParentID {
		updates["parent_id"] = nil
	} else if req.ParentID != nil {
		updates["parent_id"] = *req.ParentID
	}
	updates["sort_order"] = req.SortOrder

	updates["updated_at"] = time.Now()

	err = s.repo.UpdateAppSystem(id, updates)
	if err != nil {
		return nil, err
	}

	// Fetch updated AppSystem
	updatedOrg, err := s.repo.GetAppSystemByID(id)
	if err != nil {
		return nil, err
	}

	return &models.AppSystemResponse{
		ID:        updatedOrg.ID,
		Name:      updatedOrg.Name,
		Desc:      updatedOrg.Desc,
		Category:  updatedOrg.Category,
		Href:      updatedOrg.Href,
		Icon:      updatedOrg.Icon,
		CreatedAt: updatedOrg.CreatedAt,
		UpdatedAt: updatedOrg.UpdatedAt,
	}, nil
}

// DeleteAppSystem soft deletes an AppSystem
func (s *AppSystemService) DeleteAppSystem(id int) error {
	// Verify AppSystem exists
	_, err := s.repo.GetAppSystemByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteAppSystem(id)
}

// SearchAppSystems searches AppSystems by keyword
func (s *AppSystemService) SearchAppSystems(keyword string, page, pageSize int) (*models.AppSystemListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return s.GetAllAppSystems(page, pageSize)
	}

	orgs, total, err := s.repo.SearchAppSystems(keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.AppSystemListResponse{
		Data:       make([]models.AppSystemResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		resp.Data[i] = models.AppSystemResponse{
			ID:        org.ID,
			Name:      org.Name,
			Desc:      org.Desc,
			Category:  org.Category,
			Href:      org.Href,
			Icon:      org.Icon,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		}
	}

	return resp, nil
}

// CreateAppSystemService implements the port interface
func (s *AppSystemService) CreateAppSystemService(req models.CreateAppSystemRequest) error {
	_, err := s.CreateAppSystem(&req)
	return err
}

// CreateAppSystemWithResponse creates and returns the response with ID
func (s *AppSystemService) CreateAppSystemWithResponse(req models.CreateAppSystemRequest) (*models.AppSystemResponse, error) {
	return s.CreateAppSystem(&req)
}

// GetAllAppSystemsService implements the port interface
func (s *AppSystemService) GetAllAppSystemsService(limit, offset int) (models.AppSystemListResponse, error) {
	// Convert offset to page number (offset = (page - 1) * limit)
	page := offset/limit + 1
	resp, err := s.GetAllAppSystems(page, limit)
	if err != nil {
		return models.AppSystemListResponse{}, err
	}
	return *resp, nil
}

// GetAppSystemsByCategoryService implements the port interface
func (s *AppSystemService) GetAppSystemsByCategoryService(category string, limit, offset int) (models.AppSystemListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetAppSystemsByCategory(category, page, limit)
	if err != nil {
		return models.AppSystemListResponse{}, err
	}
	return *resp, nil
}

// GetAppSystemByIDService implements the port interface
func (s *AppSystemService) GetAppSystemByIDService(id int) (*models.AppSystemResponse, error) {
	return s.GetAppSystemByID(id)
}

// UpdateAppSystemService implements the port interface
func (s *AppSystemService) UpdateAppSystemService(id int, req models.UpdateAppSystemRequest) error {
	_, err := s.UpdateAppSystem(id, &req)
	return err
}

// DeleteAppSystemService implements the port interface
func (s *AppSystemService) DeleteAppSystemService(id int) error {
	return s.DeleteAppSystem(id)
}

// SearchAppSystemsService implements the port interface
func (s *AppSystemService) SearchAppSystemsService(keyword string, limit, offset int) (models.AppSystemListResponse, error) {
	page := offset/limit + 1
	resp, err := s.SearchAppSystems(keyword, page, limit)
	if err != nil {
		return models.AppSystemListResponse{}, err
	}
	return *resp, nil
}
