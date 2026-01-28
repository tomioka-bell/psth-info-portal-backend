package services

import (
	"errors"
	"strings"
	"time"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	"backend/internal/repositories"
)

type OrganizationService struct {
	repo *repositories.OrganizationRepository
}

func NewOrganizationService(repo *repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// CreateOrganization creates a new organization
func (s *OrganizationService) CreateOrganization(req *models.CreateOrganizationRequest) (*models.OrganizationResponse, error) {
	// Validation
	if req.Name == "" || req.Desc == "" || req.Category == "" || req.Href == "" {
		return nil, errors.New("name, desc, category, and href are required")
	}

	// Validate category
	validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
	if !validCategories[strings.ToLower(req.Category)] {
		return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
	}

	org := &domains.Organization{
		Name:     strings.TrimSpace(req.Name),
		Desc:     strings.TrimSpace(req.Desc),
		Category: strings.ToLower(req.Category),
		Href:     strings.TrimSpace(req.Href),
		Icon:     strings.TrimSpace(req.Icon),
	}

	err := s.repo.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	return &models.OrganizationResponse{
		ID:        org.ID,
		Name:      org.Name,
		Desc:      org.Desc,
		Category:  org.Category,
		Href:      org.Href,
		Icon:      org.Icon,
		FileName:  org.FileName,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}, nil
}

// GetOrganizationByID retrieves organization by ID
func (s *OrganizationService) GetOrganizationByID(id int) (*models.OrganizationResponse, error) {
	org, err := s.repo.GetOrganizationByID(id)
	if err != nil {
		return nil, err
	}

	return &models.OrganizationResponse{
		ID:        org.ID,
		Name:      org.Name,
		Desc:      org.Desc,
		Category:  org.Category,
		Href:      org.Href,
		Icon:      org.Icon,
		FileName:  org.FileName,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}, nil
}

// GetAllOrganizations retrieves all organizations with pagination
func (s *OrganizationService) GetAllOrganizations(page, pageSize int) (*models.OrganizationListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	orgs, total, err := s.repo.GetAllOrganizations(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.OrganizationListResponse{
		Data:       make([]models.OrganizationResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		resp.Data[i] = models.OrganizationResponse{
			ID:        org.ID,
			Name:      org.Name,
			Desc:      org.Desc,
			Category:  org.Category,
			Href:      org.Href,
			Icon:      org.Icon,
			FileName:  org.FileName,
			CreatedAt: org.CreatedAt,
			UpdatedAt: org.UpdatedAt,
		}
	}

	return resp, nil
}

// GetOrganizationsByCategory retrieves organizations by category
func (s *OrganizationService) GetOrganizationsByCategory(category string, page, pageSize int) (*models.OrganizationListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Validate category
	validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
	category = strings.ToLower(category)
	if !validCategories[category] {
		return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
	}

	orgs, total, err := s.repo.GetOrganizationsByCategory(category, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.OrganizationListResponse{
		Data:       make([]models.OrganizationResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		resp.Data[i] = models.OrganizationResponse{
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

// UpdateOrganization updates an existing organization
func (s *OrganizationService) UpdateOrganization(id int, req *models.UpdateOrganizationRequest) (*models.OrganizationResponse, error) {
	// Verify organization exists
	_, err := s.repo.GetOrganizationByID(id)
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
		validCategories := map[string]bool{"office": true, "production": true, "quality": true, "support": true}
		category := strings.ToLower(req.Category)
		if !validCategories[category] {
			return nil, errors.New("invalid category. Must be one of: office, production, quality, support")
		}
		updates["category"] = category
	}
	if req.Href != "" {
		updates["href"] = strings.TrimSpace(req.Href)
	}
	if req.Icon != "" {
		updates["icon"] = strings.TrimSpace(req.Icon)
	}
	if req.FileName != "" {
		updates["file_name"] = strings.TrimSpace(req.FileName)
	}

	updates["updated_at"] = time.Now()

	err = s.repo.UpdateOrganization(id, updates)
	if err != nil {
		return nil, err
	}

	// Fetch updated organization
	updatedOrg, err := s.repo.GetOrganizationByID(id)
	if err != nil {
		return nil, err
	}

	return &models.OrganizationResponse{
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

// DeleteOrganization soft deletes an organization
func (s *OrganizationService) DeleteOrganization(id int) error {
	// Verify organization exists
	_, err := s.repo.GetOrganizationByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteOrganization(id)
}

// SearchOrganizations searches organizations by keyword
func (s *OrganizationService) SearchOrganizations(keyword string, page, pageSize int) (*models.OrganizationListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return s.GetAllOrganizations(page, pageSize)
	}

	orgs, total, err := s.repo.SearchOrganizations(keyword, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := &models.OrganizationListResponse{
		Data:       make([]models.OrganizationResponse, len(orgs)),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	for i, org := range orgs {
		resp.Data[i] = models.OrganizationResponse{
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

// CreateOrganizationService implements the port interface
func (s *OrganizationService) CreateOrganizationService(req models.CreateOrganizationRequest) error {
	_, err := s.CreateOrganization(&req)
	return err
}

// GetAllOrganizationsService implements the port interface
func (s *OrganizationService) GetAllOrganizationsService(limit, offset int) (models.OrganizationListResponse, error) {
	// Convert offset to page number (offset = (page - 1) * limit)
	page := offset/limit + 1
	resp, err := s.GetAllOrganizations(page, limit)
	if err != nil {
		return models.OrganizationListResponse{}, err
	}
	return *resp, nil
}

// GetOrganizationsByCategoryService implements the port interface
func (s *OrganizationService) GetOrganizationsByCategoryService(category string, limit, offset int) (models.OrganizationListResponse, error) {
	page := offset/limit + 1
	resp, err := s.GetOrganizationsByCategory(category, page, limit)
	if err != nil {
		return models.OrganizationListResponse{}, err
	}
	return *resp, nil
}

// GetOrganizationByIDService implements the port interface
func (s *OrganizationService) GetOrganizationByIDService(id int) (*models.OrganizationResponse, error) {
	return s.GetOrganizationByID(id)
}

// UpdateOrganizationService implements the port interface
func (s *OrganizationService) UpdateOrganizationService(id int, req models.UpdateOrganizationRequest) error {
	_, err := s.UpdateOrganization(id, &req)
	return err
}

// DeleteOrganizationService implements the port interface
func (s *OrganizationService) DeleteOrganizationService(id int) error {
	return s.DeleteOrganization(id)
}

// SearchOrganizationsService implements the port interface
func (s *OrganizationService) SearchOrganizationsService(keyword string, limit, offset int) (models.OrganizationListResponse, error) {
	page := offset/limit + 1
	resp, err := s.SearchOrganizations(keyword, page, limit)
	if err != nil {
		return models.OrganizationListResponse{}, err
	}
	return *resp, nil
}
