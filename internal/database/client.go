package database

import (
	"MicroserviceGo/internal/models"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, Id string) (*models.Customer, error)

	GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, Id string) (*models.Product, error)

	GetAllVendors(ctx context.Context) ([]models.Vendors, error)
	AddVendor(ctx context.Context, vendor *models.Vendors) (*models.Vendors, error)
	GetVendorById(ctx context.Context, Id string) (*models.Vendors, error)

	GetAllServices(ctx context.Context) ([]models.Services, error)
	AddService(ctx context.Context, service *models.Services) (*models.Services, error)
	GetServiceById(ctx context.Context, Id string) (*models.Services, error)
}

type Client struct {
	DB *gorm.DB
}

func (client Client) Ready() bool {
	//TODO implement me
	var ready string
	tx := client.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

type DBConnSettings struct {
	Hostname string
	Username string
	Password string
	Dbname   string
	Port     int
	Ssl      string
}

func NewDatabaseClient(settings DBConnSettings) (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		settings.Hostname,
		settings.Username,
		settings.Password,
		settings.Dbname,
		settings.Port,
		settings.Ssl)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
		Logger:      newLogger,
	})

	if err != nil {
		return nil, err
	}

	client := Client{
		DB: db,
	}
	return client, nil
}
