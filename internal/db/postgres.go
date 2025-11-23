package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"avito_internship_PR/internal/config"
	"avito_internship_PR/internal/db/models"
)

type DB struct {
	GormDB *gorm.DB
}

func NewPostgresConnection(cfg *config.DB) (*DB, error) {
	cfg.DBString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	dsn := cfg.DBString
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to database!")
	return &DB{GormDB: gormDB}, nil
}

func AutoMigrate(db *gorm.DB) error {
	if db.Dialector.Name() == "postgres" {
		db.Exec("CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');")
	}
	return db.AutoMigrate(&models.User{}, &models.Team{}, &models.TeamMember{}, &models.PullRequest{})
}
