package models

type Services struct {
	ServicesId string `gorm:"primaryKey" json:"servicesId"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
}
