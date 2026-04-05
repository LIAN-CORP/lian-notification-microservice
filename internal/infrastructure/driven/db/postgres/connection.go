package postgres

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(databaseURL string, log *zap.SugaredLogger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Errorw("Failed to connect to database", "error", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorw("Failed to get database sql.DB", "error", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	
	log.Infow("Database connection established")
	return db, nil
}