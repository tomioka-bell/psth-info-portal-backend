package services

import (
	"errors"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	portRepositories "backend/internal/core/ports/repositories"
	portServices "backend/internal/core/ports/services"
)

type OrganizationDocService struct {
	repo portRepositories.OrganizationDocRepository
}

func NewOrganizationDocService(repo portRepositories.OrganizationDocRepository) portServices.OrganizationDocService {
	return &OrganizationDocService{repo: repo}
}

func (s *OrganizationDocService) CreateOrganizationDoc(req *models.CreateOrganizationDocRequest) (*models.OrganizationDocResponse, error) {
	// Validate required fields
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	if req.Desc == "" {
		return nil, errors.New("desc is required")
	}
	if req.Department == "" {
		return nil, errors.New("department is required")
	}

	doc := domains.OrganizationDoc{
		Name:       req.Name,
		Desc:       req.Desc,
		Department: req.Department,
		FileName:   req.FileName,
	}

	if err := s.repo.CreateOrganizationDoc(&doc); err != nil {
		return nil, err
	}

	return &models.OrganizationDocResponse{
		OrganizationDocID: doc.OrganizationDocID,
		Name:              doc.Name,
		Desc:              doc.Desc,
		Department:        doc.Department,
		FileName:          doc.FileName,
		CreatedAt:         doc.CreatedAt,
		UpdatedAt:         doc.UpdatedAt,
	}, nil
}

func (s *OrganizationDocService) GetOrganizationDocByID(id int) (*models.OrganizationDocResponse, error) {
	doc, err := s.repo.GetOrganizationDocByID(id)
	if err != nil {
		return nil, err
	}

	return &models.OrganizationDocResponse{
		OrganizationDocID: doc.OrganizationDocID,
		Name:              doc.Name,
		Desc:              doc.Desc,
		Department:        doc.Department,
		FileName:          doc.FileName,
		CreatedAt:         doc.CreatedAt,
		UpdatedAt:         doc.UpdatedAt,
	}, nil
}

func (s *OrganizationDocService) GetAllOrganizationDocs(limit, offset int) (*models.OrganizationDocListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.GetAllOrganizationDocs(limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.OrganizationDocResponse
	for _, doc := range docs {
		responses = append(responses, models.OrganizationDocResponse{
			OrganizationDocID: doc.OrganizationDocID,
			Name:              doc.Name,
			Desc:              doc.Desc,
			Department:        doc.Department,
			FileName:          doc.FileName,
			CreatedAt:         doc.CreatedAt,
			UpdatedAt:         doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.OrganizationDocListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *OrganizationDocService) GetOrganizationDocByDepartment(department string, limit, offset int) (*models.OrganizationDocListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.GetOrganizationDocsByDepartment(department, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.OrganizationDocResponse
	for _, doc := range docs {
		responses = append(responses, models.OrganizationDocResponse{
			OrganizationDocID: doc.OrganizationDocID,
			Name:              doc.Name,
			Desc:              doc.Desc,
			Department:        doc.Department,
			FileName:          doc.FileName,
			CreatedAt:         doc.CreatedAt,
			UpdatedAt:         doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.OrganizationDocListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

func (s *OrganizationDocService) UpdateOrganizationDoc(id int, req *models.UpdateOrganizationDocRequest) (*models.OrganizationDocResponse, error) {
	// Validate
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	if req.Desc == "" {
		return nil, errors.New("desc is required")
	}
	if req.Department == "" {
		return nil, errors.New("department is required")
	}

	updates := map[string]interface{}{
		"name":       req.Name,
		"desc":       req.Desc,
		"department": req.Department,
	}

	if req.FileName != "" {
		updates["file_name"] = req.FileName
	}

	if err := s.repo.UpdateOrganizationDoc(id, updates); err != nil {
		return nil, err
	}

	// Get updated doc
	return s.GetOrganizationDocByID(id)
}

func (s *OrganizationDocService) DeleteOrganizationDoc(id int) error {
	return s.repo.DeleteOrganizationDoc(id)
}

func (s *OrganizationDocService) SearchOrganizationDocs(keyword string, limit, offset int) (*models.OrganizationDocListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	docs, total, err := s.repo.SearchOrganizationDocs(keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.OrganizationDocResponse
	for _, doc := range docs {
		responses = append(responses, models.OrganizationDocResponse{
			OrganizationDocID: doc.OrganizationDocID,
			Name:              doc.Name,
			Desc:              doc.Desc,
			Department:        doc.Department,
			FileName:          doc.FileName,
			CreatedAt:         doc.CreatedAt,
			UpdatedAt:         doc.UpdatedAt,
		})
	}

	page := offset/limit + 1
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &models.OrganizationDocListResponse{
		Data:       responses,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, nil
}

// ===== Port Interface Implementation =====

// CreateOrganizationDocService implements the port interface
func (s *OrganizationDocService) CreateOrganizationDocService(req models.CreateOrganizationDocRequest) error {
	_, err := s.CreateOrganizationDoc(&req)
	return err
}

// GetAllOrganizationDocService implements the port interface
func (s *OrganizationDocService) GetAllOrganizationDocService(page, pageSize int) (*models.OrganizationDocListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetAllOrganizationDocs(pageSize, offset)
}

// GetOrganizationDocByIDService implements the port interface
func (s *OrganizationDocService) GetOrganizationDocByIDService(id int) (*models.OrganizationDocResponse, error) {
	return s.GetOrganizationDocByID(id)
}

// GetOrganizationDocByDepartmentService implements the port interface
func (s *OrganizationDocService) GetOrganizationDocByDepartmentService(department string, page, pageSize int) (*models.OrganizationDocListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.GetOrganizationDocByDepartment(department, pageSize, offset)
}

// UpdateOrganizationDocService implements the port interface
func (s *OrganizationDocService) UpdateOrganizationDocService(id int, req models.UpdateOrganizationDocRequest) error {
	_, err := s.UpdateOrganizationDoc(id, &req)
	return err
}

// DeleteOrganizationDocService implements the port interface
func (s *OrganizationDocService) DeleteOrganizationDocService(id int) error {
	return s.DeleteOrganizationDoc(id)
}

// SearchOrganizationDocService implements the port interface
func (s *OrganizationDocService) SearchOrganizationDocService(keyword string, page, pageSize int) (*models.OrganizationDocListResponse, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	return s.SearchOrganizationDocs(keyword, pageSize, offset)
}
