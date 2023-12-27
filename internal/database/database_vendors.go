package database

import (
	"MicroserviceGo/internal/models"
	"context"
)

func (client Client) GetAllVendors(ctx context.Context) ([]models.Vendors, error) {
	var vendors []models.Vendors
	result := client.DB.WithContext(ctx).Where(models.Vendors{}).Find(&vendors)
	return vendors, result.Error
}
