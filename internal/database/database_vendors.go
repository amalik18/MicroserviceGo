package database

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (client Client) GetAllVendors(ctx context.Context) ([]models.Vendors, error) {
	var vendors []models.Vendors
	result := client.DB.WithContext(ctx).Find(&vendors)
	return vendors, result.Error
}

func (client Client) AddVendor(ctx context.Context, vendor *models.Vendors) (*models.Vendors, error) {
	vendor.VendorId = uuid.NewString()
	result := client.DB.WithContext(ctx).
		Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflicError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (client Client) GetVendorById(ctx context.Context, Id string) (*models.Vendors, error) {
	vendor := &models.Vendors{}
	result := client.DB.WithContext(ctx).
		Where(&models.Vendors{
			VendorId: Id,
		}).
		First(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "vendors",
				Id:     Id,
			}
		}
		return nil, result.Error
	}
	return vendor, nil
}
