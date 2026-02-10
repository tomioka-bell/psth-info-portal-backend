package services

import (
	"errors"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	portRepositories "backend/internal/core/ports/repositories"
	portServices "backend/internal/core/ports/services"
)

type WelfareBenefitService struct {
	repo portRepositories.WelfareBenefitRepository
}

func NewWelfareBenefitService(repo portRepositories.WelfareBenefitRepository) portServices.WelfareBenefitService {
	return &WelfareBenefitService{repo: repo}
}

func (s *WelfareBenefitService) CreateWelfareBenefit(req *models.CreateWelfareBenefitRequest) (*models.WelfareBenefitResponse, error) {
	// Validate required fields
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.Description == "" {
		return nil, errors.New("description is required")
	}
	if req.Category == "" {
		return nil, errors.New("category is required")
	}

	benefit := domains.WelfareBenefit{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		FileName:    req.FileName,
		ImageURL:    req.ImageURL,
	}

	if err := s.repo.CreateWelfareBenefit(&benefit); err != nil {
		return nil, err
	}

	return &models.WelfareBenefitResponse{
		WelfareBenefitID: benefit.WelfareBenefitID,
		Title:            benefit.Title,
		ImageURL:         benefit.ImageURL,
		Description:      benefit.Description,
		Category:         benefit.Category,
		FileName:         benefit.FileName,
		CreatedAt:        benefit.CreatedAt,
		UpdatedAt:        benefit.UpdatedAt,
	}, nil
}

func (s *WelfareBenefitService) GetWelfareBenefitByID(id int) (*models.WelfareBenefitResponse, error) {
	benefit, err := s.repo.GetWelfareBenefitByID(id)
	if err != nil {
		return nil, err
	}

	return &models.WelfareBenefitResponse{
		WelfareBenefitID: benefit.WelfareBenefitID,
		Title:            benefit.Title,
		ImageURL:         benefit.ImageURL,
		Description:      benefit.Description,
		Category:         benefit.Category,
		FileName:         benefit.FileName,
		CreatedAt:        benefit.CreatedAt,
		UpdatedAt:        benefit.UpdatedAt,
	}, nil
}

