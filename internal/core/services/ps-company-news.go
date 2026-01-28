package services

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"backend/internal/core/domains"
	"backend/internal/core/models"
	ports "backend/internal/core/ports/repositories"
	servicesports "backend/internal/core/ports/services"
	"backend/internal/pkgs/logs"
)

type CompanyNewsService struct {
	companyNewsRepo ports.CompanyNewsRepository
}

func NewCompanyNewsService(companyNewsRepo ports.CompanyNewsRepository) servicesports.CompanyNewsService {
	return &CompanyNewsService{companyNewsRepo: companyNewsRepo}
}

func (s *CompanyNewsService) CreateCompanyNewsService(req models.CompanyNewsResp) error {
	newID := uuid.New()

	domainISR := domains.CompanyNews{
		CompanyNewsID:    newID,
		CompanyNewsPhoto: req.CompanyNewsPhoto,
		Title:            req.Title,
		Content:          req.Content,
		Category:         req.Category,
		UsernameCreator:  req.UsernameCreator,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if s.companyNewsRepo == nil {
		log.Println("UserRepo is nil")
		return fmt.Errorf("user repository is not initialized")
	}

	err := s.companyNewsRepo.CreateCompanyNews(&domainISR)
	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *CompanyNewsService) GetCompanyNews(limit, offset int) (models.CompanyNewsListResp, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	query, total, err := s.companyNewsRepo.GetCompanyNews(limit, offset)
	if err != nil {
		return models.CompanyNewsListResp{}, err
	}

	jobs := make([]models.CompanyNewsReq, 0, len(query))
	for _, job := range query {
		jobs = append(jobs, models.CompanyNewsReq{
			CompanyNewsID:    job.CompanyNewsID,
			CompanyNewsPhoto: job.CompanyNewsPhoto,
			Title:            job.Title,
			Content:          job.Content,
			Category:         job.Category,
			UsernameCreator:  job.UsernameCreator,
			CreatedAt:        job.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        job.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// fmt.Println("CompanyNewsID : ", jobs)

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return models.CompanyNewsListResp{
		Data:       jobs,
		Total:      total,
		TotalPages: totalPages,
		Limit:      limit,
		Offset:     offset,
	}, nil
}

func (s *CompanyNewsService) GetCompanyNewsByTitle(title string) (models.CompanyNewsReq, error) {
	job, err := s.companyNewsRepo.GetCompanyNewsByTitle(title)
	if err != nil {
		return models.CompanyNewsReq{}, err
	}

	jobReq := models.CompanyNewsReq{
		CompanyNewsID:    job.CompanyNewsID,
		CompanyNewsPhoto: job.CompanyNewsPhoto,
		Title:            job.Title,
		Content:          job.Content,
		Category:         job.Category,
		UsernameCreator:  job.UsernameCreator,
		CreatedAt:        job.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        job.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return jobReq, nil
}

func (s *CompanyNewsService) UpdateCompanyNewsService(companyNewsID string, req models.CompanyNewsResp) error {
	log.Printf("[UpdateCompanyNewsService] Starting update for ID: %s\n", companyNewsID)
	log.Printf("[UpdateCompanyNewsService] Request data: %+v\n", req)

	if companyNewsID == "" {
		log.Println("[UpdateCompanyNewsService] Company news ID is empty!")
		return fmt.Errorf("company news ID is required")
	}

	if s.companyNewsRepo == nil {
		log.Println("[UpdateCompanyNewsService] CompanyNewsRepo is nil")
		return fmt.Errorf("company news repository is not initialized")
	}

	updates := make(map[string]interface{})

	if req.CompanyNewsPhoto != "" {
		log.Printf("[UpdateCompanyNewsService] Setting company_news_photo: %s\n", req.CompanyNewsPhoto)
		updates["company_news_photo"] = req.CompanyNewsPhoto
	}
	if req.Title != "" {
		log.Printf("[UpdateCompanyNewsService] Setting title: %s\n", req.Title)
		updates["title"] = req.Title
	}
	if req.Content != "" {
		log.Printf("[UpdateCompanyNewsService] Setting content: %s\n", req.Content)
		updates["content"] = req.Content
	}
	if req.Category != "" {
		log.Printf("[UpdateCompanyNewsService] Setting category: %s\n", req.Category)
		updates["category"] = req.Category
	}
	if req.UsernameCreator != "" {
		log.Printf("[UpdateCompanyNewsService] Setting username_creator: %s\n", req.UsernameCreator)
		updates["username_creator"] = req.UsernameCreator
	}

	if len(updates) == 0 {
		log.Println("[UpdateCompanyNewsService] No fields to update - all provided values are empty!")
		return fmt.Errorf("no fields to update")
	}

	log.Printf("[UpdateCompanyNewsService] Total fields to update: %d\n", len(updates))
	log.Printf("[UpdateCompanyNewsService] Updates map: %+v\n", updates)

	err := s.companyNewsRepo.UpdateCompanyNewsWithMap(companyNewsID, updates)
	if err != nil {
		log.Printf("[UpdateCompanyNewsService] Repository error: %v\n", err)
		logs.Error(err)
		return fmt.Errorf("failed to update company news: %w", err)
	}

	log.Println("[UpdateCompanyNewsService] Update completed successfully")
	return nil
}

func (s *CompanyNewsService) DeleteCompanyNewsService(companyNewsID string) error {
	log.Printf("[DeleteCompanyNewsService] Starting delete for ID: %s\n", companyNewsID)

	if companyNewsID == "" {
		log.Println("[DeleteCompanyNewsService] Company news ID is empty!")
		return fmt.Errorf("company news ID is required")
	}

	if s.companyNewsRepo == nil {
		log.Println("[DeleteCompanyNewsService] CompanyNewsRepo is nil")
		return fmt.Errorf("company news repository is not initialized")
	}

	err := s.companyNewsRepo.DeleteCompanyNews(companyNewsID)
	if err != nil {
		log.Printf("[DeleteCompanyNewsService] Repository error: %v\n", err)
		logs.Error(err)
		return fmt.Errorf("failed to delete company news: %w", err)
	}

	log.Println("[DeleteCompanyNewsService] Delete completed successfully")
	return nil
}
