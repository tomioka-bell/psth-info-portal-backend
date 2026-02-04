package repositories

import (
	"gorm.io/gorm"

	"backend/internal/core/domains"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

// GetAllTablesCount returns total count from all main tables
func (r *DashboardRepository) GetAllTablesCount() (map[string]int64, error) {
	counts := make(map[string]int64)

	// Count from ps_employees (no soft delete)
	var usersCount int64
	if err := r.db.Model(&domains.PSEmployee{}).Count(&usersCount).Error; err != nil {
		return nil, err
	}
	counts["users"] = usersCount

	// Count from ps_app_systems
	var appSystemsCount int64
	if err := r.db.Model(&domains.AppSystem{}).Where("deleted_at IS NULL").Count(&appSystemsCount).Error; err != nil {
		return nil, err
	}

	counts["app_systems"] = appSystemsCount
	var newsCount int64
	if err := r.db.Model(&domains.CompanyNews{}).Count(&newsCount).Error; err != nil {
		return nil, err
	}
	counts["company_news"] = newsCount

	var orgsCount int64
	if err := r.db.Model(&domains.Organization{}).Where("deleted_at IS NULL").Count(&orgsCount).Error; err != nil {
		return nil, err
	}
	counts["organizations"] = orgsCount

	var safetyCount int64
	if err := r.db.Model(&domains.SafetyDocument{}).Where("deleted_at IS NULL").Count(&safetyCount).Error; err != nil {
		return nil, err
	}
	counts["safety_documents"] = safetyCount

	var welfareCount int64
	if err := r.db.Model(&domains.WelfareBenefit{}).Where("deleted_at IS NULL").Count(&welfareCount).Error; err != nil {
		return nil, err
	}
	counts["welfare_benefits"] = welfareCount

	return counts, nil
}

// MonthlyStats holds monthly data
type MonthlyStats struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

// GetMonthlyStats returns monthly counts for a specific table
func (r *DashboardRepository) GetMonthlyStats(tableName string) ([]MonthlyStats, error) {
	switch tableName {
	case "ps_users":
		return r.GetUserMonthlyStats()
	case "ps_app_systems":
		return r.GetAppSystemMonthlyStats()
	case "ps_company_news":
		return r.GetNewsMonthlyStats()
	case "ps_safetys_documents":
		return r.GetSafetyDocMonthlyStats()
	case "welfare_benefits":
		return r.GetWelfareBenefitMonthlyStats()
	default:
		return nil, nil
	}
}

// GetUserMonthlyStats returns monthly user creation stats
func (r *DashboardRepository) GetUserMonthlyStats() ([]MonthlyStats, error) {
	// ps_employees doesn't track creation dates, return empty
	return []MonthlyStats{}, nil
}

// GetAppSystemMonthlyStats returns monthly app system creation stats
func (r *DashboardRepository) GetAppSystemMonthlyStats() ([]MonthlyStats, error) {
	var stats []MonthlyStats

	query := `
		SELECT 
			CONVERT(VARCHAR(7), created_at, 120) as month,
			COUNT(*) as count
		FROM ps_app_systems
		WHERE deleted_at IS NULL
		GROUP BY CONVERT(VARCHAR(7), created_at, 120)
		ORDER BY month DESC
		OFFSET 0 ROWS FETCH NEXT 12 ROWS ONLY
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetNewsMonthlyStats returns monthly company news creation stats
func (r *DashboardRepository) GetNewsMonthlyStats() ([]MonthlyStats, error) {
	var stats []MonthlyStats

	query := `
		SELECT 
			CONVERT(VARCHAR(7), created_at, 120) as month,
			COUNT(*) as count
		FROM ps_company_news
		WHERE deleted_at IS NULL
		GROUP BY CONVERT(VARCHAR(7), created_at, 120)
		ORDER BY month DESC
		OFFSET 0 ROWS FETCH NEXT 12 ROWS ONLY
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSafetyDocMonthlyStats returns monthly safety documents creation stats
func (r *DashboardRepository) GetSafetyDocMonthlyStats() ([]MonthlyStats, error) {
	var stats []MonthlyStats

	query := `
		SELECT 
			CONVERT(VARCHAR(7), created_at, 120) as month,
			COUNT(*) as count
		FROM ps_safetys_documents
		WHERE deleted_at IS NULL
		GROUP BY CONVERT(VARCHAR(7), created_at, 120)
		ORDER BY month DESC
		OFFSET 0 ROWS FETCH NEXT 12 ROWS ONLY
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetWelfareBenefitMonthlyStats returns monthly welfare benefits creation stats
func (r *DashboardRepository) GetWelfareBenefitMonthlyStats() ([]MonthlyStats, error) {
	var stats []MonthlyStats

	query := `
		SELECT 
			CONVERT(VARCHAR(7), created_at, 120) as month,
			COUNT(*) as count
		FROM welfare_benefits
		WHERE deleted_at IS NULL
		GROUP BY CONVERT(VARCHAR(7), created_at, 120)
		ORDER BY month DESC
		OFFSET 0 ROWS FETCH NEXT 12 ROWS ONLY
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// CategoryStats holds category breakdown data
type CategoryStats struct {
	Category string `json:"category"`
	Count    int64  `json:"count"`
}

// GetAppSystemsByCategory returns app systems grouped by category
func (r *DashboardRepository) GetAppSystemsByCategory() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			category,
			COUNT(*) as count
		FROM ps_app_systems
		WHERE deleted_at IS NULL
		GROUP BY category
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetOrganizationsByCategory returns organizations grouped by category
func (r *DashboardRepository) GetOrganizationsByCategory() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			category,
			COUNT(*) as count
		FROM ps_organizations
		WHERE deleted_at IS NULL
		GROUP BY category
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSafetyDocsByCategory returns safety documents grouped by category
func (r *DashboardRepository) GetSafetyDocsByCategory() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			category,
			COUNT(*) as count
		FROM ps_safetys_documents
		WHERE deleted_at IS NULL
		GROUP BY category
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSafetyDocsByDepartment returns safety documents grouped by department
func (r *DashboardRepository) GetSafetyDocsByDepartment() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			department as category,
			COUNT(*) as count
		FROM ps_safetys_documents
		WHERE deleted_at IS NULL AND department IS NOT NULL
		GROUP BY department
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetWelfareBenefitsByCategory returns welfare benefits grouped by category
func (r *DashboardRepository) GetWelfareBenefitsByCategory() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			category,
			COUNT(*) as count
		FROM welfare_benefits
		WHERE deleted_at IS NULL
		GROUP BY category
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// GetNewsByCategory returns company news grouped by category
func (r *DashboardRepository) GetNewsByCategory() ([]CategoryStats, error) {
	var stats []CategoryStats

	query := `
		SELECT 
			category,
			COUNT(*) as count
		FROM company_news
		WHERE deleted_at IS NULL
		GROUP BY category
		ORDER BY count DESC
	`

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}