func (s *WelfareBenefitService) GetAllWelfareBenefits(limit, offset int) (*models.WelfareBenefitListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	benefits, total, err := s.repo.GetAllWelfareBenefits(limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.WelfareBenefitResponse
	for _, benefit := range benefits {
		responses = append(responses, models.WelfareBenefitResponse{
			WelfareBenefitID: benefit.WelfareBenefitID,
			Title:            benefit.Title,
			ImageURL:         benefit.ImageURL,
			Description:      benefit.Description,
			Category:         benefit.Category,
			FileName:         benefit.FileName,
			CreatedAt:        benefit.CreatedAt,
			UpdatedAt:        benefit.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.WelfareBenefitListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *WelfareBenefitService) GetWelfareBenefitByCategory(category string, limit, offset int) (*models.WelfareBenefitListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	benefits, total, err := s.repo.GetWelfareBenefitsByCategory(category, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.WelfareBenefitResponse
	for _, benefit := range benefits {
		responses = append(responses, models.WelfareBenefitResponse{
			WelfareBenefitID: benefit.WelfareBenefitID,
			Title:            benefit.Title,
			ImageURL:         benefit.ImageURL,
			Description:      benefit.Description,
			Category:         benefit.Category,
			FileName:         benefit.FileName,
			CreatedAt:        benefit.CreatedAt,
			UpdatedAt:        benefit.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.WelfareBenefitListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *WelfareBenefitService) UpdateWelfareBenefit(id int, req *models.UpdateWelfareBenefitRequest) (*models.WelfareBenefitResponse, error) {
	// Validate
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.Description == "" {
		return nil, errors.New("description is required")
	}
	if req.Category == "" {
		return nil, errors.New("category is required")
	}

	updates := map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"category":    req.Category,
	}

	if req.FileName != "" {
		updates["file_name"] = req.FileName
	}

	if req.ImageURL != "" {
		updates["image_url"] = req.ImageURL
	}

	if err := s.repo.UpdateWelfareBenefit(id, updates); err != nil {
		return nil, err
	}

	// Get updated benefit
	return s.GetWelfareBenefitByID(id)
}

func (s *WelfareBenefitService) DeleteWelfareBenefit(id int) error {
	return s.repo.DeleteWelfareBenefit(id)
}

func (s *WelfareBenefitService) SearchWelfareBenefits(keyword string, limit, offset int) (*models.WelfareBenefitListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	benefits, total, err := s.repo.SearchWelfareBenefits(keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.WelfareBenefitResponse
	for _, benefit := range benefits {
		responses = append(responses, models.WelfareBenefitResponse{
			WelfareBenefitID: benefit.WelfareBenefitID,
			Title:            benefit.Title,
			Description:      benefit.Description,
			Category:         benefit.Category,
			FileName:         benefit.FileName,
			CreatedAt:        benefit.CreatedAt,
			UpdatedAt:        benefit.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.WelfareBenefitListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

// Get all welfare benefits without pagination (service helper)
func (s *WelfareBenefitService) GetAllWelfareBenefitsNoPagination() (*models.WelfareBenefitListResponse, error) {
	benefits, total, err := s.repo.GetAllWelfareBenefitsNoPagination()
	if err != nil {
		return nil, err
	}

	var responses []models.WelfareBenefitResponse
	for _, benefit := range benefits {
		responses = append(responses, models.WelfareBenefitResponse{
			WelfareBenefitID: benefit.WelfareBenefitID,
			Title:            benefit.Title,
			ImageURL:         benefit.ImageURL,
			Description:      benefit.Description,
			Category:         benefit.Category,
			FileName:         benefit.FileName,
			CreatedAt:        benefit.CreatedAt,
			UpdatedAt:        benefit.UpdatedAt,
		})
	}

	return &models.WelfareBenefitListResponse{
		Data:       responses,
		Total:      total,
		Page:       1,
		PageSize:   int(total),
		TotalPages: 1,
	}, nil
}

// Get total count of welfare benefits
func (s *WelfareBenefitService) GetWelfareBenefitsCount() (int64, error) {
	// delegate to repository count
	return s.repo.GetWelfareBenefitsCount()
}

// ===== Port Interface Implementation =====

// CreateWelfareBenefitService implements the port interface
func (s *WelfareBenefitService) CreateWelfareBenefitService(req models.CreateWelfareBenefitRequest) error {
	_, err := s.CreateWelfareBenefit(&req)
	return err
}

// GetAllWelfareBenefitService implements the port interface
func (s *WelfareBenefitService) GetAllWelfareBenefitService(page, pageSize int) (*models.WelfareBenefitListResponse, error) {
	// Return all welfare benefits (no pagination).
	// Ignore incoming page/pageSize and fetch all records.
	const largeLimit = 1000000
	return s.GetAllWelfareBenefits(largeLimit, 0)
}

// GetWelfareBenefitByIDService implements the port interface
func (s *WelfareBenefitService) GetWelfareBenefitByIDService(id int) (*models.WelfareBenefitResponse, error) {
	return s.GetWelfareBenefitByID(id)
}

// GetWelfareBenefitByCategoryService implements the port interface
func (s *WelfareBenefitService) GetWelfareBenefitByCategoryService(category string, page, pageSize int) (*models.WelfareBenefitListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetWelfareBenefitByCategory(category, pageSize, offset)
}

// GetAllWelfareBenefitsNoPaginationService implements the port interface
func (s *WelfareBenefitService) GetAllWelfareBenefitsNoPaginationService() (*models.WelfareBenefitListResponse, error) {
	return s.GetAllWelfareBenefitsNoPagination()
}

// UpdateWelfareBenefitService implements the port interface
func (s *WelfareBenefitService) UpdateWelfareBenefitService(id int, req models.UpdateWelfareBenefitRequest) error {
	_, err := s.UpdateWelfareBenefit(id, &req)
	return err
}

// DeleteWelfareBenefitService implements the port interface
func (s *WelfareBenefitService) DeleteWelfareBenefitService(id int) error {
	return s.DeleteWelfareBenefit(id)
}

// SearchWelfareBenefitService implements the port interface
func (s *WelfareBenefitService) SearchWelfareBenefitService(keyword string, page, pageSize int) (*models.WelfareBenefitListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.SearchWelfareBenefits(keyword, pageSize, offset)
}

// GetWelfareBenefitsCountService implements the port interface
func (s *WelfareBenefitService) GetWelfareBenefitsCountService() (int64, error) {
	return s.GetWelfareBenefitsCount()
}
