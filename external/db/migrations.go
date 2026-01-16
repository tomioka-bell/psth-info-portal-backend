package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

type MigrationRunner struct {
	db *gorm.DB
}

func NewMigrationRunner(db *gorm.DB) *MigrationRunner {
	return &MigrationRunner{db: db}
}

func (m *MigrationRunner) RunMigrations() error {
	migrationsDir := "migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Printf("Migrations directory does not exist: %s", migrationsDir)
		return nil
	}

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		if err := m.executeMigration(filePath); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
		}

		fmt.Printf("✓ Migration executed: %s\n", file.Name())
	}

	return nil
}

func (m *MigrationRunner) executeMigration(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Split SQL statements by GO batch separator (SQL Server specific)
	statements := m.splitStatements(string(content))

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		if err := m.db.Exec(stmt).Error; err != nil {
			return fmt.Errorf("failed to execute SQL: %w", err)
		}
	}

	return nil
}

// splitStatements splits SQL content by GO batch separator
func (m *MigrationRunner) splitStatements(content string) []string {
	lines := strings.Split(content, "\n")
	var statements []string
	var currentStmt strings.Builder

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "--") {
			continue
		}

		// Check if line is a GO batch separator
		if strings.ToUpper(trimmedLine) == "GO" {
			if stmt := strings.TrimSpace(currentStmt.String()); stmt != "" {
				statements = append(statements, stmt)
				currentStmt.Reset()
			}
			continue
		}

		// Add line to current statement
		if currentStmt.Len() > 0 {
			currentStmt.WriteString("\n")
		}
		currentStmt.WriteString(line)
	}

	// Add any remaining statement
	if stmt := strings.TrimSpace(currentStmt.String()); stmt != "" {
		statements = append(statements, stmt)
	}

	return statements
}
