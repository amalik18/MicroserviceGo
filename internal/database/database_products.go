package database

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (client Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := client.DB.WithContext(ctx).
		Where(models.Product{VendorId: vendorId}).
		Find(&products)
	return products, result.Error
}

func (client Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductId = uuid.NewString()
	result := client.DB.WithContext(ctx).
		Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflicError{}
		}
		return nil, result.Error
	}
	return product, nil
}

func (client Client) GetProductById(ctx context.Context, Id string) (*models.Product, error) {
	product := &models.Product{}
	result := client.DB.WithContext(ctx).
		Where(&models.Product{
			ProductId: Id,
		}).
		First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "products",
				Id:     Id,
			}
		}
		return nil, result.Error
	}
	return product, nil
}
