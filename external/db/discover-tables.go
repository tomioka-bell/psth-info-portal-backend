package database

import (
	"fmt"

	"gorm.io/gorm"
)

// TableInfo represents information about a database table
type TableInfo struct {
	TableName  string
	ColumnName string
	DataType   string
}

// DiscoverEmployeeTables finds tables containing employee-related data
func DiscoverEmployeeTables(db *gorm.DB) ([]TableInfo, error) {
	var tables []TableInfo

	// Query to find tables with UHR_ or AD_ prefixed columns
	query := `
		SELECT 
			TABLE_NAME,
			COLUMN_NAME,
			DATA_TYPE
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA = 'dbo'
		  AND (COLUMN_NAME LIKE 'UHR_%' OR COLUMN_NAME LIKE 'AD_%')
		ORDER BY TABLE_NAME, COLUMN_NAME
	`

	if err := db.Raw(query).Scan(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

// PrintTableStructure prints the discovered table structure
func PrintTableStructure(db *gorm.DB) error {
	tables, err := DiscoverEmployeeTables(db)
	if err != nil {
		return err
	}

	if len(tables) == 0 {
		fmt.Println("❌ No tables with UHR_ or AD_ columns found in the database")
		return nil
	}

	fmt.Println("\n✓ Found employee-related tables:")
	fmt.Println("================================================")

	currentTable := ""
	for _, t := range tables {
		if t.TableName != currentTable {
			fmt.Printf("\nTable: %s\n", t.TableName)
			currentTable = t.TableName
		}
		fmt.Printf("  - %s (%s)\n", t.ColumnName, t.DataType)
	}

	return nil
}
