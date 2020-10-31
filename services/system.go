package services

import (
	"github.com/sctskw/attend.io/db"
	"github.com/sctskw/attend.io/models"
)

type SystemService interface {
	GetStatus() *models.SystemStatus
}

type systemService struct {
	db db.DatabaseClient
}

func NewSystemService(client db.DatabaseClient) SystemService {
	return &systemService{db: client}
}

func (s *systemService) GetStatus() *models.SystemStatus {
	return models.NewSystemStatus(200, "attend.io is alive.")
}
