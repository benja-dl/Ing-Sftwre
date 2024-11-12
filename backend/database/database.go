package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/luccasniccolas/monitorWebApp/config"
)

func ConnectDatabase(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s  password=%s host=%s port=5432 sslmode=disable",
		cfg.DBUser, cfg.DBName, cfg.DBPass, cfg.DBHost)

	db, err := sql.Open("postgres", connStr)

	return db, err
}
