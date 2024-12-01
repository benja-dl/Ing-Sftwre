package pkg

import (
	"app/models"
	"gorm.io/gorm"
	"log"
)

// Ejecutar migraciones
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrated successfully.")
}
