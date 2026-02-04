package services

import (
	"backend/internal/repositories"
)

type DashboardService struct {
	repo *repositories.DashboardRepository
}

func NewDashboardService(repo *repositories.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

// GetDashboardStats returns all counts and monthly stats
func (s *DashboardService) GetDashboardStats() (map[string]interface{}, error) {
	counts, err := s.repo.GetAllTablesCount()
	if err != nil {
		return nil, err
	}

	// Get monthly stats for each table
	userMonthly, _ := s.repo.GetUserMonthlyStats()
	appSystemsMonthly, _ := s.repo.GetAppSystemMonthlyStats()
	newsMonthly, _ := s.repo.GetNewsMonthlyStats()
	safetyMonthly, _ := s.repo.GetSafetyDocMonthlyStats()
	welfareMonthly, _ := s.repo.GetWelfareBenefitMonthlyStats()

	return map[string]interface{}{
		"counts": counts,
		"monthly_stats": map[string]interface{}{
			"users":            userMonthly,
			"app_systems":      appSystemsMonthly,
			"company_news":     newsMonthly,
			"safety_documents": safetyMonthly,
			"welfare_benefits": welfareMonthly,
		},
	}, nil
}

// GetTableCount returns count for a specific table
func (s *DashboardService) GetTableCount(tableName string) (int64, error) {
	counts, err := s.repo.GetAllTablesCount()
	if err != nil {
		return 0, err
	}

	if count, ok := counts[tableName]; ok {
		return count, nil
	}

	return 0, nil
}

// GetTableMonthlyStats returns monthly stats for a specific table
func (s *DashboardService) GetTableMonthlyStats(tableName string) (interface{}, error) {
	switch tableName {
	case "users":
		return s.repo.GetUserMonthlyStats()
	case "app_systems":
		return s.repo.GetAppSystemMonthlyStats()
	case "company_news":
		return s.repo.GetNewsMonthlyStats()
	case "safety_documents":
		return s.repo.GetSafetyDocMonthlyStats()
	case "welfare_benefits":
		return s.repo.GetWelfareBenefitMonthlyStats()
	default:
		return s.repo.GetMonthlyStats(tableName)
	}
}

// GetAppSystemsCategory returns app systems grouped by category
func (s *DashboardService) GetAppSystemsCategory() ([]repositories.CategoryStats, error) {
	return s.repo.GetAppSystemsByCategory()
}

// GetOrganizationsCategory returns organizations grouped by category
func (s *DashboardService) GetOrganizationsCategory() ([]repositories.CategoryStats, error) {
	return s.repo.GetOrganizationsByCategory()
}

// GetSafetyDocsCategory returns safety documents grouped by category
func (s *DashboardService) GetSafetyDocsCategory() ([]repositories.CategoryStats, error) {
	return s.repo.GetSafetyDocsByCategory()
}

// GetSafetyDocsDepartment returns safety documents grouped by department
func (s *DashboardService) GetSafetyDocsDepartment() ([]repositories.CategoryStats, error) {
	return s.repo.GetSafetyDocsByDepartment()
}

// GetWelfareBenefitsCategory returns welfare benefits grouped by category
func (s *DashboardService) GetWelfareBenefitsCategory() ([]repositories.CategoryStats, error) {
	return s.repo.GetWelfareBenefitsByCategory()
}

// GetNewsCategory returns company news grouped by category
func (s *DashboardService) GetNewsCategory() ([]repositories.CategoryStats, error) {
	return s.repo.GetNewsByCategory()
}
