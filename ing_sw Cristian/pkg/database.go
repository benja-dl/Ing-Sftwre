package pkg

import (
	"app/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		cfg.DBName, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		//log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}
