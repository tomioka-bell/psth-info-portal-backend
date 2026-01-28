package domains

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CompanyNews struct {
	CompanyNewsID    uuid.UUID `gorm:"type:uniqueidentifier;primaryKey;default:NEWID()" json:"company_news_id"`
	CompanyNewsPhoto string    `json:"company_news_photo"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Category         string    `json:"category"`
	UsernameCreator  string    `json:"username_creator"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Scan implements the Scanner interface for database/sql
func (cn *CompanyNews) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch v := src.(type) {
	case []byte:
		fmt.Printf("[Scan] Received bytes: %v (len=%d)\n", v, len(v))
		if len(v) == 16 {
			// SQL Server stores GUID in little-endian format
			// We need to swap the byte order for the first three components
			swappedBytes := make([]byte, 16)

			// Swap first 4 bytes (Data1)
			swappedBytes[0], swappedBytes[1], swappedBytes[2], swappedBytes[3] = v[3], v[2], v[1], v[0]

			// Swap next 2 bytes (Data2)
			swappedBytes[4], swappedBytes[5] = v[5], v[4]

			// Swap next 2 bytes (Data3)
			swappedBytes[6], swappedBytes[7] = v[7], v[6]

			// Rest stays the same (Data4)
			copy(swappedBytes[8:], v[8:])

			fmt.Printf("[Scan] Original bytes: %v\n", v)
			fmt.Printf("[Scan] Swapped bytes:  %v\n", swappedBytes)

			id, err := uuid.FromBytes(swappedBytes)
			if err != nil {
				fmt.Printf("[Scan] Error converting bytes to UUID: %v\n", err)
				return err
			}
			fmt.Printf("[Scan] Final UUID: %s\n", id.String())
			cn.CompanyNewsID = id
		} else {
			fmt.Printf("[Scan] Parsing as string: %s\n", string(v))
			id, err := uuid.Parse(string(v))
			if err != nil {
				return err
			}
			cn.CompanyNewsID = id
		}
	case string:
		fmt.Printf("[Scan] Received string: %s\n", v)
		id, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		cn.CompanyNewsID = id
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return nil
}

// Value implements the Valuer interface for database/sql
func (cn CompanyNews) Value() (driver.Value, error) {
	return cn.CompanyNewsID.String(), nil
}

// MarshalJSON converts CompanyNews to JSON with UUID as string
func (cn *CompanyNews) MarshalJSON() ([]byte, error) {
	type Alias struct {
		CompanyNewsID    string    `json:"company_news_id"`
		CompanyNewsPhoto string    `json:"company_news_photo"`
		Title            string    `json:"title"`
		Content          string    `json:"content"`
		Category         string    `json:"category"`
		UsernameCreator  string    `json:"username_creator"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}
	return json.Marshal(&Alias{
		CompanyNewsID:    cn.CompanyNewsID.String(),
		CompanyNewsPhoto: cn.CompanyNewsPhoto,
		Title:            cn.Title,
		Content:          cn.Content,
		Category:         cn.Category,
		UsernameCreator:  cn.UsernameCreator,
		CreatedAt:        cn.CreatedAt,
		UpdatedAt:        cn.UpdatedAt,
	})
}

// type CompanyNews struct {
// 	CompanyNewsID    string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"  json:"company_news_id"`
// 	CompanyNewsPhoto string    `json:"company_news_photo"`
// 	Title            string    `json:"title"`
// 	Content          string    `json:"content"`
// 	Category         string    `json:"category"`
// 	UsernameCreator  string    `json:"username_creator"`
// 	CreatedAt        time.Time `json:"created_at"`
// 	UpdatedAt        time.Time `json:"updated_at"`
// }
