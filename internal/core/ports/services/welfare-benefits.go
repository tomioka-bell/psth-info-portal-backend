package ports

import "backend/internal/core/models"

type WelfareBenefitService interface {
	CreateWelfareBenefitService(req models.CreateWelfareBenefitRequest) error
	GetAllWelfareBenefitService(page, pageSize int) (*models.WelfareBenefitListResponse, error)
	GetWelfareBenefitByIDService(id int) (*models.WelfareBenefitResponse, error)
	GetWelfareBenefitByCategoryService(category string, page, pageSize int) (*models.WelfareBenefitListResponse, error)
	UpdateWelfareBenefitService(id int, req models.UpdateWelfareBenefitRequest) error
	DeleteWelfareBenefitService(id int) error
	SearchWelfareBenefitService(keyword string, page, pageSize int) (*models.WelfareBenefitListResponse, error)
}
