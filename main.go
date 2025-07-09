package main

import (
	"fmt"
	"for-docker/config"
	"for-docker/models"
	"for-docker/repository"
	"net/http"
	// "os"

	"github.com/joho/godotenv"
	// "gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	config.InitDB()

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

	http.ListenAndServe(":10000", nil)
}
