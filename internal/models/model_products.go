package models

type Product struct {
	ProductId string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `gorm:"type:numeric" json:"price"`
	VendorId  string  `json:"vendorId"`
}
