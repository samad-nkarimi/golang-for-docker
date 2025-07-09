package repository

import "for-docker/models"

type ClientIPRepository interface {
	Save(ip *models.ClientIP) error
	GetLast5() ([]models.ClientIP, error)
}
