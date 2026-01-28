package repositories

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/core/domains"
	ports "backend/internal/core/ports/repositories"
)

type CompanyNewsRepositoryDB struct {
	db *gorm.DB
}

func NewCompanyNewsRepositoryDB(db *gorm.DB) ports.CompanyNewsRepository {
	if err := db.AutoMigrate(&domains.CompanyNews{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &CompanyNewsRepositoryDB{db: db}
}

func (r *CompanyNewsRepositoryDB) CreateCompanyNews(n *domains.CompanyNews) error {
	now := time.Now()
	if n.CreatedAt.IsZero() {
		n.CreatedAt = now
	}
	if n.UpdatedAt.IsZero() {
		n.UpdatedAt = now
	}

	const q = `
	INSERT INTO company_news
		(company_news_id, company_news_photo, title, content, category, username_creator, created_at, updated_at)
	OUTPUT CONVERT(NVARCHAR(36), INSERTED.company_news_id) AS company_news_id, INSERTED.created_at, INSERTED.updated_at
	VALUES
		(NEWID(), ?, ?, ?, ?, ?, ?, ?);
`

	type ret struct {
		CompanyNewsID string
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	var out ret
	if err := r.db.
		Raw(q,
			n.CompanyNewsPhoto,
			n.Title,
			n.Content,
			n.Category,
			n.UsernameCreator,
			n.CreatedAt,
			n.UpdatedAt,
		).
		Scan(&out).Error; err != nil {

		fmt.Printf("CreateCompanyNews error: %v\n", err)
		return err
	}

	id, err := uuid.Parse(out.CompanyNewsID)
	if err != nil {
		fmt.Printf("failed to parse company news id: %v\n", err)
		return err
	}

	n.CompanyNewsID = id
	n.CreatedAt = out.CreatedAt
	n.UpdatedAt = out.UpdatedAt
	return nil
}

func (r *CompanyNewsRepositoryDB) GetCompanyNewsByID(companyNewsID string) (domains.CompanyNews, error) {
	const q = `
		SELECT company_news_id, company_news_photo, title, content, category,
		       username_creator, created_at, updated_at
		FROM company_news
		WHERE company_news_id = $1
		LIMIT 1;
	`

	var row domains.CompanyNews
	if err := r.db.Raw(q, companyNewsID).Scan(&row).Error; err != nil {
		return domains.CompanyNews{}, err
	}
	return row, nil
}

func (r *CompanyNewsRepositoryDB) GetAllCompanyNews() ([]domains.CompanyNews, error) {
	const q = `
		SELECT CONVERT(NVARCHAR(36), company_news_id) AS company_news_id, 
		       company_news_photo, title, content, category,
		       username_creator, created_at, updated_at
		FROM company_news
		ORDER BY created_at DESC;
	`

	type companyNewsRow struct {
		CompanyNewsID    string    `gorm:"column:company_news_id"`
		CompanyNewsPhoto string    `gorm:"column:company_news_photo"`
		Title            string    `gorm:"column:title"`
		Content          string    `gorm:"column:content"`
		Category         string    `gorm:"column:category"`
		UsernameCreator  string    `gorm:"column:username_creator"`
		CreatedAt        time.Time `gorm:"column:created_at"`
		UpdatedAt        time.Time `gorm:"column:updated_at"`
	}

	var rows []companyNewsRow
	if err := r.db.Raw(q).Scan(&rows).Error; err != nil {
		return nil, err
	}

	list := make([]domains.CompanyNews, 0, len(rows))
	for _, row := range rows {
		id, err := uuid.Parse(row.CompanyNewsID)
		if err != nil {
			fmt.Printf("GetAllCompanyNews parse UUID error: %v\n", err)
			continue
		}
		list = append(list, domains.CompanyNews{
			CompanyNewsID:    id,
			CompanyNewsPhoto: row.CompanyNewsPhoto,
			Title:            row.Title,
			Content:          row.Content,
			Category:         row.Category,
			UsernameCreator:  row.UsernameCreator,
			CreatedAt:        row.CreatedAt,
			UpdatedAt:        row.UpdatedAt,
		})
	}

	return list, nil
}

func (r *CompanyNewsRepositoryDB) GetCompanyNews(limit, offset int) ([]domains.CompanyNews, int64, error) {
	var total int64

	const countQ = `SELECT COUNT(*) FROM company_news;`
	if err := r.db.Raw(countQ).Scan(&total).Error; err != nil {
		fmt.Printf("GetCompanyNews count error: %v\n", err)
		return nil, 0, err
	}

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	const q = `
        SELECT CONVERT(NVARCHAR(36), company_news_id) AS company_news_id, 
               company_news_photo, title, content, category,
               username_creator, created_at, updated_at
        FROM company_news
        ORDER BY created_at DESC
        OFFSET ? ROWS
        FETCH NEXT ? ROWS ONLY;
    `

	type companyNewsRow struct {
		CompanyNewsID    string    `gorm:"column:company_news_id"`
		CompanyNewsPhoto string    `gorm:"column:company_news_photo"`
		Title            string    `gorm:"column:title"`
		Content          string    `gorm:"column:content"`
		Category         string    `gorm:"column:category"`
		UsernameCreator  string    `gorm:"column:username_creator"`
		CreatedAt        time.Time `gorm:"column:created_at"`
		UpdatedAt        time.Time `gorm:"column:updated_at"`
	}

	var rows []companyNewsRow
	if err := r.db.Raw(q, offset, limit).Scan(&rows).Error; err != nil {
		fmt.Printf("GetCompanyNews list error: %v\n", err)
		return nil, 0, err
	}

	list := make([]domains.CompanyNews, 0, len(rows))
	for _, row := range rows {
		id, err := uuid.Parse(row.CompanyNewsID)
		if err != nil {
			fmt.Printf("GetCompanyNews parse UUID error: %v\n", err)
			continue
		}
		list = append(list, domains.CompanyNews{
			CompanyNewsID:    id,
			CompanyNewsPhoto: row.CompanyNewsPhoto,
			Title:            row.Title,
			Content:          row.Content,
			Category:         row.Category,
			UsernameCreator:  row.UsernameCreator,
			CreatedAt:        row.CreatedAt,
			UpdatedAt:        row.UpdatedAt,
		})
	}

	return list, total, nil
}

func (r *CompanyNewsRepositoryDB) UpdateCompanyNewsWithMap(companyNewsID string, updates map[string]interface{}) error {
	if err := r.db.Model(&domains.CompanyNews{}).
		Where("company_news_id = ?", companyNewsID).
		Updates(updates).Error; err != nil {
		fmt.Printf("UpdateMeetingRoomWithMap error: %v\n", err)
		return err
	}
	return nil
}

func (r *CompanyNewsRepositoryDB) GetCompanyNewsCount() (int64, error) {
	const q = `SELECT COUNT(*) AS cnt FROM company_news;`
	var cnt int64
	if err := r.db.Raw(q).Scan(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func (r *CompanyNewsRepositoryDB) GetCompanyNewsByTitle(title string) (domains.CompanyNews, error) {
	const q = `
		SELECT TOP 1
			CONVERT(NVARCHAR(36), company_news_id) AS company_news_id,
			company_news_photo,
			title,
			content,
			category,
			username_creator,
			created_at,
			updated_at
		FROM company_news
		WHERE title = ?
		ORDER BY created_at DESC;
	`

	type companyNewsRow struct {
		CompanyNewsID    string    `gorm:"column:company_news_id"`
		CompanyNewsPhoto string    `gorm:"column:company_news_photo"`
		Title            string    `gorm:"column:title"`
		Content          string    `gorm:"column:content"`
		Category         string    `gorm:"column:category"`
		UsernameCreator  string    `gorm:"column:username_creator"`
		CreatedAt        time.Time `gorm:"column:created_at"`
		UpdatedAt        time.Time `gorm:"column:updated_at"`
	}

	var row companyNewsRow
	if err := r.db.Debug().Raw(q, title).Scan(&row).Error; err != nil {
		return domains.CompanyNews{}, err
	}

	id, err := uuid.Parse(row.CompanyNewsID)
	if err != nil {
		return domains.CompanyNews{}, fmt.Errorf("failed to parse UUID: %w", err)
	}

	return domains.CompanyNews{
		CompanyNewsID:    id,
		CompanyNewsPhoto: row.CompanyNewsPhoto,
		Title:            row.Title,
		Content:          row.Content,
		Category:         row.Category,
		UsernameCreator:  row.UsernameCreator,
		CreatedAt:        row.CreatedAt,
		UpdatedAt:        row.UpdatedAt,
	}, nil
}

func (r *CompanyNewsRepositoryDB) DeleteCompanyNews(companyNewsID string) error {
	result := r.db.Where("company_news_id = ?", companyNewsID).Delete(&domains.CompanyNews{})
	if result.Error != nil {
		fmt.Printf("DeleteCompanyNews error: %v\n", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("company news with ID %s not found", companyNewsID)
	}
	return nil
}
