package database

import (
	"MicroserviceGo/internal/models"
	"context"
)

func (client Client) GetAllServices(ctx context.Context) ([]models.Services, error) {
	var services []models.Services
	result := client.DB.WithContext(ctx).Where(models.Services{}).Find(&services)
	return services, result.Error
}
