package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDataBase() *gorm.DB {
	dsn := fmt.Sprintf(
		"sqlserver://%v:%v@%v:%v?database=%v&encrypt=disable",
		os.Getenv("SQLSERVER_USER"),
		os.Getenv("SQLSERVER_PASSWORD"),
		os.Getenv("SQLSERVER_HOST"),
		os.Getenv("SQLSERVER_PORT"),
		os.Getenv("SQLSERVER_DB"),
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 false,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("Database connection successful!")
	return db
}

// package database

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func InitDataBase() *gorm.DB {
// 	dsn := fmt.Sprintf(
// 		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
// 		os.Getenv("POSTGRES_HOST"),
// 		os.Getenv("POSTGRES_USER"),
// 		os.Getenv("POSTGRES_PASSWORD"),
// 		os.Getenv("POSTGRES_DB"),
// 		os.Getenv("POSTGRES_PORT"),
// 	)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 		DryRun:                 false,
// 		PrepareStmt:            true,
// 		Logger:                 logger.Default.LogMode(logger.Silent),
// 	})
// 	if err != nil {
// 		log.Fatalf("failed to connect database: %v", err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("failed to get generic database object: %v", err)
// 	}

// 	if err := sqlDB.Ping(); err != nil {
// 		log.Fatalf("failed to ping database: %v", err)
// 	}

// 	fmt.Println("Database connection successful!")
// 	return db
// }
