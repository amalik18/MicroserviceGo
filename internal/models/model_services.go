package models

type Services struct {
	ServicesId string  `gorm:"primaryKey" json:"servicesId"`
	Name       string  `json:"name"`
	Price      float64 `gorm:"type:numeric" json:"price"`
}
