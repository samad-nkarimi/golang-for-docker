package config

import (
	"fmt"
	"for-docker/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbURL string) {
	// dsn := os.Getenv(dbURL)
	println("")
	println(dbURL)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
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

func LoadEnv() {
	env := os.Getenv("ENV")

	if env == "" || env == "development" {
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Println("Warning: Could not load .env.dev file")
		} else {
			println("env loaded")
		}
	}
	// For production, rely on real environment variables
}
