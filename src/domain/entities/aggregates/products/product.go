package products

import (
	"kurdi-go/src/domain/entities"
	"kurdi-go/src/domain/entities/aggregates/categories"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	SKU        string                `json:"sku"`
	Details    ProductDetails        `json:"details"`
	Categories []categories.Category `json:"categories"`
	ProductPrice
	ProductQuantity
	entities.Entity
}

func (m *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Product) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	return
}

func (m *Product) AfterFind(tx *gorm.DB) (err error) {
	if m.IsDiscounted {
		m.SellingPrice = m.SellingPrice - m.Discount
	}
	return
}
