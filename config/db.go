package config

import (
	"fmt"
	"for-docker/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL") // e.g. postgres://user:pass@localhost:5432/db?sslmode=disable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	DB = db

	// Auto-migrate tables
	err = db.AutoMigrate(
		&models.ClientIP{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	fmt.Println("✅ Database connected")
}
