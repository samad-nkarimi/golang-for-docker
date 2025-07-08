package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	// "time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	_ = godotenv.Load()

	if os.Getenv("ENV") == "development" {
		println("in env")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found (skipping)")
		}
	} else {
		println("no env")
	}

	// Get database URL from env
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect to PostgreSQL
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	// Create table if not exists
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS client_ips (
	// 		id SERIAL PRIMARY KEY,
	// 		ip TEXT NOT NULL,
	// 		created_at TIMESTAMPTZ DEFAULT NOW()
	// 	)
	// `)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	http.HandleFunc("/", handleRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)

	// Insert IP into DB
	_, err := db.Exec("INSERT INTO client_ips (ip) VALUES ($1)", clientIP)
	if err != nil {
		http.Error(w, "Failed to insert IP", 500)
		log.Println("Insert error:", err)
		return
	}

	// Query last 5 IPs
	rows, err := db.Query(`
		SELECT ip FROM client_ips
		ORDER BY created_at DESC
		LIMIT 5
	`)
	if err != nil {
		http.Error(w, "Failed to query DB", 500)
		log.Println("Query error:", err)
		return
	}
	defer rows.Close()

	var ips []string
	for rows.Next() {
		var ip string
		if err := rows.Scan(&ip); err == nil {
			ips = append(ips, ip)
		}
	}

	// Return IPs as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ips)
}

// Extract client IP from request
func getClientIP(r *http.Request) string {
	// Try headers first (for reverse proxies)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}

	// Fallback to remote address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
