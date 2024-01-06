package database

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (client Client) GetAllServices(ctx context.Context) ([]models.Services, error) {
	var services []models.Services
	result := client.DB.WithContext(ctx).Where(models.Services{}).Find(&services)
	return services, result.Error
}

func (client Client) AddService(ctx context.Context, service *models.Services) (*models.Services, error) {
	service.ServicesId = uuid.NewString()
	result := client.DB.WithContext(ctx).
		Create(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflicError{}
		}
		return nil, result.Error
	}
	return service, nil
}

func (client Client) GetServiceById(ctx context.Context, Id string) (*models.Services, error) {
	service := &models.Services{}
	result := client.DB.WithContext(ctx).
		Where(&models.Services{
			ServicesId: Id,
		}).
		First(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "services",
				Id:     Id,
			}
		}
		return nil, result.Error
	}
	return service, nil
}
