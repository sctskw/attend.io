package services

import (
	"github.com/sctskw/attend.io/models"
)

func GetSystemStatus() *models.SystemStatus {
	return models.NewSystemStatus(200, "attend.io is running")
}
