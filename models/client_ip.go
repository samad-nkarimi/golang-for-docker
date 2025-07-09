package models

import "time"

type ClientIP struct {
	ID        uint   `gorm:"primaryKey"`
	IPAddress string `gorm:"not null"`
	CreatedAt time.Time
}
