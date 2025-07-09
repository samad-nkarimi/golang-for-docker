package main

import (
	"fmt"
	"for-docker/config"
	"for-docker/models"
	"for-docker/repository"
	"net/http"
	"os"
	// "os"
	// "github.com/joho/godotenv"
	// "gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	config.LoadEnv()

	port := getEnv("PORT", "10000")
	dbURL := getEnv("DATABASE_URL", "postgres://sam:sampass@localhost:5432/mygodb?sslmode=disable")
	print(dbURL)
	config.InitDB(dbURL)

	// if os.Getenv("ENV") == "development" {
	// 	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.ClientIP{})
	// 	fmt.Println("🧹 All client IPs deleted")
	// }

	repo := repository.NewClientIPRepository()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		clientIP := &models.ClientIP{IPAddress: ip}
		_ = repo.Save(clientIP)

		ips, _ := repo.GetLast5()
		for _, item := range ips {
			fmt.Fprintf(w, "IP: %s\n", item.IPAddress)
		}
	})

	http.ListenAndServe(":"+port, nil)
}
