package repository

import (
	"for-docker/config"
	"for-docker/models"
)

type clientIPRepo struct{}

func NewClientIPRepository() ClientIPRepository {
	return &clientIPRepo{}
}

func (r *clientIPRepo) Save(ip *models.ClientIP) error {
	return config.DB.Create(ip).Error
}

func (r *clientIPRepo) GetLast5() ([]models.ClientIP, error) {
	var ips []models.ClientIP
	err := config.DB.Order("created_at desc").Limit(5).Find(&ips).Error
	return ips, err
}
