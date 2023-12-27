package database

import (
	"MicroserviceGo/internal/models"
	"context"
)

func (client Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := client.DB.WithContext(ctx).
		Where(models.Product{VendorId: vendorId}).
		Find(&products)
	return products, result.Error
}
