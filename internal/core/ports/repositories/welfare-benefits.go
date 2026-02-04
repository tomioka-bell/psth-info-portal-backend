package ports

import "backend/internal/core/domains"

type WelfareBenefitRepository interface {
	CreateWelfareBenefit(benefit *domains.WelfareBenefit) error
	GetWelfareBenefitByID(id int) (*domains.WelfareBenefit, error)
	GetAllWelfareBenefits(page, pageSize int) ([]domains.WelfareBenefit, int64, error)
	GetWelfareBenefitsByCategory(category string, page, pageSize int) ([]domains.WelfareBenefit, int64, error)
	UpdateWelfareBenefit(id int, updates map[string]interface{}) error
	DeleteWelfareBenefit(id int) error
	SearchWelfareBenefits(keyword string, page, pageSize int) ([]domains.WelfareBenefit, int64, error)
}
