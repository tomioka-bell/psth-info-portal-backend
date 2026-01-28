package ports

import "backend/internal/core/domains"

type CompanyNewsRepository interface {
	CreateCompanyNews(companyNews *domains.CompanyNews) error
	GetCompanyNewsByID(companyNewsID string) (domains.CompanyNews, error)
	GetAllCompanyNews() ([]domains.CompanyNews, error)
	GetCompanyNews(limit, offset int) ([]domains.CompanyNews, int64, error)
	UpdateCompanyNewsWithMap(companyNewsID string, updates map[string]interface{}) error
	GetCompanyNewsCount() (int64, error)
	GetCompanyNewsByTitle(title string) (domains.CompanyNews, error)
	DeleteCompanyNews(companyNewsID string) error
}
