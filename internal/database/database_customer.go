package database

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (client Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := client.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}

func (client Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := client.DB.WithContext(ctx).
		Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflicError{}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (client Client) GetCustomerById(ctx context.Context, Id string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := client.DB.WithContext(ctx).
		Where(&models.Customer{
			CustomerID: Id,
		}).
		First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "customer",
				Id:     Id,
			}
		}
		return nil, result.Error
	}
	return customer, nil
}
